package utils

import (
	"url-shortener/models"

	"go.osspkg.com/x/algorithms/encoding/base62"
	"gorm.io/gorm"
)

func CreatShortID(id uint64) string{
    
	// Definimos el codificador base 62
	encoder := base62.New("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	// Codificamos el id
	encodeID := encoder.Encode(id)
    
	// Devolvemos
	return encodeID
}

// Busca la url y aumenta el numero de clicks
func ManageVisit(short_url string, db *gorm.DB) (string,error) {
    
	// Defimos el modelo
	var url models.URLTable
    
	tx := db.Begin()

	if tx.Error != nil{
		return "",tx.Error
	}

	// Buscamos la url
	if err := tx.Where("short_url = ?",short_url).First(&url).Error; err != nil{
		tx.Rollback()
		return "", err
	}

	// Aumentamos los likes
	if err := tx.Model(&url).
			Update("clicks", gorm.Expr("clicks + ?",1)).Error; err != nil{
				tx.Rollback()
				return "",err
			}
    
	// Verificamos el estado
	if err := tx.Commit().Error; err != nil {
			return "", err
	}

	return  url.OriginalURL,nil
}