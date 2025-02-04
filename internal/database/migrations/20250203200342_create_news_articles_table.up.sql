-- auto-generated definition
create table news_articles
(
    uuid          uuid        primary key,
    group_id      uuid                not null
        constraint news_articles_group_id_foreign
            references news_groups
            on update cascade on delete cascade,
    default_title varchar(255)          not null,
    alias         varchar(255)          not null
        constraint news_articles_alias_unique unique,
    published     boolean default false not null,
    user_id       bigint,
    created_at    timestamp(0),
    updated_at    timestamp(0),
    deleted_at    timestamp(0),
    published_at  date
);

create index idx_news_articles_published_at
    on news_articles (published_at);

create index idx_news_articles_published
    on news_articles (published);

create index idx_news_articles_user_id
    on news_articles (user_id);
