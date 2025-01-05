-- 表单聚合表
CREATE TABLE IF NOT EXISTS forms_aggregate (
  form_id CHAR(36) NOT NULL,
  form_name VARCHAR(255) NOT NULL,
  related_ids JSON,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW(),

  PRIMARY KEY (form_id, form_name)
);

-- 表单字段聚合表
CREATE TABLE IF NOT EXISTS form_fields_aggregate (
  field_id CHAR(36) PRIMARY KEY NOT NULL DEFAULT (uuid()),
  form_id CHAR(36) NOT NULL,
  label VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW()
);

-- 表单表
CREATE TABLE IF NOT EXISTS forms (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW()
);



-- 字段表
CREATE TABLE IF NOT EXISTS forms_fields (
  field_id CHAR(36) PRIMARY KEY,
  form_id CHAR(36) NOT NULL,
  label VARCHAR(255) NOT NULL UNIQUE,
  type VARCHAR(255) NOT NULL,
  required BOOLEAN NOT NULL DEFAULT FALSE,
  placeholder VARCHAR(255) NOT NULL DEFAULT '',
  validation_rules JSON,
  options JSON,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW()
);

-- 关联表
-- 用途 = 住宅 and 交易类型 = 出售
CREATE TABLE IF NOT EXISTS forms_fields_dependencies (
  id CHAR(36) PRIMARY KEY, -- 主键
  form_id CHAR(36) NOT NULL, -- 表单
  field_id CHAR(36) NOT NULL, -- 被依赖字段 租金
  related_field_id CHAR(36) NOT NULL, -- 关联字段 用途
  condition_value VARCHAR(255) NOT NULL DEFAULT '', -- 条件值 住宅
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW()
);