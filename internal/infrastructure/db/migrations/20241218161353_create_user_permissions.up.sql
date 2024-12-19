
create table if not exists user_permissions (
	id char(36) PRIMARY KEY not null default (uuid()),
  name varchar(255) not null,
  description varchar(255) not null,
  action varchar(255) not null,
  create_at TIMESTAMP default current_timestamp,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);