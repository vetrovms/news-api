-- auto-generated definition
create table news_groups
(
    id            bigserial             primary key,
    default_title varchar(255)          not null,
    alias         varchar(255)          not null
        constraint news_groups_alias_unique unique,
    published     boolean default false not null,
    deleted_at    timestamp(0),
    created_at    timestamp(0),
    updated_at    timestamp(0)
);

comment on column news_groups.default_title is 'Default Title';

comment on column news_groups.alias is 'Alias';

comment on column news_groups.published is 'Published';

alter table news_groups
    owner to postgres;

create index news_groups_published_index
    on news_groups (published);