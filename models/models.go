package models

import (
	"gorm.io/gorm"
)

// Modelo de json que usara nuestra api
type JsonURLInfo struct {
	OriginalURL string `json:"original_url" validate:"required,url"`
}

// Modelo de la base de datos
type URLTable struct {
	// Definimos el modelo de gorm para que cree los campos de id y fechas de manipulacion
	gorm.Model
	OriginalURL string `gorm:"not null"`
	ShortURL string `gorm:"uniqueIndex; not null"`
	Clicks int `gorm:"default:0"`
}

// Configuramos el nombre de la tabla
func (URLTable) TableName() string {
	return "urls" // nombre de la tabla
}
