package main

import (
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/db"
	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
  // SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
  FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
  g := gen.NewGenerator(gen.Config{
    OutPath: "internal/infrastructure/dao/query",
    Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
  })

  gormdb, _ := db.GetDB()
  g.UseDB(gormdb) // reuse your gorm db

  // Generate basic type-safe DAO API for struct `model.User` following conventions
  g.ApplyBasic(model.UserPermission{})

  // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
  g.ApplyInterface(func(Querier){}, model.UserPermission{})

  // Generate the code
  g.Execute()
}