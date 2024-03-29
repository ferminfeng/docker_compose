# 从官方基础版本构建
FROM php:7.4-fpm

# replace sources.list
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN \
    export mc="-j$(nproc)" \
    && apt-get update \
    && apt-get install -y \
        # for git
        git \
        libcurl4-gnutls-dev \
        # for iconv mcrypt
        libmcrypt-dev \
        # for gd
        libfreetype6-dev libjpeg62-turbo-dev libpng-dev \
        # for bz2
        libbz2-dev \
        # for gmp
        libgmp-dev \
        # for soap wddx xmlrpc tidy xsl
        libxml2-dev libtidy-dev libxslt1-dev \
        # for zip
        libzip-dev \
        # for snmp
        libsnmp-dev snmp \
        # for pgsql pdo_pgsql
        libpq-dev \
        # for pspell
        libpspell-dev \
        # for pdo_firebird
        firebird-dev \
        # for pdo_dblib
        freetds-dev \
        # for ldap
        libldap2-dev \
        # for imap
        libc-client-dev libkrb5-dev \
        # for intl
        libicu-dev \
        # for gearman
        libgearman-dev \
        # for magick
        libmagickwand-dev \
        # for memcached
        zlib1g-dev libmemcached-dev \
        # for mongodb
        autoconf pkg-config libssl-dev \
        # for odbc pdo_odbc
        unixodbc-dev \
        # for mcrypt
        libmcrypt-dev \
        # for librabbitmq-dev (amqp)
        librabbitmq-dev \
        # for curl
        curl \
        # for vim
        vim \
        # for supervisor
        supervisor \
    #  ========== docker-php-ext install ===============================
    # for gd
    && docker-php-ext-configure gd --with-freetype=/usr/include/ --with-jpeg=/usr/include/ \
    && docker-php-ext-install $mc gd \
    # for bz2
    && docker-php-ext-install $mc bz2 \
    # for gmp
    && ln -s /usr/include/x86_64-linux-gnu/gmp.h /usr/include/gmp.h \
    && docker-php-ext-install $mc gmp \
    # for soap xmlrpc tidy xsl
    && docker-php-ext-install $mc soap xmlrpc tidy xsl \
    # for zip
    && docker-php-ext-install $mc zip \
    # for snmp
    && docker-php-ext-install $mc snmp \
    # for pgsql pdo_pgsql
    && docker-php-ext-install $mc pgsql pdo_pgsql \
    # for pspell
    && docker-php-ext-install $mc pspell \
    # for pdo_firebird
    && docker-php-ext-install $mc pdo_firebird \
    # for pdo_dblib
    && docker-php-ext-configure pdo_dblib --with-libdir=lib/x86_64-linux-gnu \
    && docker-php-ext-install $mc pdo_dblib \
    # for ldap
    && docker-php-ext-configure ldap --with-libdir=lib/x86_64-linux-gnu \
    && docker-php-ext-install $mc ldap \
    # for imap
    && docker-php-ext-configure imap --with-kerberos --with-imap-ssl \
    && docker-php-ext-install $mc imap \
    # for intl
    && docker-php-ext-install $mc intl \
    # no dependency extension
    && docker-php-ext-install $mc bcmath \
    && docker-php-ext-install $mc calendar \
    && docker-php-ext-install $mc exif \
    && docker-php-ext-install $mc gettext \
    && docker-php-ext-install $mc sockets \
    && docker-php-ext-install $mc dba \
    && docker-php-ext-install $mc mysqli \
    && docker-php-ext-install $mc pcntl \
    && docker-php-ext-install $mc pdo_mysql \
    && docker-php-ext-install $mc shmop \
    && docker-php-ext-install $mc sysvmsg \
    && docker-php-ext-install $mc sysvsem \
    && docker-php-ext-install $mc sysvshm \
    # ================ Install PECL extensions ====================
    # for mcrypt
    && pecl install mcrypt-1.0.3 && docker-php-ext-enable mcrypt \
    # for redis
    && pecl install redis && docker-php-ext-enable redis \
    # for imagick require PHP
    && pecl install imagick && docker-php-ext-enable imagick \
    # for memcached require PHP
    && pecl install memcached && docker-php-ext-enable memcached \
    # for mongodb
    && pecl install mongodb && docker-php-ext-enable mongodb \
    # for swoole
    && pecl install swoole && docker-php-ext-enable swoole \
    # for msgpack
    && pecl install msgpack && docker-php-ext-enable msgpack \
    # for yar
    && pecl install yar && docker-php-ext-enable yar \
    # for yaf
    && pecl install yaf && docker-php-ext-enable yaf \
    # for amqp
    && pecl install amqp && docker-php-ext-enable amqp \
    # clear tmp data
    && docker-php-source delete \
    && apt-get clean all \
    && rm -rf /var/lib/apt/lists/*  \
    && rm -rf /tmp/* \
    && rm -rf /var/tmp/* \
    && rm -rf /usr/share/doc/* \
    && echo 'PHP 7.4 extensions installed.'

# install composer
RUN curl -sS https://getcomposer.org/installer | php -- --filename=composer --install-dir=/bin
ENV PATH /root/.composer/vendor/bin:$PATH

# 切换目录
WORKDIR /var/www 
RUN chown -R www-data:www-data /var/www/

# vim
# RUN echo -e 'set encoding=utf-8' "\n" 'set fileencoding=utf-8' >> ~/.vimrc
# RUN source ~/.vimrc

# RUN supervisord -c /etc/supervisor/supervisord.conf