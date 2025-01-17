FROM golang:latest

ARG ARG_NEWS_API_DB_DSN
ARG ARG_NEWS_API_DB_URL
ARG ARG_NEWS_API_WEB_PORT
ARG ARG_NEWS_API_LOG_PATH
ARG ARG_NEWS_API_UPLOAD_PATH

ENV NEWS_API_DB_DSN $ARG_NEWS_API_DB_DSN
ENV NEWS_API_DB_URL $ARG_NEWS_API_DB_URL
ENV NEWS_API_WEB_PORT $ARG_NEWS_API_WEB_PORT
ENV NEWS_API_LOG_PATH $ARG_NEWS_API_LOG_PATH
ENV NEWS_API_UPLOAD_PATH $ARG_NEWS_API_UPLOAD_PATH

# системні пакети
RUN apt-get update && \
    apt-get -y install --no-install-recommends --no-install-suggests \
    vim \
    nano \
    git \
    lsb-release \
    curl

# утиліта для міграцій golang-migrate
RUN wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb && \
    dpkg -i migrate.linux-amd64.deb

# очистка
RUN apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# гіт
RUN rm -rf /app \
    && mkdir -p /app \
    && cd /app \
    && git init \
    && git remote -v \
    && git remote add origin https://github.com/vetrovms/news-api.git \
    && git remote -v \
    && git fetch origin master \
    && git checkout master \
    && git config --global user.name "Thread" \
    && git config --global user.email thread@vipdesign.com

WORKDIR /app

# GO install
RUN go get -u github.com/gofiber/fiber/v2 && \
    go get gorm.io/gorm && \
    go get gorm.io/driver/postgres && \
    go get github.com/subosito/gotenv && \
    go mod tidy

# білд
RUN go build -C cmd -o app

# міграції та запуск застосунку
ENTRYPOINT migrate -database "$NEWS_API_DB_URL" -path internal/database/migrations up && cmd/app