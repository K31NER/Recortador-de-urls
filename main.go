package main

import (
	"log"
	"url-shortener/db"
	"url-shortener/models"
	"url-shortener/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Instanaciamos la app
	app := fiber.New()
    
	// Definimos el middleware de logging
	app.Use(logger.New())
    
	// Creamos la conexion con la base de datos
	conn, err := db.Connect()
    
	// Validamos que no ocurra un error
	if err != nil{
		log.Fatal("Error al conectar con base de datos")
	}

	// Realizamos la migracion
	if err := conn.AutoMigrate(&models.URLTable{}); err != nil {
	    log.Fatal("Error en migración:", err)
    }
	
	// Ruta principal
	app.Get("/",func (c *fiber.Ctx) error  {
		return c.JSON(fiber.Map{
			"message":"server is running",})
	})	
    
	// Configuramos las rutas
	routers.SetupRoutes(app,conn)

	// Arrancamos el servidor
	log.Fatal(app.Listen(":8080"))
}