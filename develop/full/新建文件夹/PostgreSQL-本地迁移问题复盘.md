# PostgreSQL 数据无法导入本地库：原因与处理记录

更新日期：2026-10-27

## 1. 当前环境与结论

- 来源 A：PostgreSQL 18.6，数据库 `erp2027`，导出工具为 Navicat。
- 目标 B：Docker 中的 PostgreSQL 18.6（Debian `18.6-1.pgdg13+2`），Compose 服务名 `full_postgres`，宿主机端口 `14120`。用户确认目标库没有需要保留的数据。
- 导出文件：`C:\Users\fermin\Downloads\public.sql`，是 Navicat 生成的普通 SQL，包含序列、表、数据、函数和索引等对象。
- **核心问题有两层**：Navicat 的导入/导出顺序与扩展对象处理不合适；目标 PostgreSQL 镜像又没有安装 `pg_bigm` 的服务器端文件。A、B 都是 PostgreSQL 18.6，并不能保证两边安装了相同扩展。
- **截至本文编写时尚未导入成功**：用于补装 `pg_bigm` 的本地镜像构建被网络、证书问题阻断；当前 SQL 也仍包含扩展内部对象定义，不能把镜像构建成功等同于迁移完成。

## 2. 报错及原因：按发生顺序

| 阶段 | 代表错误 | 已确认的原因 |
| --- | --- | --- |
| Navicat 直接“数据传输” | `relation "action_category_id_seq" does not exist` | 建表语句的 `id` 默认值引用该序列，但执行建表时目标库还没有该序列。需要先有序列，再建表。 |
| 改为导入 Navicat SQL | `relation "public.approval_instances" does not exist` | `INSERT` 执行时表尚不存在；前面建表语句未成功执行或未被导入。勾选“遇错继续”只能跳过错误，不能补齐前置对象。 |
| 导入原始 SQL | `type "gtrgm" does not exist` | 原始导出曾尝试直接 `CREATE TYPE gtrgm`，但其输入/输出函数及所属的 `pg_trgm` 扩展对象未按正确依赖关系建立。 |
| 目标库安装 `pg_trgm` 后再次导入 | `cannot drop type gtrgm because extension pg_trgm requires it` | SQL 又尝试 `DROP TYPE gtrgm`。该类型归 `pg_trgm` 扩展管理，不能作为普通独立对象删除。**不应为通过导入而删除整个 `pg_trgm` 扩展。** |
| 继续导入 | `could not access file "pg_bigm": No such file or directory` | SQL 包含 `AS '$libdir/pg_bigm'` 的 C 函数及 `gin_bigm_ops` 索引；目标容器没有相应的 `pg_bigm.so` 等扩展文件。仅执行 `CREATE EXTENSION` 或手工建函数，无法凭空提供服务器端库。 |
| 为 `pg_bigm` 构建自定义镜像 | Docker Hub token 连接超时、APT 包下载失败、Debian HTTPS `certificate verify failed` | 镜像构建需要从 Docker Hub、Debian 软件源、GitHub 获取内容；当前 Docker 构建网络访问不稳定，且 HTTPS 证书链校验失败。日志中的 `198.18.0.x` 可能与本机代理/网络转发环境有关，但仅凭该地址不能确定具体配置错误。 |

这些错误不是同一个问题。前几项是对象依赖及导出 SQL 的问题，`pg_bigm` 错误是目标服务器缺少扩展文件，最后一项是补装扩展时的镜像构建网络问题。

## 3. 已做过的尝试与得到的经验

1. 从 Navicat 的直接数据传输改为导出、执行 `public.sql`，解决思路转向控制对象创建顺序。但 Navicat SQL 将扩展内部对象也展开为 `DROP/CREATE`，因此引出 `gtrgm` 和 `pg_bigm` 错误。
2. 在 B 库处理 `pg_trgm` 后，`DROP TYPE gtrgm` 被 PostgreSQL 正确阻止。这证明扩展已管理该对象，不能用普通 DDL 覆盖扩展内部类型。当前磁盘上的 `public.sql` 已不再包含最初的 `gtrgm` 类型段，而是从序列定义开始；这与最早贴出的报错片段不同。
3. 检查当前文件：约第 8033 行仍有 `AS '$libdir/pg_bigm'`，约第 8880 行起的索引使用 `public.gin_bigm_ops`；另有 `$libdir/pg_trgm` 函数定义。因此当前 SQL 仍不是“只含业务表和数据”的干净脚本。
4. 修改现有 `D:\www\docker_compose\develop\full\docker-compose.yml`：`full_postgres` 现指向 `local/full-postgres:18.6-pgbigm`，并配置本地 `build` 与 `shared_preload_libraries=pg_bigm`。新增 `D:\www\docker_compose\develop\full\etc\postgres-pgbigm\Dockerfile`，以官方 `postgres:18.6` 为基础构建 `pg_bigm` 1.2。**配置已写入文件，镜像未成功构建。**
5. 构建时先后遇到 Docker Hub 认证服务无法连接、HTTP Debian 软件源连接失败、切到 HTTPS 后证书校验失败。`apt-get update` 即使打印“Reading package lists”，也可能因部分索引失败而继续使用旧索引；不能据此认为构建环境已恢复。
6. 曾检查 `nnaka2992/docker-postgres-pgbigm`：查到的版本清单最高为 PostgreSQL 17.2，仓库主要提供 Dockerfile/构建流程，没有核实到可直接替换当前 `postgres:18.6` 的已发布镜像。因此它目前不能作为 PG18.6 的即用方案。

