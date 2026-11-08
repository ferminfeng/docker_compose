# PostgreSQL 本地镜像与数据导入

已发布的 `ferminfu/postgres-18.3-pgbigm:v0.2` 包含 PostgreSQL 18.3、`pg_bigm` 1.2 和 `zh_CN.UTF-8` 系统 locale。其他电脑直接拉取该镜像即可。

## 后续完整构建与发布

以后需要更新镜像时，在 `develop/full` 目录用唯一的 `Dockerfile` 完整构建，使用新的版本标签；这一步会重新编译 `pg_bigm`：

```powershell
$imageTag = 'vNEXT' # 替换为新的版本标签
docker build --platform linux/amd64 -f ./etc/postgres-pgbigm/Dockerfile -t ferminfu/postgres-18.3-pgbigm:$imageTag ./etc/postgres-pgbigm
docker run --rm --platform linux/amd64 --entrypoint locale ferminfu/postgres-18.3-pgbigm:$imageTag -a
docker push ferminfu/postgres-18.3-pgbigm:$imageTag
```

第二条输出必须包含 `zh_CN.utf8`。推送前确认 Docker Desktop 已登录 `ferminfu`；不要覆盖已发布的版本标签。

## 每台电脑的容器配置

在 PostgreSQL 服务的 Compose 配置中使用已发布镜像，不需要在每台电脑上构建。Compose 文件已声明 `full_postgres_data` 具名卷；截至本次检查，运行中的旧容器仍使用匿名卷，重建后才会挂载新具名卷。两种卷的数据都存放在宿主机的 Docker 存储中；PostgreSQL 18 镜像的挂载点是 `/var/lib/postgresql`：

```yaml
services:
  full_postgres:
    image: ferminfu/postgres-18.3-pgbigm:v0.2
    volumes:
      - full_postgres_data:/var/lib/postgresql

volumes:
  full_postgres_data:
```

其他现有配置（端口、环境变量、启动命令和网络）保持原值。首次使用此具名卷时会得到空的本地数据库；旧匿名卷不会自动删除。确认需要重新导入后，再拉取镜像并重建服务：

```powershell
docker compose pull full_postgres
docker compose up -d --force-recreate full_postgres
docker exec full_postgres locale -a
```

最后一条必须包含 `zh_CN.utf8`。不要使用 `docker compose down -v` 或清理旧卷命令。

## 导出并导入示例数据库

以下操作在新容器就绪后执行。示例地址 `192.0.2.10` 是保留的文档用 IP，示例库名为 `demo_erp`；执行前替换命令中的示例地址和库名。`pg_dump` 只读取源库，`-W` 会交互式提示输入源库密码；不要把密码写进命令或文档。先从宿主机 PowerShell 进入容器：

```powershell
docker exec -it full_postgres bash
```

在容器内执行。首次恢复要求本地尚无同名库，因此第一条查询应无输出；`pg_dump` 必须正常退出：

```bash
psql -U postgres -d postgres -Atc "SELECT datname FROM pg_database WHERE datname='demo_erp';"
pg_dump -h 192.0.2.10 -p 5432 -U postgres -W -d demo_erp -Fc -b -f /tmp/source.dump
exit
```

回到宿主机 PowerShell，在仓库外的备份目录复制 dump，再次进入容器。若容器重建，可用 `docker cp ./source.dump full_postgres:/tmp/source.dump` 放回新容器：

```powershell
docker cp full_postgres:/tmp/source.dump ./source.dump
docker exec -it full_postgres bash
```

在容器内核对归档内容并恢复：

```bash
pg_restore -l /tmp/source.dump
pg_restore -U postgres -d postgres -C --no-owner --no-acl --exit-on-error -v /tmp/source.dump
```

`pg_restore -C` 会使用源库的 `zh_CN.UTF-8` 排序规则创建同名库，再恢复表、数据、索引、函数、触发器、扩展和序列值。`--no-owner --no-acl` 不复制源库的角色归属和授权；开发副本通常不需要这些集群级对象。若恢复报错，保留第一条错误并停止，不要向可能只恢复一部分的数据库重复运行。

## 验证与服务切换

继续在容器内执行验证，结束后退出到宿主机：

```bash
psql -U postgres -d postgres -c "SELECT datname, datcollate, datctype FROM pg_database WHERE datname='demo_erp';"
psql -U postgres -d demo_erp -c "SELECT extname, extversion FROM pg_extension WHERE extname IN ('pg_bigm', 'pg_trgm');"
psql -U postgres -d demo_erp -c "SELECT (SELECT count(*) FROM pg_class WHERE relnamespace='public'::regnamespace AND relkind IN ('r','p')) AS tables, (SELECT count(*) FROM pg_class WHERE relnamespace='public'::regnamespace AND relkind='S') AS sequences;"
exit
```

确认数据后，将 ERP 后端数据库连接改为本机映射端口对应的目标库，再重启后端。数据库 dump 应保存在仓库外，避免误提交业务数据。
