package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/YohanaZapata/Sistema-de-turnos/internal/domain"
	"github.com/YohanaZapata/Sistema-de-turnos/pkg/store"
)

func main() {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
	if err != nil {
		log.Fatal(err)
	}

	dentistRepository := store.NewSqlStore(db)
	dentistService := domain.NewService(dentistRepository)
	dentistHandler := domain.NewHandler(dentistService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	dentistRoutes := r.Group("/dentistas")
	{
		dentistRoutes.GET("/:id", dentistHandler.GetDentistaByID)
		dentistRoutes.POST("", dentistHandler.CreateDentista)
		dentistRoutes.DELETE("/:id", dentistHandler.DeleteDentista)
		dentistRoutes.PATCH("/:id", dentistHandler.UpdateDentista)
		dentistRoutes.PUT("/:id", dentistHandler.UpdateDentista)
		dentistRoutes.GET("/", dentistHandler.GetAllDentistas)
		dentistRoutes.GET("/code/:codeValue", dentistHandler.GetDentistaByCodeValue)
		dentistRoutes.GET("/published", dentistHandler.GetDentistasByPublished)
		dentistRoutes.GET("/expiration/:expiration", dentistHandler.GetDentistasByExpiration)
		dentistRoutes.GET("/price/:price", dentistHandler.GetDentistasByPrice)
		dentistRoutes.GET("/quantity/:quantity", dentistHandler.GetDentistasByQuantity)
		dentistRoutes.GET("/matricula/:matricula", dentistHandler.GetDentistasByMatricula)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
