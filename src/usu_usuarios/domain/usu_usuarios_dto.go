package usuusuarios

type DtoIngreso struct {
	Uid                   string `json:"uid"`
	Cedula                string `json:"cedula"`
	Fullname              string `json:"fullname"`
	Nombre                string `json:"nombre"`
	Apellido              string `json:"apellido"`
	Email                 string `json:"email"`
	Contrasena            string `json:"contrasena"`
	Celular               string `json:"celular"`
	Genero                string `json:"genero"`
	FechaRegistro         string `json:"fecha_registro"`
	CodCiudad             int    `json:"cod_ciudad"`
	Cumpleanos            string `json:"cumpleanos"`
	Estado                string `json:"estado"`
	Foto                  string `json:"foto"`
	PlataformaOrigen      string `json:"plataforma_origen"`
	TipoCuenta            string `json:"tipo_cuenta"`
	Token                 string `json:"token"`
	PlataformaMovil       string `json:"plataforma_movil"`
	UidCms                string `json:"uid_cms"`
	Nota                  string `json:"nota"`
	EnvioMail             string `json:"envio_email"`
	EnvioSms              string `json:"envio_sms"`
	IpCajero              string `json:"ip_cajero"`
	CodPais               string `json:"cod_pais"`
	ComerciosInscritos    int    `json:"comercios_inscritos"`
	Visitas               int    `json:"visitas"`
	EnviarPush            string `json:"enviar_push"`
	FechaIngresoApp       string `json:"fecha_ingreso_app"`
	FechaPerfilCompletado string `json:"fecha_perfil_completado"`
	TipoDocumento         int    `json:"tipo_documento"`
	Intentos              int    `json:"intentos"`
	FechaIntentoIngreso   string `json:"fecha_intento_ingreso"`
	CoinsTest             int    `json:"coins_test"`
	IdSucursalRegistro    int    `json:"id_sucursal_registro"`
	DeviceId              string `json:"device_id"`
	CedulaWl              string `json:"cedula_wl"` // cedula white label (wl)
}

func (d *DtoIngreso) ToDtoSalida() DtoSalida {
	return DtoSalida{
		Uid:                   d.Uid,
		Cedula:                d.Uid,
		IdTipoDocumento:       d.TipoDocumento,
		Nombre:                d.Nombre,
		Apellido:              d.Apellido,
		Fullname:              d.Fullname,
		Email:                 d.Email,
		Celular:               d.Celular,
		Genero:                d.Genero,
		FechaRegistroLeal:     d.FechaRegistro,
		Cumpleanos:            d.Cumpleanos,
		EstadoLeal:            d.Estado,
		PlataformaMovil:       d.PlataformaMovil,
		EnvioMailGlobal:       boolToBinary(d.EnvioMail, "activo"),
		EnvioSmsGlobal:        boolToBinary(d.EnvioSms, "activo"),
		EnvioPushGlobal:       boolToBinary(d.EnviarPush, "activo"),
		FechaPerfilCompletado: d.FechaPerfilCompletado,
		IdCiudad:              d.CodCiudad,
		CodPais:               d.CodPais,
		IdsComercios:          make([]int, 0, 3),
	}
}

func boolToBinary(value string, condition string) int {
	if value == condition {
		return 1
	}
	return 0
}

type DtoSalida struct {
	Uid                   string `json:"uid"`
	Cedula                string `json:"cedula"`
	IdTipoDocumento       int    `json:"id_tipo_documento"`
	Nombre                string `json:"nombre"`
	Apellido              string `json:"apellido"`
	Fullname              string `json:"fullname"`
	Email                 string `json:"email"`
	Celular               string `json:"celular"`
	Genero                string `json:"genero"`
	FechaRegistroLeal     string `json:"fecha_registro_leal"`
	Cumpleanos            string `json:"cumpleanos"`
	CumpleanosMes         int    `json:"cumpleanos_mes"`
	CumpleanosDia         int    `json:"cumpleanos_dia"`
	GrupoEdad             string `json:"grupo_edad"`
	Edad                  int    `json:"edad"`
	EstadoLeal            string `json:"estado_leal"`
	PlataformaMovil       string `json:"plataforma_movil"`
	EnvioMailGlobal       int    `json:"envio_mail_global"`
	EnvioSmsGlobal        int    `json:"envio_sms_global"`
	EnvioPushGlobal       int    `json:"envio_push_global"`
	FechaPerfilCompletado string `json:"fecha_perfil_completado"`
	IdCiudad              int    `json:"id_ciudad"`
	CodPais               string `json:"cod_pais"`
	EstadoUsuario         string `json:"estadoUsuario"`
	IdsComercios          []int  `json:"ids_comercios"`
}
