-- 表单聚合表
CREATE TABLE IF NOT EXISTS forms_aggregate (
  form_id CHAR(36) NOT NULL,
  field_id CHAR(36) NOT NULL,
  form_name VARCHAR(255) NOT NULL,
  field_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW(),

  PRIMARY KEY (form_id, field_id),
  UNIQUE KEY unique_form_name (form_name, field_name),
  UNIQUE KEY unique_form_field (form_id, field_id, field_name)
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
CREATE TABLE IF NOT EXISTS forms_fields_dependencies (
  id CHAR(36) PRIMARY KEY,
  form_id CHAR(36) NOT NULL,
  field_id CHAR(36) NOT NULL,
  depends_on CHAR(36) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE NOW()
);