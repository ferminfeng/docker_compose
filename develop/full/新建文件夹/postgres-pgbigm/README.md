# PostgreSQL 18.6 with pg_bigm

This custom image uses the official `postgres:18.6` image and builds `pg_bigm` 1.2 from the upstream release source. The upstream 1.2 documentation lists PostgreSQL 18 as a tested platform.

## Build and start

Set `POSTGRES_PASSWORD` in a `.env` file beside `compose.yaml`, then run:

```powershell
docker compose build
docker compose up -d
```

The Compose file expects an existing Docker network named `full_net`. Create it first if needed:

```powershell
docker network create full_net
```

This is a separate test container: it publishes host port `14121`, uses container name `full_postgres_pgbigm_test`, and stores data in `postgres_pgbigm_test_data`. It will not replace or mount the existing `full_postgres` container's data.

## Existing database data

Do not point this service at an existing PostgreSQL data directory until you have a verified backup and confirmed how the current container stores its data. The provided named volume starts empty; it does not automatically reuse an existing container's anonymous volume. Restore a `pg_dump` backup into this test container. Once verified, plan a separate cutover to port `14120` and the existing data storage. Never run `docker compose down -v` if the named volume contains data you need.

If you already use other `shared_preload_libraries`, include `pg_bigm` together with those library names in the Compose `command` value instead of replacing them.

## Enable the extension

After the server starts, connect to each database that needs the extension and run:

```sql
CREATE EXTENSION IF NOT EXISTS pg_bigm;
```

Verify availability and installation with:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'pg_bigm';
```

The SQL dump should not drop or recreate extension-owned objects such as `bigm_similarity`, `gin_bigm_ops`, or the `gtrgm` type. Let PostgreSQL extensions manage their own functions, operator classes, and types. The dump's table definitions and data can then be imported, followed by the `pg_bigm` indexes if the extension was not available when those index statements ran.

## References

- [Official PostgreSQL Docker image](https://hub.docker.com/_/postgres)
- [pg_bigm upstream repository](https://github.com/pgbigm/pg_bigm)
- [pg_bigm 1.2 documentation](https://github.com/pgbigm/pg_bigm/blob/REL1_2_STABLE/docs/pg_bigm_en.md)
