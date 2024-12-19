package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/r1005410078/meida-admin-server/internal/domain/permissions"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/db"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/repository"
	"github.com/r1005410078/meida-admin-server/internal/interfaces/http"
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
	
	userPermissionsHandlers := http.NewUserPermissionsHandlers(permissions.NewPermissionsService(repository.NewPermissionsRepository(mysqlDb)))

	v1 := r.Group("/v1")
	permissionsRouter := v1.Group("/user-permissions")
	permissionsRouter.POST("/save", userPermissionsHandlers.Save)
	permissionsRouter.GET("/list", userPermissionsHandlers.List)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}