package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/app/services"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/role/handler"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/repository"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)

type RoleHttpHandlers struct {
	aggregateRepo *repository.RoleAggregateRepository
	server *services.RoleServices
	eventBus *shared.EventBus
}

func NewRoleHandlers(aggregateRepo *repository.RoleAggregateRepository, eventBus *shared.EventBus, server *services.RoleServices) *RoleHttpHandlers {
	return &RoleHttpHandlers{
		aggregateRepo,
		server,
		eventBus,
	}
}

// 保存角色
func (r *RoleHttpHandlers) Save(c *gin.Context)  {
	var body command.SaveRoleCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// 执行保存角色命令
	if err := handler.NewSaveRoleCommandHandler(r.aggregateRepo, r.eventBus).Handle(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 获取角色列表
func (r *RoleHttpHandlers) GetRoleList(c *gin.Context)  {
	data, err := r.server.GetRoleList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// 删除角色
func (r *RoleHttpHandlers) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := handler.NewDeleteRoleCommandHandler(r.aggregateRepo, r.eventBus).Handle(&command.DeleteRoleCommand{
		Id: id,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 删除角色权限
func (r *RoleHttpHandlers) DeleteRolePermission(c *gin.Context) {
	body := command.DeletePermissionCommand{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
 
	if err := handler.NewDeletePermissionHandler(r.aggregateRepo, r.eventBus).Handle(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}