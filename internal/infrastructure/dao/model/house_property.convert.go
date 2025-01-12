package model

import "github.com/r1005410078/meida-admin-server/internal/domain/house/events"

func FormCreateHouseEvent(e *events.CreateHouseEvent) *HouseProperty {
	house := &HouseProperty{
		ID: *e.ID,
		Purpose: *e.Purpose,
		TransactionType: *e.TransactionType,
		HouseStatus: e.HouseStatus,
		OwnerName: e.OwnerName,
		Phone: e.Phone,
		CommunityAddress: e.Community,
		FloorRangeMin: e.FloorRangeMin,
		FloorRangeMax: e.FloorRangeMax,
		BuildingNumber: e.BuildingNumber,
		UnitNumber: e.UnitNumber,
		DoorNumber: e.DoorNumber,
		
	}

	if e.HouseDetails != nil {
		house.FloorNumber = e.HouseDetails.FloorNumber
		house.FloorNumberFrom = e.HouseDetails.FloorNumberFrom
		house.FloorNumberTo = e.HouseDetails.FloorNumberTo
		house.Title = e.HouseDetails.Title
		house.CarHeight = e.HouseDetails.CarHeight
		house.LayoutRoom = e.HouseDetails.LayoutRoom
		house.LayoutHall = e.HouseDetails.LayoutHall
		house.LayoutKitchen = e.HouseDetails.LayoutKitchen
		house.LayoutBathroom = e.HouseDetails.LayoutBathroom
		house.LayoutBalcony = e.HouseDetails.LayoutBalcony
		house.Stairs = e.HouseDetails.Stairs
		house.Rooms = e.HouseDetails.Rooms
		house.ActualRate = e.HouseDetails.ActualRate
		house.Level = e.HouseDetails.Level
		house.FloorHeight = e.HouseDetails.FloorHeight
		house.ProgressDepth = e.HouseDetails.ProgressDepth
		house.DoorWidth = e.HouseDetails.DoorWidth
		house.BuildingArea = e.HouseDetails.BuildingArea
		house.UseArea  = e.HouseDetails.UseArea
		house.SalePrice  = e.HouseDetails.SalePrice
		house.RentPrice = e.HouseDetails.RentPrice
		house.RentLowPrice = e.HouseDetails.RentLowPrice
		house.DownPayment = e.HouseDetails.DownPayment
		house.SaleLowPrice = e.HouseDetails.SaleLowPrice
		house.HouseType = e.HouseDetails.HouseType
		house.HouseOrientation = e.HouseDetails.HouseOrientation
		house.HouseDecoration = e.HouseDetails.HouseDecoration
		house.DiscountYearLimit = e.HouseDetails.DiscountYearLimit
		house.ViewMethod = e.HouseDetails.ViewMethod
		house.PaymentMethod = e.HouseDetails.PaymentMethod
		house.PropertyTax  = e.HouseDetails.PropertyTax
		house.BuildingStructure = e.HouseDetails.BuildingStructure
		house.BuildingYear = e.HouseDetails.BuildingYear
		house.PropertyRights = e.HouseDetails.PropertyRights
		house.PropertyYearLimit = e.HouseDetails.PropertyYearLimit
		house.CertificateDate = e.HouseDetails.CertificateDate
		house.HandoverDate = e.HouseDetails.HandoverDate
		house.Degree = e.HouseDetails.Degree
		house.Household = e.HouseDetails.Household
		house.Source = e.HouseDetails.Source
		house.DelegateNumber = e.HouseDetails.DelegateNumber
		house.UniqueHousing = e.HouseDetails.UniqueHousing
		house.FullPayment = e.HouseDetails.FullPayment
		house.Mortgage = e.HouseDetails.Mortgage
		house.Urgent = e.HouseDetails.Urgent
		house.Support = e.HouseDetails.Support
		house.PresentState = e.HouseDetails.PresentState
		house.ExternalSync = e.HouseDetails.ExternalSync
		house.Remark = e.HouseDetails.Remark
	}

	return house
}

