FROM postgres:18.6 AS build

ARG PG_BIGM_VERSION=1.2-20250903

RUN sed -i 's|http://deb.debian.org|https://deb.debian.org|g' /etc/apt/sources.list.d/debian.sources \
    && apt-get -o Acquire::Retries=5 update \
    && apt-get -o Acquire::Retries=5 install -y --no-install-recommends gcc make libc6-dev postgresql-server-dev-18 ca-certificates wget \
    && rm -rf /var/lib/apt/lists/*

RUN wget -O /tmp/pg_bigm.tar.gz \
        "https://github.com/pgbigm/pg_bigm/archive/refs/tags/v${PG_BIGM_VERSION}.tar.gz" \
    && mkdir -p /tmp/pg_bigm \
    && tar -xzf /tmp/pg_bigm.tar.gz --strip-components=1 -C /tmp/pg_bigm \
    && cd /tmp/pg_bigm \
    && make USE_PGXS=1 PG_CONFIG="$(command -v pg_config)" \
    && make USE_PGXS=1 PG_CONFIG="$(command -v pg_config)" install

FROM postgres:18.6

COPY --from=build /usr/local/lib/postgresql/pg_bigm.so /usr/local/lib/postgresql/pg_bigm.so
COPY --from=build /usr/local/share/postgresql/extension/pg_bigm* /usr/local/share/postgresql/extension/
