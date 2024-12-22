package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/handler"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/repository"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)

type RoleHttpHandlers struct {
	repo *repository.RoleAggregateRepository
	eventBus *shared.EventBus
}

func NewRoleHandlers(repo *repository.RoleAggregateRepository, eventBus *shared.EventBus) *RoleHttpHandlers {
	return &RoleHttpHandlers{
		repo,
		eventBus,
	}
}

func (r *RoleHttpHandlers) Save(c *gin.Context)  {
	var body command.SaveRoleCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	// 执行保存角色命令
	if err := handler.NewSaveRoleCommandHandler(r.repo, r.eventBus).Handle(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
}