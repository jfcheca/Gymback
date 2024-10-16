package handler

import (
	"errors"
	"net/http"
//	"os"
	"strconv"
//	"strings"

	//	"text/template/parse"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/Checa_Budai_FinalBack3/internal/domain"
	"github.com/jfcheca/Checa_Budai_FinalBack3/internal/odontologo"
	"github.com/jfcheca/Checa_Budai_FinalBack3/pkg/web"
)

type odontoHandler struct {
	s odontologo.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s odontologo.Service) *odontoHandler {
	return &odontoHandler{
		s: s,
	}
}

var listaOdontologo []domain.Odontologo
var ultimoID int = 1

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA NUEVO ODONTOLOGO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *odontoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odonto domain.Odontologo
		odonto.ID = ultimoID
		ultimoID++
		err := c.ShouldBindJSON(&odonto)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		// Crear el odontólogo utilizando el servicio
		createdOdonto, err := h.s.Create(odonto)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create odontologo"))
			return
		}
		// Devolver el odontólogo creado con su ID asignado a la base de datos
		c.JSON(200, createdOdonto)
	}

}