package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&CompanyModel{},
		&NodeModel{},
		&GameModel{},
		&EdgeModel{})

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	// move this to an interface approach
	companyController := NewCompanyController("/v1")
	nodeController := NewNodeController("/v1")
	gameController := NewGameController("/v1")
	edgeController := NewEdgeController("/v1")

	companyController.Router(r)
	nodeController.Router(r)
	gameController.Router(r)
	edgeController.Router(r)
	r.Run()
}
