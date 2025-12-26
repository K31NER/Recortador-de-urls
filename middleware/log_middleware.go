package middleware

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

// Definimos el middleware para imprimir las solicitudes recibidas
func LoggerMiddleware(c *fiber.Ctx) error {
	// Ejecutamos el siguiente manejador
	err := c.Next()

	// Por defecto el código es el de la respuesta
	code := c.Response().StatusCode()

	// Si hubo un error (como un 404 o 405), Fiber lo devuelve en 'err'
	// pero aún no lo ha puesto en la respuesta. Debemos extraerlo:
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
	}

	// Log a mostrar con el código correcto
	log.Printf("Solicitud recibida: %s %s - %d", c.Method(), c.Path(), code)

	return err
}
