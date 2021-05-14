package usu_historial_puntos

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type DtoIngreso struct {
	IdHistorial       int     `json:"id_historial"`
	Uid               string  `json:"uid"`
	Fecha             string  `json:"fecha"`
	FechaVencimiento  string  `json:"fecha_vencimiento"`
	Puntos            int     `json:"puntos"`
	BalancePuntos     int     `json:"balance_puntos"`
	IdEstadoPuntos    int     `json:"id_estado_puntos"`
	Tipo              string  `json:"tipo"`
	IdPremio          int     `json:"id_premio"`
	UidCms            string  `json:"uid_cms"`
	AccionRed         string  `json:"accion_red"`
	Valor             float64 `json:"valor"`
	IdComercio        int     `json:"id_comercio"`
	IdSucursal        int     `json:"id_sucursal"`
	IpCajero          string  `json:"ip_cajero"`
	Factura           string  `json:"factura"`
	Nota              string  `json:"nota"`
	IdFranquicia      int     `json:"id_franquicia"`
	IdTipoTransaccion int     `json:"id_tipo_transaccion"`
	CodigoBin         string  `json:"codigo_bin"`
}

func (d *DtoIngreso) ToDtoSalida() DtoSalida {
	return DtoSalida{
		IdHistorial:         d.IdHistorial,
		Uid:                 d.Uid,
		Fecha:               d.Fecha,
		FechaVencimiento:    d.FechaVencimiento,
		Puntos:              d.Puntos,
		BalancePuntos:       d.BalancePuntos,
		IdEstadoPuntos:      d.IdEstadoPuntos,
		Tipo:                d.Tipo,
		IdPremio:            d.IdPremio,
		UidCms:              d.UidCms,
		AccionRed:           d.AccionRed,
		Valor:               d.Valor,
		IdComercio:          d.IdComercio,
		IdSucursal:          d.IdSucursal,
		IpCajero:            d.IpCajero,
		Factura:             d.Factura,
		Nota:                d.Nota,
		IdFranquicia:        d.IdFranquicia,
		IdTipoTransaccion:   d.IdTipoTransaccion,
		CodigoBin:           d.CodigoBin,
		Productos:           make([]string, 0, 3),
		PremiosReclamados:   make([]string, 0, 3),
		Categorias:          make([]int, 0, 3),
		IdPremiosReclamados: make([]int, 0, 3),
		IdSucursales:        make([]int, 0, 3),
	}
}

type DtoSalida struct {
	IdHistorial               int      `json:"id_historial"`
	Uid                       string   `json:"uid"`
	Fecha                     string   `json:"fecha"`
	FechaVencimiento          string   `json:"fecha_vencimiento"`
	Puntos                    int      `json:"puntos"`
	BalancePuntos             int      `json:"balance_puntos"`
	IdEstadoPuntos            int      `json:"id_estado_puntos"`
	Tipo                      string   `json:"tipo"`
	IdPremio                  int      `json:"id_premio"`
	UidCms                    string   `json:"uid_cms"`
	AccionRed                 string   `json:"accion_red"`
	Valor                     float64  `json:"valor"`
	IdComercio                int      `json:"id_comercio"`
	IdSucursal                int      `json:"id_sucursal"`
	IpCajero                  string   `json:"ip_cajero"`
	Factura                   string   `json:"factura"`
	Nota                      string   `json:"nota"`
	IdFranquicia              int      `json:"id_franquicia"`
	IdTipoTransaccion         int      `json:"id_tipo_transaccion"`
	CodigoBin                 string   `json:"codigo_bin"`
	Frecuencia                float64  `json:"frecuencia"`
	TicketPromedio            float64  `json:"ticket_promedio"`
	GastoTotal                float64  `json:"gasto_total"`
	CantidadPremiosReclamados int      `json:"premios_reclamados_cantidad"`
	Transacciones             int      `json:"transacciones"`
	Banderazos                int      `json:"banderazos"`
	Cargas                    int      `json:"cargas"`
	Redenciones               int      `json:"redenciones"`
	Visitas                   int      `json:"visitas"`
	UltimaVisita              string   `json:"ultima_visita"`
	Etiquetas                 int      `json:"etiquetas"`
	Productos                 []string `json:"codigo_item"`
	Categorias                []int    `json:"categorias"`
	PremiosReclamados         []string `json:"premios_reclamdos"`
	IdPremiosReclamados       []int    `json:"id_premios_reclamados"`
	IdSucursales              []int    `json:"id_sucursales"`
}

func (d *DtoSalida) VerificarFactura() bool {
	factura, err := regexp.Match(`^U[0-9]+\-?`, []byte(d.Factura))
	if err != nil {
		log.Printf("error regex, %v", err)
		return false
	}

	return d.Tipo == "carga" && !factura && !strings.HasPrefix(d.Factura, fmt.Sprintf("%d-%d-", d.IdComercio, d.IdSucursal))
}
