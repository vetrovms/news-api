-- auto-generated definition
create table file_uploads
(
    id                   bigserial
        primary key,
    entity_type          varchar(255)                               not null,
    entity_id            integer      default 0                     not null,
    name                 varchar(255)                               not null,
    path                 varchar(255) default ''::character varying not null,
    deleted_at           timestamp(0),
    created_at           timestamp(0),
    updated_at           timestamp(0)
);

comment on column file_uploads.entity_type is 'Entity';

comment on column file_uploads.entity_id is 'Entity ID';

comment on column file_uploads.name is 'Name';

comment on column file_uploads.path is 'path';

alter table file_uploads
    owner to postgres;
