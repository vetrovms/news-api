create table file_uploads
(
    uuid                   uuid        primary key,
    entity_type          varchar(255)                               not null,
    entity_id            uuid                                       not null,
    name                 varchar(255)                               not null,
    path                 varchar(255) default ''::character varying not null,
    deleted_at           timestamp(0),
    created_at           timestamp(0),
    updated_at           timestamp(0)
);
