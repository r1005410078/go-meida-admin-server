create table if not exists house_property_aggregate (
  id UUID not null primary key,
  address VARCHAR(255) not null, -- 地址,不可重复
  is_synced BOOLEAN not null default false,  -- 是否在外网同步
  deleted_at TIMESTAMP, -- 删除时间
  created_at TIMESTAMP not null default now(),
  updated_at TIMESTAMP not null default current_timestamp on update now()
)

create table if not exists house_property (
  id UUID not null primary key,   
  purpose VARCHAR(255) not null, -- 用途
  transaction_type VARCHAR(255)  not null, -- 交易类型
  house_status VARCHAR(255), -- 交易状态
  owner_name VARCHAR(255)  not null, -- 业主姓名
  phone VARCHAR(255)  not null, -- 联系电话
  community_address VARCHAR(255) not null, -- 小区地址
  floor_range json, -- 楼层 { min, max }
  door_number json, -- 门牌号 { building_number, unit_number, door_number }
  title VARCHAR(255), -- 房源标题
  tags json, -- 推荐标签
  car_height DECIMAL(5, 2), -- 车位高度
  layout json, -- 户型 { room, hall, kitchen, bathroom, balcony }
  stairs json, -- 梯户 { stairs, rooms }
  actual_rate DECIMAL(5, 2), -- 实率
  level INTEGER, -- 级别
  floor_height DECIMAL(5, 2), -- 层高
  progress_depth DECIMAL(5, 2), -- 进深
  door_width DECIMAL(5, 2), -- 门宽
  building_area INTEGER not null, -- 建筑面积
  use_area DECIMAL(10, 2) not null, -- 使用面积
  sale_price DECIMAL(10, 2) not null default 0, -- 售价
  rent_price DECIMAL(10, 2) not null default 0, -- 租价
  rent_low_price DECIMAL(10, 2), -- 出租低价
  down_payment DECIMAL(10, 2), -- 首付
  sale_low_price DECIMAL(10, 2), -- 出售低价
  house_type VARCHAR(255), -- 房屋类型
  house_orientation VARCHAR(255), -- 房屋朝向
  house_decoration VARCHAR(255), -- 装修
  discount_year_limit INTEGER, -- 满减年限
  view_method VARCHAR(255), -- 看房方式
  payment_method VARCHAR(255), -- 付款方式
  property_tax DECIMAL(10, 2), -- 房源税费
  building_structure VARCHAR(255), -- 建筑结构
  building_year VARCHAR(255), -- 建筑年代
  property_rights VARCHAR(255), -- 产权性质
  property_year_limit INTEGER, -- 产权年限
  certificate_date DATE, -- 产权日期
  handover_date DATE, -- 交房日期
  degree VARCHAR(255), -- 学位
  household VARCHAR(255), -- 户口
  source VARCHAR(255), -- 来源
  delegate_number VARCHAR(255), -- 委托编号
  unique_housing BOOLEAN , -- 唯一住房
  full_payment BOOLEAN , -- 全款
  mortgage BOOLEAN , -- 抵押
  urgent BOOLEAN , -- 急切
  support VARCHAR(255), -- 配套
  present_state VARCHAR(255), -- 现状
  external_sync BOOLEAN not null default false, -- 外网同步
  remark TEXT , -- 备注
  images json , -- 图片
  location VARCHAR(255) , -- 位置
  deleted_at timestamp,
  created_at timestamp not null default now(),
  updated_at timestamp not null default current_timestamp on update now()
);

-- 房源修改快照
create table if not exists house_property_history (
  id UUID not null primary key,
  house_property_id UUID not null,
  house_property_data json not null,
  deleted_at timestamp,
  created_at TIMESTAMP not null default now(),
  updated_at TIMESTAMP not null default current_timestamp on update now()
)