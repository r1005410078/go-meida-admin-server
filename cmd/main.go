package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/app/services"
	"github.com/r1005410078/meida-admin-server/internal/domain/permissions"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/db"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/repository"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/http"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取当前工作目录: %v", err)
	}

	// 确保日志目录存在
	logDir := filepath.Join(currentDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("无法创建日志目录: %v", err)
	}

	// 配置日志
	logConfig := zap.NewProductionConfig()
	
	// 设置输出路径
	logConfig.OutputPaths = []string{
		"stdout",
		filepath.Join(logDir, "server.log"),
	}
	logConfig.ErrorOutputPaths = []string{
		"stderr",
		filepath.Join(logDir, "error.log"),
	}

	// 配置编码器
	logConfig.EncoderConfig.TimeKey = "timestamp"
	logConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logConfig.EncoderConfig.StacktraceKey = "stacktrace"
	logConfig.EncoderConfig.MessageKey = "message"
	logConfig.EncoderConfig.LevelKey = "level"
	logConfig.EncoderConfig.CallerKey = "caller"
	logConfig.EncoderConfig.FunctionKey = "func"
	
	// 开启开发环境更详细的日志
	if gin.Mode() == gin.DebugMode {
		logConfig.Development = true
		logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	
	logger, err := logConfig.Build(
		zap.AddCaller(),      // 添加调用者信息
		zap.AddCallerSkip(1), // 跳过一层调用栈，直接显示业务代码位置
	)
	if err != nil {
		log.Fatalf("无法初始化日志: %v", err)
	}
	defer logger.Sync()

	// 替换全局logger
	zap.ReplaceGlobals(logger)

	r := gin.Default()

	// 根据环境变量设置 Gin 模式，默认是 debug 模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}

	mysqlDb, err := db.GetDB()
	if err != nil {
		panic(err)
	}

	redisDb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "rts2778205", // no password set
		DB:		  0,  // use default DB
	})
	
	// 初始化角色服务
	roleRepo := repository.NewRoleRepository(mysqlDb)
	roleServices := services.NewRepoServices(roleRepo, logger)

	// 初始化用户服务
	userRepo := repository.NewUserRepository(mysqlDb, redisDb, context.Background())
	userServices := services.NewUserServices(userRepo, logger)

	// 注册事件总线
	bus := shared.NewEventBus()

	// 注册角色相关事件处理器
	bus.Register(roleServices.DeleteRoleEventHandle)
	bus.Register(roleServices.RoleDeleteFailedEventHandle)
	bus.Register(roleServices.SaveRoleEventHandle)
	bus.Register(roleServices.RoleSaveFailedEventHandle)

	// 注册用户相关事件处理器
	bus.Register(userServices.SaveUserEventHandle)
	bus.Register(userServices.SaveUserFailedEventHandle)
	bus.Register(userServices.DeleteUserHandle)
	bus.Register(userServices.DeleteUserFailedEventHandle)
	bus.Register(userServices.AssoicatedRolesEventHandle)
	bus.Register(userServices.AssoicatedRolesFailedEventHandle)
	bus.Register(userServices.SaveUserStatusEventHandle)
	bus.Register(userServices.SaveUserStatusFailedEventHandle)

	// 注册用户注册事件处理器
	bus.Register(userServices.RegisterEventEventHandle)
	bus.Register(userServices.RegisterFailedEventHandle)
	bus.Register(userServices.LoginInEventHandle)
	bus.Register(userServices.LoginFailedEventHandle)
	bus.Register(userServices.LogoutEventHandle)
	bus.Register(userServices.LogoutFailedEventHandle)

	// 初始化表单服务
	formsAggregateRepo := repository.NewFormsAggregateRepository(mysqlDb)
	formRepo := repository.NewFormsRepository(mysqlDb)
	formServices := services.NewFormsServices(formRepo, logger)

	// 初始化表单事件处理器
	bus.Register(formServices.CreateFormsEventHandle)
	bus.Register(formServices.UpdateFormsEventHandle)
	bus.Register(formServices.SaveFormsFailedEventHandle)
	bus.Register(formServices.DeleteFormsEventHandle)
	bus.Register(formServices.DeleteFormsFailedEventHandle)
	bus.Register(formServices.CreateFormsFieldsEventHanlde)
	bus.Register(formServices.UpdateFormsFieldsEventHanlde)
	bus.Register(formServices.DeleteFormsFiledsEventHanlde)
	bus.Register(formServices.CreateFormsFiledsFailedEventHanlde)
	bus.Register(formServices.UpdateFormsFieldsFailedEventHanlde)
	bus.Register(formServices.DeleteFormsFiledsFailedEventHanlde)
	bus.Register(formServices.CreateDependsOnEventHandle)
	bus.Register(formServices.UpdateDependsOnEventHandle)
	bus.Register(formServices.DeleteDependsOnEventHandle)

	// 路由分组
	v1 := r.Group("/v1")

	// 初始化表单处理器
	formsHandlers := http.NewFormsHandlers(formsAggregateRepo, bus, formServices)
	formsRouter := v1.Group("/forms")
	formsRouter.POST("/save", formsHandlers.SaveFormsHandler)
	formsRouter.POST("/delete", formsHandlers.DeleteFormsHandler)
	formsRouter.POST("/save-fields", formsHandlers.SaveFormsFieldsHandler)
	formsRouter.POST("/delete-fields", formsHandlers.DeleteFormsFiledsHandler)
	formsRouter.POST("/save-depends-on", formsHandlers.SaveDependsOnHandler)

	// 初始化权限处理器
	userPermissionsHandlers := http.NewUserPermissionsHandlers(
		permissions.NewPermissionsService(
			repository.NewPermissionsRepository(mysqlDb),
		),
	)


	// 权限路由
	permissionsRouter := v1.Group("/user-permissions")
	permissionsRouter.POST("/save", userPermissionsHandlers.Save)
	permissionsRouter.GET("/list", userPermissionsHandlers.List)

	// 角色路由
	roleHttpHandlers := http.NewRoleHandlers(repository.NewRoleAggregateRepository(mysqlDb), bus, roleServices)
	roleRouter := v1.Group("/role")
	roleRouter.POST("/save", roleHttpHandlers.Save)
	roleRouter.GET("/list", roleHttpHandlers.GetRoleList)
	roleRouter.DELETE("/delete/:id", roleHttpHandlers.DeleteRole)
	roleRouter.POST("/delete-permission", roleHttpHandlers.DeleteRolePermission)

	// 用户路由
	userHttpHandlers := http.NewUserHandlers(repository.NewUserAggregateRepository(mysqlDb, redisDb, true), bus, userServices)
	userRouter := v1.Group("/user")
	userRouter.POST("/save", userHttpHandlers.Save)
	userRouter.GET("/list", userHttpHandlers.GetUserList)
	userRouter.DELETE("/delete/:id", userHttpHandlers.DeleteUser)
	userRouter.POST("/status", userHttpHandlers.SaveUserStatus)
	userRouter.POST("/associated-roles", userHttpHandlers.AssoicatedRoles)
	userRouter.POST("/register", userHttpHandlers.RegisterUserHandler)
	userRouter.POST("/login", userHttpHandlers.LoginUserHandler)
	userRouter.POST("/logout", userHttpHandlers.LogoutUserHandler)

	// 发送验证码
	userRouter.POST("/send-verify-code", userHttpHandlers.SendVerifyCodeHandler)
	userRouter.POST("/refresh-token", userHttpHandlers.RefreshTokenHandler)

	// 启动 Gin
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}