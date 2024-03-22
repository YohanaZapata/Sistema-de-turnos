package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/YohanaZapata/Sistema-de-turnos/internal/domain"
	"github.com/YohanaZapata/Sistema-de-turnos/pkg/store"
)

func main() {
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	store := store.NewSqlStore(db)

	dentistService := domain.NewService(store)
	dentistHandler := domain.NewHandler(dentistService)

	r := gin.Default()

	r.GET("/ping/:time", func(c *gin.Context) {
		tm := c.Param("time")
		c.String(200, "pong after "+tm+" second(s)")
	})

	r.POST("/dentists", dentistHandler.CreateDentist)

	dentistsGroup := r.Group("/dentists")
	{
		dentistsGroup.GET("/:id", dentistHandler.GetDentistByID)
		dentistsGroup.GET("", dentistHandler.GetAllDentists)
		dentistsGroup.PUT("/:id", dentistHandler.UpdateDentist)
		dentistsGroup.DELETE("/:id", dentistHandler.DeleteDentist)
	}

	err = r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
