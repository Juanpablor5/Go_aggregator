package usu_usuarios_comercios

import (
	usuusuarios "leal.co/listas-aggregator/src/usu_usuarios/domain"
)

type DtoIngreso struct {
	Uid                 string `json:"uid"`
	IdComercio          int    `json:"id_comercio"`
	UltimaActualizacion string `json:"ultima_actualizacion"`
	PuntosActivos       int    `json:"puntos_activos"`
	PuntosVencidos      int    `json:"puntos_vencidos"`
	PuntosUsados        int    `json:"puntos_usados"`
	PuntosTotales       int    `json:"puntos_totales"`
	Estado              string `json:"estado"`
	FechaRegistro       string `json:"fecha_registro"`
	Email               string `json:"email"`
	Sms                 string `json:"sms"`
	IdEstatus           int    `json:"id_estatus"`
	Nota                string `json:"nota"`
	Regalos             int    `json:"regalos"`
	EnviarPush          string `json:"enviar_push"`
	Editado             int    `json:"editado"`
	FplusmRank          int    `json:"FplusM_Rank"`
	RRank               int    `json:"R_rank"`
	Nps                 int    `json:"nps"`
	PuntosPorVencer     int    `json:"puntos_por_vencer"`
	TipoUsuario         int    `json:"tipo_usuario"`
	Tyc                 int    `json:"tyc"`
}

func (u *DtoIngreso) ToDtoSalida() DtoSalida {
	return DtoSalida{
		Uid:                   u.Uid,
		IdComercio:            u.IdComercio,
		FechaRegistroComercio: u.FechaRegistro,
		EstadoComercio:        u.Estado,
		EnvioMail:             boolToBinary(u.Email, "si"),
		EnvioSms:              boolToBinary(u.Sms, "si"),
		EnvioPushComercio:     boolToBinary(u.EnviarPush, "activo"),
		PuntosActivos:         u.PuntosActivos,
		PuntosVencidos:        u.PuntosVencidos,
		PuntosUsados:          u.PuntosUsados,
		PuntosTotales:         u.PuntosTotales,
		PuntosPorVencer:       u.PuntosPorVencer,
		IdEstatus:             u.IdEstatus,
		UltimaCalificacion:    handleNullInt(u.Nps, -1),
		Usuario:               usuusuarios.DtoSalida{IdsComercios: make([]int, 0, 0)},
	}
}

func boolToBinary(value string, condition string) int {
	if value == condition {
		return 1
	}
	return 0
}

func handleNullInt(value int, valueReturn int) int {
	if value != 0 {
		return value
	}
	return valueReturn
}

type DtoSalida struct {
	Uid                        string                `json:"uid"`
	IdComercio                 int                   `json:"id_comercio"`
	TipoDocumento              string                `json:"tipo_documento"`
	FechaRegistroComercio      string                `json:"fecha_registro_comercio"`
	EstadoComercio             string                `json:"estado_comercio"`
	EnvioMail                  int                   `json:"envio_mail"`
	EnvioSms                   int                   `json:"envio_sms"`
	EnvioPushComercio          int                   `json:"envio_push_comercio"`
	PuntosActivos              int                   `json:"puntos_activos"`
	PuntosVencidos             int                   `json:"puntos_vencidos"`
	PuntosUsados               int                   `json:"puntos_usados"`
	PuntosTotales              int                   `json:"puntos_totales"`
	PuntosPorVencer            int                   `json:"puntos_por_vencer"`
	IdEstatus                  int                   `json:"id_estatus"`
	Estatus                    string                `json:"estatus"`
	UltimaCalificacion         int                   `json:"ultima_calificacion"`
	IdPremiosDisponibles       string                `json:"id_premios_disponibles"`
	PremiosDisponibles         string                `json:"premios_disponibles"`
	PremiosPorVencer           int                   `json:"premios_por_vencer"`
	PremiosDisponiblesCantidad int                   `json:"premios_disponibles_cantidad"`
	BlackList                  string                `json:"black_list"`
	IdDepartamento             int                   `json:"id_departamento"`
	Ciudad                     string                `json:"ciudad"`
	Usuario                    usuusuarios.DtoSalida `json:"usuario"`
}
