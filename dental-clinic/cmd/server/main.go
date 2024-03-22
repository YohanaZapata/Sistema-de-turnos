package main

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/dentist"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia") 
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	dentistStorage := store.NewSqlStore(db)
	dentistRepo := repository.NewDentistRepository(dentistStorage)
	dentistService := service.NewDentistService(dentistRepo)
	dentistHandler := handler.NewDentistHandler(dentistService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.POST("", dentistHandler.Create())
		dentists.DELETE(":id", dentistHandler.Delete())
		dentists.PATCH(":id", dentistHandler.Update())

	}
	r.Run(":8080")
}
