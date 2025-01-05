package http

import (
	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/app/services"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/forms/handler"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)

type FormsHttpHandlers struct {
 	aggregateRepo forms.IFormsAggregateRepository
	server *services.FormsServices
	eventBus *shared.EventBus
}

func NewFormsHandlers(aggregateRepo forms.IFormsAggregateRepository, eventBus *shared.EventBus, server *services.FormsServices) *FormsHttpHandlers {
	return &FormsHttpHandlers{
		aggregateRepo,
		server,
		eventBus,
	}
}

func (h *FormsHttpHandlers) SaveFormsHandler(c *gin.Context) {
	var body command.SaveFormsCommand
	
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := handler.NewSaveFormsHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func (h *FormsHttpHandlers) DeleteFormsHandler(c *gin.Context) {
	var body command.DeleteFormsCommand

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := handler.NewDeleteFormsCommandHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(200, gin.H{"message": "success"})
}

func (h *FormsHttpHandlers) SaveFormsFieldsHandler(c *gin.Context) {
	var body command.SaveFormsFieldsCommand

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := handler.NewSaveFormsFieldsHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(200, gin.H{"message": "success"})
}

func (h *FormsHttpHandlers) DeleteFormsFiledsHandler(c *gin.Context) {
	var body command.DeleteFormsFieldsCommand

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := handler.NewDeleteFormsFiledsHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(200, gin.H{"message": "success"})
}

func (h *FormsHttpHandlers) SaveDependsOnHandler(c *gin.Context) {
	var body command.SaveDependsOnCommand

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := handler.NewSaveDependsOnCommandHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}