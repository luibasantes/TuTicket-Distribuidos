package main

import (
    //"net/http"
    //"strconv"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//Conexión a la base de datos MySQL
func init() {
    //Open a db connection
    var err error
    db, err = gorm.Open("mysql", "root:root@/tu_ticket?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("Conexión con la base de datos fallida")
    }

    //Migrate the schema
    db.AutoMigrate(&EventoModel{})

    //Para utilizar la variable db en otros scripts
    return func(c *gin.Context) {
        c.Set("DB", db)
        c.Next()
    }
}

func main() {

    router := gin.Default()

    //Rutas para front-end
    /*router.LoadHTMLGlob("templates/*")
    router.Static("/js", "./static/js")

    router.GET("/eventos", viewEventos)
    router.GET("/eventos/:evento_id", viewEvento)

    //Redirecciona a la vista eventos como página principal
    router.GET("/", func(c *gin.Context){
    c.Redirect(http.StatusMovedPermanently, "/eventos")
    })*/

    //Rutas para api
    api := router.Group("/api")
    {
        api.GET("/eventos", getAllEventos)
        api.GET("/eventos/:evento_id", getSingleEvento)
        api.GET("/eventos/:evento_id/asientos", getAsientos)
        api.POST("/comprar", buyBoleto)
    }

    router.Run()
}