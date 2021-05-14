package usuusuariosadapters

import (
	"database/sql"
	lealdb "leal.co/listas-aggregator/src/common/leal-db"
	usuusuarios "leal.co/listas-aggregator/src/usu_usuarios/domain"
	"log"
)

type UsuUsuariosSQLAdapter interface {
	TraerInfoExtra(*usuusuarios.DtoSalida) error
}

type usuUsuariosSQLAdapter struct {
	db *sql.DB
}

func NewUsuUsuariosSqlAdapter() UsuUsuariosSQLAdapter {
	return &usuUsuariosSQLAdapter{
		db: lealdb.GetDB(),
	}
}

func (u *usuUsuariosSQLAdapter) TraerInfoExtra(usu *usuusuarios.DtoSalida) error {
	idComercios := make([]int, 0, 3)
	rows, err := u.db.Query(
		`SELECT id_comercio FROM usu_usuarios_comercios uuc where uid = ?`, usu.Uid)

	if err != nil {
		log.Printf("Error buscando comercios del usuario: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		idComercios = append(idComercios, id)
	}

	usu.IdsComercios = idComercios

	return nil
}
