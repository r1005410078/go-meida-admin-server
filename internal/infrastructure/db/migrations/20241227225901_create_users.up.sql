-- 创建用户表
-- 包含用户的基本信息、认证信息、状态信息和扩展信息
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,                    -- 用户唯一标识
    username VARCHAR(255) NOT NULL,          -- 用户名，不可重复
    email VARCHAR(255) UNIQUE,                      -- 邮箱，不可重复
    phone VARCHAR(20),                              -- 手机号
    full_name VARCHAR(255),                         -- 用户全名
    avatar_url TEXT,                                -- 头像URL地址
    gender VARCHAR(10),                             -- 性别
    birthday TIMESTAMP NULL,                        -- 出生日期
    address TEXT,                                   -- 地址
    password_hash VARCHAR(255) NOT NULL,            -- 密码哈希值
    role VARCHAR(50),                               -- 用户角色
    status VARCHAR(50),                             -- 账户状态
    deleted_at TIMESTAMP,                           -- 删除时间 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- 创建时间
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,    -- 更新时间
    last_login_at TIMESTAMP,                        -- 最后登录时间
    login_attempts INTEGER DEFAULT 0,               -- 登录尝试次数
    preferences JSON,                 -- 用户偏好设置，使用JSONB存储
    tags TEXT,                                    -- 用户标签数组
    referred_by VARCHAR(255),                       -- 推荐人ID

    UNIQUE (username, deleted_at)
);

-- 创建user聚合表，用于DDD聚合根存储
CREATE TABLE IF NOT EXISTS user_aggregate (
    user_id VARCHAR(255) PRIMARY KEY,               -- 用户聚合根ID
    username VARCHAR(255) NOT NULL,                 -- 用户名
    password_hash VARCHAR(255) NOT NULL,            -- 密码哈希值
    role VARCHAR(50),                               -- 用户角色
    status VARCHAR(50),                             -- 用户状态（如：active, inactive, blocked等）
    deleted_at TIMESTAMP,                           -- 软删除时间戳
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- 创建时间
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update current_timestamp,    -- 更新时间
    UNIQUE (username, deleted_at)                   -- 用户名在未删除状态下唯一
);