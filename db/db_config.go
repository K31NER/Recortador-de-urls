package db

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Definimos la conxion global
var DB *gorm.DB

// Creamos la funcion de conexion
func Connect() (*gorm.DB, error){
	var err error // Definimos la variable para capturar errores
    
	// Creamos la conexion
	DB, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
    
	// Validamos si no hubo error
	if err != nil{
		log.Printf("No se puedo conectar con la base de datos: %v",err)
		return nil, err
	}
    
	log.Println("Conectado con la base de datos")
	return DB, nil
}