## 4. 现在卡在哪里

目前 Compose 引用的 `local/full-postgres:18.6-pgbigm` 是计划在本机构建的镜像，不是 Docker Hub 上现成可拉取的镜像。构建失败后运行 `docker-compose up -d`，出现 `pull access denied for local/full-postgres` 只是本地镜像尚不存在的后续表现；构建又会受 Docker Hub 和 Debian/GitHub 网络问题阻断。

即便以后获得安装了 `pg_bigm` 的 PG18.6 镜像，仍需重新处理 SQL：当前 `public.sql` 包含扩展自带函数、操作符类等定义。让扩展自行创建和管理这些对象，再导入业务对象，才是完整路径。

## 5. 建议的下一步顺序

1. **先选定可运行的 PG18.6 + `pg_bigm` 环境。** 可修复 Docker 构建网络/证书后构建现有 Dockerfile，或寻找已发布镜像。选第三方镜像时必须核实：确实有可拉取的标签、PostgreSQL 主版本为 18、支持 `linux/amd64`、镜像内有 `pg_bigm.so` 和扩展控制/SQL 文件。不要仅凭仓库名称或 README 推断。
2. **启动目标实例并验证扩展。** 容器能正常启动后，在目标数据库查询：

   ```sql
   SELECT name, default_version, installed_version
   FROM pg_available_extensions
   WHERE name IN ('pg_trgm', 'pg_bigm');
   ```

   `pg_available_extensions` 查不到 `pg_bigm`，说明镜像仍未安装扩展文件；查得到但 `installed_version` 为空，则该数据库尚未执行 `CREATE EXTENSION pg_bigm;`。如需 `pg_trgm`，也在目标数据库执行 `CREATE EXTENSION IF NOT EXISTS pg_trgm;`。
3. **优先从 A 库重新制作 PostgreSQL 原生备份，再恢复到空的 B 库。** 用 `pg_dump`/`pg_restore` 处理扩展归属和依赖，比继续手工删改 Navicat SQL 更可靠。先在 A 库确认扩展：

   ```sql
   SELECT extname, extversion, extnamespace::regnamespace AS schema_name
   FROM pg_extension
   WHERE extname IN ('pg_trgm', 'pg_bigm');
   ```

   若 A 库里的对象是通过扩展安装的，原生备份会按扩展记录处理，而不会像当前 Navicat SQL 一样逐条重建扩展内部函数。备份前仍需核对 A 库实际扩展状态；如果那些函数是手工装的，迁移策略要相应调整。
4. **只在目标空库做一次完整恢复，并验证。** 检查扩展已安装、业务表存在、关键表行数与 A 库相符、`pg_bigm` 索引存在、序列下一值没有落后于表中最大 ID。不要在部分成功的库上反复整份执行含 `DROP` 的 SQL；这容易产生混合状态或覆盖已导入对象。

## 6. 目前不能直接得出的结论

- 没有证据表明 PostgreSQL 18.6 本身不支持 `pg_bigm`；当前明确的问题是 B 容器缺文件以及构建环境下载失败。
- 没有完成的镜像构建、扩展验证或全库导入结果，因此不能宣称本地库已修好。
- `nnaka2992/docker-postgres-pgbigm` 可以作为构建方式参考，但目前未核实到适合 PG18.6 的现成镜像。

相关文件：`C:\Users\fermin\Downloads\public.sql`、`D:\www\docker_compose\develop\full\docker-compose.yml`、`D:\www\docker_compose\develop\full\etc\postgres-pgbigm\Dockerfile`。
