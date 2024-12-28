## 创建表

migrate create -ext sql -dir internal/infrastructure/db/migrations create_users

## 生成表

migrate -path internal/infrastructure/db/migrations -database "mysql://root:123456@tcp(127.0.0.1:3306)/meida_dev" up

## 生成模型

gentool -c "./gen.tool"
go run ./internal/infrastructure/cmd/main.go
