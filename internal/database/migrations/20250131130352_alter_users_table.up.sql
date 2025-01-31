ALTER TABLE news_articles
ADD user_id bigint;

CREATE INDEX idx_news_articles_user_id
ON news_articles (user_id);