-- auto-generated definition
create table news_group_langs
(
    id         bigserial             primary key,
    rid        bigint       not null
        constraint fk_news_group_langs_id
            references news_groups
            on update cascade on delete cascade,
    loc        varchar(5)   not null,
    title      varchar(255) not null,
    created_at timestamp(0),
    updated_at timestamp(0),
    deleted_at timestamp(0)
);

comment on column news_group_langs.rid is 'Related model ID';

comment on column news_group_langs.loc is 'Language';

comment on column news_group_langs.title is 'Title';

alter table news_group_langs
    owner to postgres;

create index idx_news_group_langs_rid
    on news_group_langs (rid, loc);

create index idx_news_group_langs_rid2
    on news_group_langs (rid);
