-- auto-generated definition
create table news_article_langs
(
    id            bigserial
        primary key,
    rid               bigint       not null
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

comment on column news_article_langs.rid is 'Related model ID';

comment on column news_article_langs.loc is 'Language';

comment on column news_article_langs.title is 'Title';

comment on column news_article_langs.content is 'Content';

comment on column news_article_langs.short_description is 'Short Description';

alter table news_article_langs
    owner to postgres;

create index idx_news_article_langs_rid
    on news_article_langs (rid, loc);

create index idx_news_article_langs_rid2
    on news_article_langs (rid);
