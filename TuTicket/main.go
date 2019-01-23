package main

import (
    "net/http"
    "strconv"

    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
    //open a db connection
    var err error
    db, err = gorm.Open("mysql", "root:root@/tu_ticket?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("failed to connect database")
    }

    //Migrate the schema
    db.AutoMigrate(&evento{})
}

func main() {

    router := gin.Default()

    router.LoadHTMLGlob("templates/*")
    router.Static("/js", "./static/js")

    /*  router.GET("/eventos", viewEventos)
    router.GET("/eventos/:evento_id", viewEvento)

    //Redirecciona a la vista eventos como página principal
    router.GET("/", func(c *gin.Context){
        c.Redirect(http.StatusMovedPermanently, "/eventos")
    })*/

    api := router.Group("/api")
    {
        api.GET("/eventos", getAllEventos)
        api.GET("/eventos/:evento_id", getSingleEvento)
        api.GET("/evento/:evento_id/:asiento_id", getAsiento)
        api.POST("/evento/:evento_id/:asiento_id/comprar", buyBoleto)
    }

    router.Run()
}