package main

import (
	"Online-Test-GO/internal/pkg/config"
	"Online-Test-GO/internal/pkg/repository/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 16 << 20

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	postgres.NewPostgres()

	//repository
	//userRepo := users.NewRepository(postgresDB)
	//taskRepo := tasks.NewRepository(postgresDB)
	//projectRepo := projects.NewRepository(postgresDB)
	//
	//////service
	//userService := users2.NewService(userRepo)
	//tasksService := tasks2.NewService(taskRepo)
	//projectsService := projects2.NewService(projectRepo)
	//
	//userUseCase := users3.NewUseCase(userService, userRepo)
	//taskUseCaseService := tasks3.NewUseCase(tasksService)
	//projectsUseCaseService := projects3.NewUseCase(projectsService)
	//
	//userController := users4.NewController(userUseCase)
	//taskController := tasks4.NewController(taskUseCaseService)
	//projectsController := projects4.NewController(projectsUseCaseService)
	//exportController := export4.NewController(userUseCase, taskUseCaseService, projectsUseCaseService)

	api := r.Group("api")
	{
		v1 := api.Group("v1")

		v1.GET("/time", func(c *gin.Context) {
			now := time.Now()
			c.JSON(http.StatusOK, gin.H{
				"message": "ok!",
				"status":  true,
				"data": map[string]interface{}{
					"time":            now.Format("15:04"),
					"time_in_seconds": now.Hour()*3600 + now.Minute()*60 + now.Second(),
					"unix":            now.Unix(),
					"date":            now.Format("02.01.2006"),
					"week_day":        now.Weekday(),
					"full_date":       now.Format("02.01.2006 15:04:06"),
					"month":           now.Month(),
					"day":             now.Day(),
					"year":            now.Year(),
					"hour":            now.Hour(),
					"minute":          now.Minute(),
					"second":          now.Second(),
				},
			})
		})
		// @router
		//{
		//	user_router.Router(v1, userController)
		//	task_router.Router(v1, taskController)
		//	project_router.Router(v1, projectsController)
		//	export.Router(v1, exportController)
		//}
	}

	log.Fatalln(r.Run(":" + config.GetConf().Port))
}
