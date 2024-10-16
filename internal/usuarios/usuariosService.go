package usuarios

import "github.com/jfcheca/Gymback/internal/domain"

type Service interface {

	CrearUsuario(p domain.Usuarios) (domain.Usuarios, error)

}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// CrearUsuario crea un nuevo usuario
func (s *service) CrearUsuario(p domain.Usuarios) (domain.Usuarios, error) {
    return s.r.CrearUsuario(p)
}