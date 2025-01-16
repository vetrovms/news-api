-- auto-generated definition
create table news_articles
(
    id            bigserial
        primary key,
    group_id      bigint                not null
        constraint news_articles_group_id_foreign
            references news_groups
            on update cascade on delete cascade,
    default_title varchar(255)          not null,
    alias         varchar(255)          not null,
    published     boolean default false not null,
    created_at    timestamp(0),
    updated_at    timestamp(0),
    deleted_at    timestamp(0),
    published_at  date
);

comment on column news_articles.group_id is 'News Group ID';

comment on column news_articles.default_title is 'Default Ttitle';

comment on column news_articles.alias is 'Alias';

comment on column news_articles.published_at is 'Published At';

alter table news_articles
    owner to postgres;

create index idx_news_articles_published_at
    on news_articles (published_at);

create index idx_news_articles_published
    on news_articles (published);
