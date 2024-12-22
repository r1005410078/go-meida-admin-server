
-- 权限聚合表
create table if not exists role_aggregate (
    id char(36) PRIMARY KEY not null default (uuid()),
    name varchar(255) unique not null,
    permission_ids json not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);

-- 角色表
create table if not exists roles (
    id char(36) PRIMARY KEY not null default (uuid()),
    name varchar(255) unique not null,
    description varchar(255),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);

-- 角色 与 权限关联表
create table if not exists roles_permission (
    role_id char(36) not null,
    permission_id char(36) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,

    PRIMARY KEY (role_id, permission_id)
);

