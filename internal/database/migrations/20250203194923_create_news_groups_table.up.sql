create table news_groups
(
    uuid            uuid             primary key,
    default_title varchar(255)          not null,
    alias         varchar(255)          not null
        constraint news_groups_alias_unique unique,
    published     boolean default false not null,
    user_id       bigint,
    deleted_at    timestamp(0),
    created_at    timestamp(0),
    updated_at    timestamp(0)
);

create index news_groups_published_index
    on news_groups (published);

create index news_groups_user_id_index
    on news_groups (user_id);