package main

import (
	"log"
	"os"
	"swiping-card/controller"
	"swiping-card/lib"
	"swiping-card/middle"
	"swiping-card/model"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.POST("/api/admin/login", controller.AdminLoginPost)
	r.POST("/api/user/login", controller.UserLogin)
	r.POST("/api/test/login", controller.TestLogin)
	r.GET("/api/banks", controller.BankGet)
	r.GET("/api/ways", controller.WayGet)
	r.GET("/api/cycles", controller.CycleGet)
	login := r.Group("/api")
	login.Use(middle.IsLogin())
	{
		login.GET("/cards", controller.CardGet)
		login.POST("/cards", controller.CardPost)
		login.POST("/plans", controller.PlanPost)
		login.GET("/plans", controller.PlanGet)
		login.GET("/dos", controller.DoGet)
		login.PUT("/dos/:id", controller.DoPut)
		admin := login.Group("")
		admin.Use(middle.IsAdmin())
		{
			admin.POST("/banks", controller.BankPost)
			admin.POST("/ways", controller.WayPost)
			admin.POST("/cycles", controller.CyclePost)
		}
	}

	return r
}

func dbinstance() {
	db := lib.DB()
	db.AutoMigrate(
		&model.Bank{},
		&model.Card{},
		&model.User{},
		&model.Way{},
		&model.Admin{},
		&model.Cycle{},
		&model.Plan{},
		&model.Do{},
	)
}

func main() {
	c := cron.New(cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	// c.AddFunc("* * * * *", func() {
	c.AddFunc("@daily", func() {
		if err := controller.DoPost(); err != nil {
			log.Panicln(err.Error())
		}
	})
	c.Start()
	defer c.Stop()

	r := setupRouter()
	dbinstance()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
