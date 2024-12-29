package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/app/services"
	"github.com/r1005410078/meida-admin-server/internal/domain/user"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/command"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/handler"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)

type UserHttpHandlers struct {
	aggregateRepo user.IUserAggregateRepository
	server       *services.UserServices
	eventBus     *shared.EventBus
}

func NewUserHandlers(aggregateRepo user.IUserAggregateRepository, eventBus *shared.EventBus, server *services.UserServices) *UserHttpHandlers {
	return &UserHttpHandlers{
		aggregateRepo,
		server,
		eventBus,
	}
}

// 保存用户
func (h *UserHttpHandlers) Save(c *gin.Context) {
	var body command.SaveUserCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 执行保存用户命令
	if err := handler.NewSaveUserCommandHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 获取用户列表
func (h *UserHttpHandlers) GetUserList(c *gin.Context) {
	data, err := h.server.List()
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

// 删除用户
func (h *UserHttpHandlers) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := handler.NewDeleteUserCommandHandler(h.aggregateRepo, h.eventBus).Handle(&command.DeleteUserCommand{
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

// 更新用户状态
func (h *UserHttpHandlers) SaveUserStatus(c *gin.Context) {
	var body command.UserStatusCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := handler.NewUserStatusCommandHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 关联角色
func (h *UserHttpHandlers) AssoicatedRoles(c *gin.Context) {
	var body command.AssociatedRolesCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := handler.NewAssoicatedRolesCommandHandler(h.aggregateRepo, h.eventBus).Handle(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 注册用户
func (h *UserHttpHandlers) RegisterUserHandler(c *gin.Context) {
	var body command.RegisterCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := handler.NewRegisterUserHandler(h.aggregateRepo, h.eventBus).Handle(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 登陆
func (h *UserHttpHandlers) LoginUserHandler(c *gin.Context) {
	var body command.LoginInCommand
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := handler.NewLoginUserHandler(h.aggregateRepo, h.eventBus).Handle(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 退出
func (h *UserHttpHandlers) LogoutUserHandler(c *gin.Context) {
	var body command.LoggedOutCommand

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 获取用户id
	if userId, err := h.server.GetUserIdByToken(body.Token) ; err == nil {
		body.UserId = *userId
	}

	if err := handler.NewLogoutUserHandler(h.aggregateRepo, h.eventBus).Handle(c, body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 发送验证码
func (h *UserHttpHandlers) SendVerifyCodeHandler(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 发送验证码
	if err := h.server.SendEmailCode(body.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// 刷新token
func (h *UserHttpHandlers) RefreshTokenHandler(c *gin.Context) {
	body := struct {
		Token string `json:"token" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	h.server.RefreshLoginToken(body.Token)
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}