package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/app/services"
	"github.com/r1005410078/meida-admin-server/internal/domain/permissions"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/db"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/repository"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/http"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)

func main() {
	// 初始化 Gin 路由
	r := gin.Default()

	// 根据环境变量设置 Gin 模式，默认是 debug 模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}



	mysqlDb, err := db.GetDB()
	if err != nil {
		panic(err)}
	
	// 初始化角色服务
  roleRepo := repository.NewRoleRepository(mysqlDb)
	roleServices := services.NewRepoServices(roleRepo)

	// 注册事件
	bus := shared.NewEventBus()
	bus.Register(roleServices.DeleteRoleEventHandle)
	bus.Register(roleServices.RoleDeleteFailedEventHandle)
	bus.Register(roleServices.SaveRoleEventHandle)
	bus.Register(roleServices.RoleSaveFailedEventHandle)

	
	userPermissionsHandlers := http.NewUserPermissionsHandlers(permissions.NewPermissionsService(repository.NewPermissionsRepository(mysqlDb)))

	v1 := r.Group("/v1")
	permissionsRouter := v1.Group("/user-permissions")
	permissionsRouter.POST("/save", userPermissionsHandlers.Save)
	permissionsRouter.GET("/list", userPermissionsHandlers.List)


	roleHttpHandlers := http.NewRoleHandlers(repository.NewRoleAggregateRepository(mysqlDb), bus, roleServices)
	roleRouter := v1.Group("/role")
	roleRouter.POST("/save", roleHttpHandlers.Save)
	roleRouter.GET("/list", roleHttpHandlers.GetRoleList)

	// 启动 Gin
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}