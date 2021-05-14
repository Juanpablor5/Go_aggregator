package usuusuarioscomerciosadapters

import (
	"database/sql"
	"fmt"
	lealdb "leal.co/listas-aggregator/src/common/leal-db"
	domain "leal.co/listas-aggregator/src/usu_usuarios_comercios/domain"
)

type UsuUsuariosComerciosSQLAdapter interface {
	TraerInfoExtra(*domain.DtoSalida) error
}

type usuUsuariosComerciosSQLAdapter struct {
	db *sql.DB
}

func NewUsuUsuariosComerciosSQLAdapter() UsuUsuariosComerciosSQLAdapter {
	return &usuUsuariosComerciosSQLAdapter{
		db: lealdb.GetDB(),
	}
}

func (u *usuUsuariosComerciosSQLAdapter) TraerInfoExtra(usu *domain.DtoSalida) error {
	//"premios_por_vencer" : 0, COALESCE(IF(t13.id>0,1,0),0) black_list, IF(t3.puntos_activos> MIN(t12.puntos) AND t3.puntos_activos-t3.puntos_por_vencer < MIN(t12.puntos) , 1,0)
	//"estadoUsuario" : "saas", IF((SELECT count(tt1.uid) FROM leal.usu_transacciones_coins tt1 WHERE t1.uid = tt1.uid) = 0, 'saas', 'ambos')
	err := u.db.QueryRow(`SELECT
							t1.uid,
							t1.cedula,
							t1.tipo_documento id_tipo_documento,
							t4.tipo_documento,
							t1.nombre,
							t1.apellido,
							t1.fullname,
							t1.email,
							t1.celular,
							t1.genero,
							DATE(t1.fecha_registro) fecha_registro_leal,
							t1.cumpleanos,
							COALESCE(MONTH(t1.cumpleanos), 0) cumpleanos_mes,
							COALESCE(DAY(t1.cumpleanos), 0) cumpleanos_dia,
							leal.fn_agrupar_edad(t1.cumpleanos,1) grupo_edad,
							IF(FLOOR(((TO_DAYS(CURDATE()) - TO_DAYS(t1.cumpleanos)) / 365.2422)) < 97,FLOOR(((TO_DAYS(CURDATE()) - TO_DAYS(t1.cumpleanos)) / 365.2422)),0) AS edad,
							t1.estado estado_leal,
							COALESCE(t1.plataforma_movil, ''),
							IF(t1.envio_mail = 'activo',1,0) envio_mail_global,
							IF(t1.envio_sms = 'activo',1,0) envio_sms_global,
							IF(t1.enviar_push = 'activo',1,0) envio_push_global,
							t1.fecha_perfil_completado,
							t1.cod_ciudad id_ciudad,
							t5.ciudad,
							t1.cod_pais,
							t5.id_departamento,
							COALESCE(t10.estatus, ''),
							COALESCE(COUNT(DISTINCT(IF(t3.puntos_activos >= t12.puntos, t12.premio, NULL))), 0) premios_disponibles_cantidad,
							COALESCE(GROUP_CONCAT(DISTINCT(IF(t3.puntos_activos >= t12.puntos, t12.id_premio, NULL))), 0) id_premios_disponibles,
							COALESCE(GROUP_CONCAT(DISTINCT(IF(t3.puntos_activos >= t12.puntos, t12.premio, NULL))), '') premios_disponibles,
							COALESCE(IF(t13.id>0, 1, 0), 0) black_list,
							IF(t3.puntos_activos> MIN(t12.puntos) AND t3.puntos_activos-t3.puntos_por_vencer < MIN(t12.puntos),	1, 0) premios_por_vencer,
							IF((SELECT count(tt1.uid) FROM leal.usu_transacciones_coins tt1 WHERE t1.uid = tt1.uid) = 0, 'saas', 'ambos') estadoUsuario
						FROM
							leal.usu_usuarios t1
						INNER JOIN leal.usu_usuarios_comercios t3 ON
							t1.uid = t3.uid
						LEFT JOIN leal.cnf_tipo_documento t4 ON
							t1.tipo_documento = t4.id_tipo_identificacion
						LEFT JOIN leal.cnf_ciudades t5 ON
							t1.cod_ciudad = t5.id_ciudad
						LEFT JOIN leal.usu_estatus t10 ON
							t3.id_estatus = t10.id_estatus
						LEFT JOIN leal.com_premios t12 ON
							t3.id_comercio = t12.id_comercio
							AND t12.estado = 'activo'
						LEFT JOIN leal.adm_email_blacklist t13 ON
							t1.email = t13.email
						WHERE
							t1.uid = ?
							AND t3.id_comercio = ?
							AND t1.estado not in ('inactivo')
							AND t3.estado not like ('inacti%')`, usu.Uid, usu.IdComercio).
		Scan(&usu.Usuario.Uid, &usu.Usuario.Cedula, &usu.Usuario.IdTipoDocumento, &usu.TipoDocumento, &usu.Usuario.Nombre, &usu.Usuario.Apellido, &usu.Usuario.Fullname,
			&usu.Usuario.Email, &usu.Usuario.Celular, &usu.Usuario.Genero, &usu.Usuario.FechaRegistroLeal, &usu.Usuario.Cumpleanos, &usu.Usuario.CumpleanosMes, &usu.Usuario.CumpleanosDia,
			&usu.Usuario.GrupoEdad, &usu.Usuario.Edad, &usu.Usuario.EstadoLeal, &usu.Usuario.PlataformaMovil, &usu.Usuario.EnvioMailGlobal, &usu.Usuario.EnvioSmsGlobal,
			&usu.Usuario.EnvioPushGlobal, &usu.Usuario.FechaPerfilCompletado, &usu.Usuario.IdCiudad, &usu.Ciudad, &usu.Usuario.CodPais, &usu.IdDepartamento, &usu.Estatus,
			&usu.PremiosDisponiblesCantidad, &usu.IdPremiosDisponibles, &usu.PremiosDisponibles, &usu.BlackList, &usu.PremiosPorVencer,
			&usu.Usuario.EstadoUsuario)
	if err != nil {
		return fmt.Errorf("error getting extra info from saas, %v", err)
	}

	return nil

}
