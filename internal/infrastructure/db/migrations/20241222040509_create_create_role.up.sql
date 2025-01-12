
-- 权限聚合表
create table if not exists role_aggregate (
    role_id char(36) PRIMARY KEY not null, -- 角色id,
    role_name varchar(255) unique not null, -- 角色名称
    permission_ids json not null, -- 权限id列表
    deleted_at timestamp, -- 删除时间
    created_at timestamp default current_timestamp, -- 创建时间
    updated_at timestamp default current_timestamp on update current_timestamp, -- 更新时间
    UNIQUE (role_name, deleted_at)
);

-- 角色表
create table if not exists roles (
    id char(36) PRIMARY KEY not null,
    name varchar(255) not null,
    description varchar(255),
    deleted_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    UNIQUE (name, deleted_at)
);

-- 角色 与 权限关联表
create table if not exists roles_permission (
    role_id char(36) not null,
    permission_id char(36) not null,
    deleted_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (role_id, permission_id)
);