func FormUpdateHouseEvent(e *events.UpdateHouseEvent) *HouseProperty {
	house := &HouseProperty{
		ID: *e.ID,
		Purpose: *e.Purpose,
		TransactionType: *e.TransactionType,
		HouseStatus: e.HouseStatus,
		OwnerName: e.OwnerName,
		Phone: e.Phone,
		CommunityAddress: e.Community,
		FloorRangeMin: e.FloorRangeMin,
		FloorRangeMax: e.FloorRangeMax,
		BuildingNumber: e.BuildingNumber,
		UnitNumber: e.UnitNumber,
		DoorNumber: e.DoorNumber,
		
	}

	if e.HouseDetails != nil {
		house.FloorNumber = e.HouseDetails.FloorNumber
		house.FloorNumberFrom = e.HouseDetails.FloorNumberFrom
		house.FloorNumberTo = e.HouseDetails.FloorNumberTo
		house.Title = e.HouseDetails.Title
		house.CarHeight = e.HouseDetails.CarHeight
		house.LayoutRoom = e.HouseDetails.LayoutRoom
		house.LayoutHall = e.HouseDetails.LayoutHall
		house.LayoutKitchen = e.HouseDetails.LayoutKitchen
		house.LayoutBathroom = e.HouseDetails.LayoutBathroom
		house.LayoutBalcony = e.HouseDetails.LayoutBalcony
		house.Stairs = e.HouseDetails.Stairs
		house.Rooms = e.HouseDetails.Rooms
		house.ActualRate = e.HouseDetails.ActualRate
		house.Level = e.HouseDetails.Level
		house.FloorHeight = e.HouseDetails.FloorHeight
		house.ProgressDepth = e.HouseDetails.ProgressDepth
		house.DoorWidth = e.HouseDetails.DoorWidth
		house.BuildingArea = e.HouseDetails.BuildingArea
		house.UseArea  = e.HouseDetails.UseArea
		house.SalePrice  = e.HouseDetails.SalePrice
		house.RentPrice = e.HouseDetails.RentPrice
		house.RentLowPrice = e.HouseDetails.RentLowPrice
		house.DownPayment = e.HouseDetails.DownPayment
		house.SaleLowPrice = e.HouseDetails.SaleLowPrice
		house.HouseType = e.HouseDetails.HouseType
		house.HouseOrientation = e.HouseDetails.HouseOrientation
		house.HouseDecoration = e.HouseDetails.HouseDecoration
		house.DiscountYearLimit = e.HouseDetails.DiscountYearLimit
		house.ViewMethod = e.HouseDetails.ViewMethod
		house.PaymentMethod = e.HouseDetails.PaymentMethod
		house.PropertyTax  = e.HouseDetails.PropertyTax
		house.BuildingStructure = e.HouseDetails.BuildingStructure
		house.BuildingYear = e.HouseDetails.BuildingYear
		house.PropertyRights = e.HouseDetails.PropertyRights
		house.PropertyYearLimit = e.HouseDetails.PropertyYearLimit
		house.CertificateDate = e.HouseDetails.CertificateDate
		house.HandoverDate = e.HouseDetails.HandoverDate
		house.Degree = e.HouseDetails.Degree
		house.Household = e.HouseDetails.Household
		house.Source = e.HouseDetails.Source
		house.DelegateNumber = e.HouseDetails.DelegateNumber
		house.UniqueHousing = e.HouseDetails.UniqueHousing
		house.FullPayment = e.HouseDetails.FullPayment
		house.Mortgage = e.HouseDetails.Mortgage
		house.Urgent = e.HouseDetails.Urgent
		house.Support = e.HouseDetails.Support
		house.PresentState = e.HouseDetails.PresentState
		house.ExternalSync = e.HouseDetails.ExternalSync
		house.Remark = e.HouseDetails.Remark
	}

	return house
}