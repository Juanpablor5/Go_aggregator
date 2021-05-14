package usuusuariosservice

import (
	usuusuariosadapters "leal.co/listas-aggregator/src/usu_usuarios/adapters"
	usuusuarios "leal.co/listas-aggregator/src/usu_usuarios/domain"
	"log"
)

type UsuUsuariosService interface {
	CreateUsuario(dto usuusuarios.DtoIngreso)
	UpdateUsuario(dto usuusuarios.DtoIngreso)
}

func NewUsuUsuariosService() UsuUsuariosService {
	return &usuUsuariosService{
		kinesisAdapter: usuusuariosadapters.NewUsuUsuariosKinesisAdapter(),
		sqlAdapter:     usuusuariosadapters.NewUsuUsuariosSqlAdapter(),
	}
}

type usuUsuariosService struct {
	kinesisAdapter usuusuariosadapters.UsuUsuariosKinsisAdapter
	sqlAdapter     usuusuariosadapters.UsuUsuariosSQLAdapter
}

func (u *usuUsuariosService) CreateUsuario(dto usuusuarios.DtoIngreso) {
	dtoSalida := dto.ToDtoSalida()

	err := u.sqlAdapter.TraerInfoExtra(&dtoSalida)
	if err != nil {
		log.Println(err)
	}

	u.kinesisAdapter.PutRecord(dtoSalida)
}

func (u *usuUsuariosService) UpdateUsuario(dto usuusuarios.DtoIngreso) {
	dtoSalida := dto.ToDtoSalida()

	err := u.sqlAdapter.TraerInfoExtra(&dtoSalida)
	if err != nil {
		log.Println(err)
	}

	u.kinesisAdapter.PutRecord(dtoSalida)
}
