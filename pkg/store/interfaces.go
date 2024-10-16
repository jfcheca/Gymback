package store

import "github.com/jfcheca/Gymback/internal/domain"

type StoreInterfaceUsuarios interface {

	CrearUsuario(odonto domain.Odontologo) error

}