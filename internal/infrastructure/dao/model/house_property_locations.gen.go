// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameHousePropertyLocation = "house_property_locations"

// HousePropertyLocation mapped from table <house_property_locations>
type HousePropertyLocation struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	HousePropertyID string    `gorm:"column:house_property_id;not null" json:"house_property_id"`
	Latitude        *float64  `gorm:"column:latitude" json:"latitude"`
	Longitude       *float64  `gorm:"column:longitude" json:"longitude"`
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName HousePropertyLocation's table name
func (*HousePropertyLocation) TableName() string {
	return TableNameHousePropertyLocation
}
