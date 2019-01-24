package main

import (
    "net/http"
    "strconv"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//Conexión a la base de datos MySQL
func init() {
    //Open a db connection
    var err error
    db, err = gorm.Open("mysql", "root:luigibasantes1@/tu_ticket?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("Conexión con la base de datos fallida")
    }

    //Migrate the schema
    db.AutoMigrate(&EventoModel{})
	db.AutoMigrate(&SedeModel{})
	db.AutoMigrate(&AsientoModel{})
	db.AutoMigrate(&LocalidadModel{})
	db.AutoMigrate(&BoletoModel{})
	db.AutoMigrate(&UsuarioModel{})

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
		api.POST("/crear/evento", createEvento)
    }

    router.Run()
}


//Obtiene todos los eventos
func getAllEventos(c *gin.Context) {
	//Encuentra todos los eventos
	var eventos []EventoModel
	db.Find(&eventos)

	if len(eventos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Eventos no encontrados!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": eventos})
}

//Obtiene un solo evento
func getSingleEvento(c *gin.Context) {
	

	//Encuentra el evento con el id que se envió como parámetro
	var evento EventoModel
  	eventoID := c.Param("evento_id")
	db.First(&evento, eventoID)

	if evento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Evento no encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": evento})
}

//Obtiene todos los asientos de un evento
func getAsientos(c *gin.Context) {
	

	//Encuentra el evento con el id que se envió como parámetro
	var evento EventoModel
  	eventoID := c.Param("evento_id")
	db.First(&evento, eventoID)

	if evento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Evento no encontrado!"})
		return
	}

	//Encuentra los asientos del evento
	var asientos []AsientoModel
	db.Model(&evento).Related(&asientos);
	
	if len(asientos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡No hay asientos!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": asientos})
}

//Compra un boleto
func buyBoleto(c *gin.Context) {
	

	//Encuentra el asiento con el id que se envió en el POST
	var asiento AsientoModel
	asientoID, _ := strconv.Atoi(c.PostForm("asientoID"))
	//db.Where("id = ?", asientoID).First(&asiento)
	db.First(&asiento, asientoID)

	if asiento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Asiento no encontrado!"})
		return
	}

	//Valida que el asiento esté disponible
	if asiento.Disponible == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "message": "¡El asiento no está disponible!"})
		return
	}
	/*
	//Encuentra el usuario con el id que se envió en el POST
	var usuario UsuarioModel
  	usuarioID, _ := strconv.Atoi(c.PostForm("usuarioID"))
	db.Where("id = ?", usuarioID).First(&asiento)
	db.First(&usuario, usuarioID)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "¡Usuario no encontrado!"})
		return
	}
	
	//Guarda el boleto comprado
	var boleto BoletoModel
	boleto = BoletoModel{AsientoId: uint(asientoID), UsuarioId: uint(usuarioID)}
	db.Save(&boleto)
	*/
	//Actualiza la disponilidad del asiento
	db.Model(&asiento).Update("Disponible", false)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "¡Boleto comprado! ¡Gracias por su compra!"})
}

// createEvento add a new Evento
func createEvento(c *gin.Context) {
	agotado, _ := strconv.Atoi(c.PostForm("agotado"))
	agotado_transf := false
	if agotado == 1 {
		agotado_transf = true
	} else {
		agotado_transf = false
	}
	
	
	sedeId, _ := strconv.Atoi(c.PostForm("sede_id"))
	evento := EventoModel{Nombre: c.PostForm("nombre"), 
	Descripcion: c.PostForm("descripcion"), 
	Fecha: c.PostForm("fecha"),
	SedeId: sedeId,
	Agotado: agotado_transf}
	db.Save(&evento)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Evento item created successfully!", "resourceId": evento.ID})
}

	

type (
	EventoModel struct {
		gorm.Model
		Nombre      string      `json:"nombre"`
		Descripcion string      `json:"descripcion"`
		Fecha       string      `json:"fecha"`
		SedeId      int        `json:"sede_id"`
		Agotado     bool        `json:"agotado"`
		Asientos []AsientoModel
	}

	SedeModel struct {
		gorm.Model
		Nombre    string      `json:"nombre"`
		Tipo      string      `json:"tipo"`
		Ciudad    string      `json:"ciudad"`
		Pais      string      `json:"pais"`
		Direccion string      `json:"direccion"`
		Eventos []EventoModel
	}

	AsientoModel struct {
		gorm.Model
		Numero      uint      `json:"numero"`
		Disponible  bool      `json:"disponible"`
		EventoId    uint      `json:"evento_id"`
		LocalidadId uint      `json:"localidad_id"`
	}

	LocalidadModel struct {
		gorm.Model
		Nombre string      `json:"nombre"`
		Precio float64     `json:"precio"`
		Asientos []AsientoModel
	}

	BoletoModel struct {
		gorm.Model
		AsientoId uint      `json:"asiento_id"`
		UsuarioId uint      `json:"usuario_id"`
	}

	UsuarioModel struct {
		gorm.Model
		Usuario    string      `json:"usuario"`
		Nombre     string      `json:"nombre"`
		Contrasena string      `json:"contrasena"`
		Boletos []BoletoModel
	}
)
