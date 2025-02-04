create table news_article_langs
(
    uuid              uuid         primary key,
    rid               uuid         not null
        constraint fk_news_article_langs_id
            references news_articles
            on update cascade on delete cascade,
    loc               varchar(5)   not null,
    title             varchar(255) not null,
    content           text         not null,
    created_at timestamp(0),
    updated_at timestamp(0),
    deleted_at        timestamp(0),
    short_description varchar(1000)
);

create index idx_news_article_langs_rid
    on news_article_langs (rid, loc);

create index idx_news_article_langs_rid2
    on news_article_langs (rid);
