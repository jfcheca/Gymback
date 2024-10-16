package handler

import (
	"errors"
	"gimnasio/internal/domain"
	"gimnasio/internal/usuarios"
	"gimnasio/pkg/web"
	"github.com/gin-gonic/gin"
)

type usuariosHandler struct {
	s usuarios.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewUsuariosHandler(s usuarios.Service) *usuariosHandler {
	return &usuariosHandler{
		s: s,
	}
}

var listaOdontologo []domain.Usuarios
var ultimoID int = 1

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA NUEVO ODONTOLOGO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *usuariosHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odonto domain.Usuarios
		odonto.ID = ultimoID
		ultimoID++
		err := c.ShouldBindJSON(&odonto)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		// Crear el odontólogo utilizando el servicio
		createdOdonto, err := h.s.CrearUsuario(odonto)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create odontologo"))
			return
		}
		// Devolver el odontólogo creado con su ID asignado a la base de datos
		c.JSON(200, createdOdonto)
	}

}