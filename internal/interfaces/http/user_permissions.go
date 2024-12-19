package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/domain/permissions"
)

type UserPermissionsHandlers struct {
	service *permissions.PermissionsService
}

func NewUserPermissionsHandlers(service *permissions.PermissionsService) *UserPermissionsHandlers {
	return &UserPermissionsHandlers{
		service,
	}
}

func (u *UserPermissionsHandlers) List(c *gin.Context)  {
	
	list, err := u.service.List();

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"data": list,
	})
}

func (u *UserPermissionsHandlers) Save(c *gin.Context) {
	var  permission permissions.Permission
	
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
	}

	if err := u.service.Save(&permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	
	c.JSON(200, gin.H{
		"data": "success",
	})
}