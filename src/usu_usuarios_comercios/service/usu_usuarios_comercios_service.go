package usu_usuarios_comercios_service

import (
	usu_usuarios_comercios_adapters "leal.co/listas-aggregator/src/usu_usuarios_comercios/adapters"
	usu_usuarios_comercios "leal.co/listas-aggregator/src/usu_usuarios_comercios/domain"
	"log"
)

type UsuUsuariosComerciosService interface {
	CreateUsuUsuariosComercios(dto usu_usuarios_comercios.DtoIngreso)
	UpdateUsuUsuariosComercios(dto usu_usuarios_comercios.DtoIngreso)
}

func NewUsuUsuariosComerciosService() UsuUsuariosComerciosService {
	return &usuUsuariosComerciosService{
		kinesisAdapter: usu_usuarios_comercios_adapters.NewUsuUsuariosComerciosKinesisAdapter(),
		sqlAdapter:     usu_usuarios_comercios_adapters.NewUsuUsuariosComerciosSQLAdapter(),
	}
}

type usuUsuariosComerciosService struct {
	kinesisAdapter usu_usuarios_comercios_adapters.UsuUsuariosComerciosKinesisAdapter
	sqlAdapter     usu_usuarios_comercios_adapters.UsuUsuariosComerciosSQLAdapter
}

func (u *usuUsuariosComerciosService) CreateUsuUsuariosComercios(dto usu_usuarios_comercios.DtoIngreso) {
	dtoSalida := dto.ToDtoSalida()
	if err := u.sqlAdapter.TraerInfoExtra(&dtoSalida); err != nil {
		log.Println(err)
	}
	u.kinesisAdapter.PutRecord(dtoSalida)
}

func (u *usuUsuariosComerciosService) UpdateUsuUsuariosComercios(dto usu_usuarios_comercios.DtoIngreso) {
	dtoSalida := dto.ToDtoSalida()
	if err := u.sqlAdapter.TraerInfoExtra(&dtoSalida); err != nil {
		log.Println(err)
	}
	u.kinesisAdapter.PutRecord(dtoSalida)
}
