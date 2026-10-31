ARG POSTGRES_IMAGE=postgres:18.3
FROM ${POSTGRES_IMAGE} AS build

ARG PG_BIGM_VERSION=1.2-20250903

RUN apt-get -o Acquire::Retries=3 -o APT::Update::Error-Mode=any update \
    && apt-get -o Acquire::Retries=3 install -y --no-install-recommends gcc make libc6-dev postgresql-server-dev-${PG_MAJOR} ca-certificates wget \
    && rm -rf /var/lib/apt/lists/*

# PostgreSQL 的 locale 头文件也依赖 ICU 开发头文件。
RUN apt-get -o Acquire::Retries=3 -o APT::Update::Error-Mode=any update \
    && apt-get -o Acquire::Retries=3 install -y --no-install-recommends libicu-dev \
    && rm -rf /var/lib/apt/lists/*

ARG PG_BIGM_SHA256=4d4fb48161e3f9e99207851299379c83b3e5d1d3c1a07c35f6a3153d577b3142
RUN wget -O /tmp/pg_bigm.tar.gz \
        "https://github.com/pgbigm/pg_bigm/archive/refs/tags/v${PG_BIGM_VERSION}.tar.gz" \
    && echo "${PG_BIGM_SHA256}  /tmp/pg_bigm.tar.gz" | sha256sum -c - \
    && mkdir -p /tmp/pg_bigm \
    && tar -xzf /tmp/pg_bigm.tar.gz --strip-components=1 -C /tmp/pg_bigm \
    && cd /tmp/pg_bigm \
    && make USE_PGXS=1 PG_CONFIG="$(command -v pg_config)" \
    && make USE_PGXS=1 PG_CONFIG="$(command -v pg_config)" install \
    && mkdir -p /extension-artifacts/lib /extension-artifacts/share \
    && cp "$(pg_config --pkglibdir)/pg_bigm.so" /extension-artifacts/lib/ \
    && cp "$(pg_config --sharedir)"/extension/pg_bigm* /extension-artifacts/share/

FROM ${POSTGRES_IMAGE}

COPY --from=build /extension-artifacts/ /tmp/extension-artifacts/
RUN cp /tmp/extension-artifacts/lib/* "$(pg_config --pkglibdir)/" \
    && cp /tmp/extension-artifacts/share/* "$(pg_config --sharedir)/extension/" \
    && rm -rf /tmp/extension-artifacts
