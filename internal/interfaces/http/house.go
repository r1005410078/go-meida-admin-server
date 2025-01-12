package http

import (
	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/app/services"
	"github.com/r1005410078/meida-admin-server/internal/domain/house"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/house/handler"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)

type HouseHttpHandlers struct {
	aggregateRepo house.IHouseAggregateRepository
	server *services.HouseServices
	eventBus *shared.EventBus
}


func NewHouseHandlers(aggregateRepo house.IHouseAggregateRepository, eventBus *shared.EventBus, server *services.HouseServices) *HouseHttpHandlers {
	return &HouseHttpHandlers{
		aggregateRepo,
		server,
		eventBus,
	}
}

// 获取房源列表
func (s *HouseHttpHandlers) ListHouseHandler(c *gin.Context) {
	// 	
}

// 保存房源
func (s *HouseHttpHandlers) SaveHouseHandler(c *gin.Context) {
	body := &command.SaveHouseCommand{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := handler.NewSaveHouseCommandHandler(s.aggregateRepo, s.eventBus).Handle(body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// 删除房源
func (s *HouseHttpHandlers) DeleteHouseHandler(c *gin.Context) {
	id := c.Param("id")

	body := &command.DeleteHouseCommand{
		ID: &id,
	}

	if err := handler.NewDeleteHouseCommandHandler(s.aggregateRepo, s.eventBus).Handle(body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// 保存房源标签
func (s *HouseHttpHandlers) SaveHouseTagsHandler(c *gin.Context) {
	body := &command.SaveHouseTagsCommand{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := handler.NewSaveTagCommandHandler(s.aggregateRepo, s.eventBus).Handle(body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// 保存房源多媒体
func (s *HouseHttpHandlers) SaveHouseMediasHandler(c *gin.Context) {
	body := &command.SaveHouseMediasCommand{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := handler.NewSaveHouseMediasCommandHandler(s.aggregateRepo, s.eventBus).Handle(body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// 保存房源经纬度
func (s *HouseHttpHandlers) SaveHouseLocationHandler(c *gin.Context) {	
	body := &command.SaveHouseLocationCommand{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := handler.NewSaveHouseLocationCommandHandler(s.aggregateRepo, s.eventBus).Handle(body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return	
	}

	c.JSON(200, gin.H{"message": "success"})
}

// 删除房源经纬度
func (s *HouseHttpHandlers) DeleteHouseLocationHandler(c *gin.Context) {
	body := &command.DeleteHouseLocationCommand{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := handler.NewDeleteHouseLocationCommandHandler(s.aggregateRepo, s.eventBus).Handle(body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}