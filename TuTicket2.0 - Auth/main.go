package main

import (
    "net/http"
	"fmt"
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
	db.AutoMigrate(&UsuarioModel{})

}

func main() {

    router := gin.Default()

    //Rutas para front-end
    /*router.LoadHTMLGlob("templates/*")
    router.Static("/js", "./static/js")

    //Redirecciona a la vista eventos como página principal
    router.GET("/", func(c *gin.Context){
    c.Redirect(http.StatusMovedPermanently, "/api/auth")
    })*/

    //Rutas para api
    api := router.Group("/api")
    {
        api.POST("/auth", Auth)
    }

    router.Run()
}


//Obtiene un solo evento
func Auth(c *gin.Context) {
	

	//Encuentra el usuario con el Usuario que se envió como parámetro
	var user UsuarioModel
  	usuario := c.PostForm("usuario")
	clave := c.PostForm("contrasena")
	db.Where("Usuario = ?", usuario).First(&user)
	//db.First(&user, usuario)
	
	fmt.Println(user.Usuario)
	fmt.Println(usuario)
	fmt.Println(user.Contrasena)
	fmt.Println(clave)
	
	
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Usuario no existe"})
		return
	}
	
	if(user.Usuario == usuario && user.Contrasena == clave){
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Ingreso Exitoso"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Usuario/Contraseña Incorrectos"})
}

type (
	UsuarioModel struct {
		gorm.Model
		Usuario    string      `json:"usuario"`
		Contrasena string      `json:"contrasena"`
	}
)
