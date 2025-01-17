# API Каталог товарів  
Веб застосунок з API за протоколом HTTP для обслуговування списку новин.

Змінні оточення для локального запуску (без докер)

```bash
export NEWS_API_LOG_PATH="/home/inside/go/src/study/news/log/log.json" && \
export NEWS_API_DB_DSN="user=postgres password=postgres host=127.0.0.1 port=5432 dbname=gonews sslmode=disable" && \
export NEWS_API_DB_URL="postgres://postgres:postgres@127.0.0.1:5432/gonews?sslmode=disable" && \
export NEWS_API_WEB_PORT=8000 && \
export NEWS_API_UPLOAD_PATH="/home/inside/go/src/study/news/uploads/"
```

Змінні оточення для запуску в докер

```bash
export NEWS_API_LOG_PATH="/app/log/log.json" && \
export NEWS_API_DB_DSN="user=postgres password=postgres host=postgres_news_api port=5432 dbname=gonews sslmode=disable" && \
export NEWS_API_DB_URL="postgres://postgres:postgres@postgres_news_api:5432/gonews?sslmode=disable" && \
export NEWS_API_WEB_PORT=8000 && \
export NEWS_API_UPLOAD_PATH="/app/uploads/" && \
export NEWS_API_POSTGRES_PASSWORD=postgres && \
export NEWS_API_POSTGRES_USER=postgres && \
export NEWS_API_POSTGRES_DATABASE=gonews && \
export POSTGRES_USER=postgres && \
export POSTGRES_DB=gonews
```
