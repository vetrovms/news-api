create table news_group_langs
(
    uuid       uuid             primary key,
    rid        uuid       not null
        constraint fk_news_group_langs_id
            references news_groups
            on update cascade on delete cascade,
    loc        varchar(5)   not null,
    title      varchar(255) not null,
    created_at timestamp(0),
    updated_at timestamp(0),
    deleted_at timestamp(0)
);

create index idx_news_group_langs_rid
    on news_group_langs (rid, loc);

create index idx_news_group_langs_rid2
    on news_group_langs (rid);
