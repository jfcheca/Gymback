package usuarios

import (
	"errors"

	"github.com/jfcheca/Gymback/internal/domain"
)

type Repository interface {

	CrearUsuario(p domain.Usuarios) (domain.Usuarios, error)

}

type repository struct {
	storage  store.StoreInterfaceUsuarios

}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceUsuarios) Repository {
	return &repository{storage}
}


//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearUsuario(p domain.Usuarios) (domain.Usuarios, error) {
    err := r.storage.CrearUsuario(p)
    if err != nil {
        return domain.Usuarios{}, errors.New("error creating usuario")
    }
    return p, nil
}