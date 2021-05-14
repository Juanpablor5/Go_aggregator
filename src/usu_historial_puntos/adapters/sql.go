package usuhistorialpuntosadapters

import (
	"database/sql"
	"fmt"
	facturasdb "leal.co/listas-aggregator/src/common/facturas-db"
	lealdb "leal.co/listas-aggregator/src/common/leal-db"
	usuhistorialpuntos "leal.co/listas-aggregator/src/usu_historial_puntos/domain"
	"log"
	"strconv"
	"strings"
)

type UsuHistorialPuntosSQLAdapter interface {
	TraerInfoExtra(*usuhistorialpuntos.DtoSalida) error
	TraerInfoFacturas(*usuhistorialpuntos.DtoSalida) error
}

type usuHistorialPuntosSQLAdapter struct {
	db         *sql.DB
	facturasdb *sql.DB
}

func NewUsuHistorialPuntosSqlAdapter() UsuHistorialPuntosSQLAdapter {
	return &usuHistorialPuntosSQLAdapter{
		db:         lealdb.GetDB(),
		facturasdb: facturasdb.GetDB(),
	}
}

// trae info de bd y parte idPremios y saca unicos y parte premios y saca unicos
func (u *usuHistorialPuntosSQLAdapter) TraerInfoExtra(usu *usuhistorialpuntos.DtoSalida) error {
	var idPremios, premios, idSucursales string
	err := u.db.QueryRow(`
    SELECT
		COALESCE(AVG(IF(t2.tipo = 'carga', t2.valor, NULL)), 0) ticket_promedio,
		TIMESTAMPDIFF(second, min(t2.fecha), now())/(count(DISTINCT(IF(t2.tipo IN('carga', 'redencion'), t2.id_historial, NULL))))/(3600 * 24) frecuencia,
		COALESCE(( SELECT SUM(tt1.valor) from leal.usu_historial_puntos tt1 where tt1.uid = t1.uid and tt1.id_comercio = t2.id_comercio ), 0) gasto_total,
		COALESCE(GROUP_CONCAT((IF(t2.id_premio != 0, t2.id_premio, NULL)) ), '') id_premio_reclamados,
		COALESCE(GROUP_CONCAT((IF(t2.id_premio != 0, t11.premio, NULL))), '') premios_reclamados,
		COALESCE(COUNT(DISTINCT(IF(t2.id_premio != 0, t2.id_historial, NULL))), 0) premios_reclamados_cantidad,
		COALESCE(GROUP_CONCAT(DISTINCT(t2.id_sucursal)), '') id_sucursales,
		COUNT(DISTINCT(IF(t2.tipo IN('carga', 'redencion', 'promocion'), t2.id_historial, NULL))) transacciones,
		COUNT(DISTINCT(IF(t2.tipo = 'banderazo', t2.id_historial, NULL))) banderazos,
		COUNT(DISTINCT(IF(t2.tipo = 'carga' , t2.id_historial, NULL))) cargas,
		COUNT(DISTINCT(IF(t2.tipo = 'redencion', t2.id_historial, NULL))) redenciones,
		COUNT(DISTINCT(IF(t2.tipo IN('carga', 'redencion', 'promocion'), t2.id_historial, NULL))) visitas,
		DATE_FORMAT(MAX(IF(t2.tipo IN('carga', 'redencion', 'promocion'), DATE(t2.fecha), NULL)), '%Y-%m-%d') ultima_visita,
		COALESCE(t3.tipo_usuario, 0) as etiquetas
	FROM
		leal.usu_usuarios t1
	INNER JOIN leal.usu_usuarios_comercios t3 ON
		t1.uid = t3.uid
	LEFT JOIN leal.usu_historial_puntos t2 ON
		t2.uid = t3.uid
		AND t3.id_comercio = t2.id_comercio
	LEFT JOIN leal.com_premios t11 ON
		t2.id_premio = t11.id_premio
		AND t2.tipo = 'redencion'
	LEFT JOIN leal.com_calificaciones_usuarios t9 on
		t2.uid = t1.uid
		AND t2.id_historial = t9.id_historial
		AND t9.estado = 'calificado'
	WHERE
		t3.id_comercio = ?
		AND t1.uid = ?`, usu.IdComercio, usu.Uid).
		Scan(&usu.TicketPromedio, &usu.Frecuencia, &usu.GastoTotal, &idPremios, &premios, &usu.CantidadPremiosReclamados, &idSucursales, &usu.Transacciones,
			&usu.Banderazos, &usu.Cargas, &usu.Redenciones, &usu.Visitas, &usu.UltimaVisita, &usu.Etiquetas)

	if err != nil {
		return fmt.Errorf("error getting extra info from saas, %v", err)
	}

	idPremiostMap := make(map[int]bool)

	for _, prem := range strings.Split(idPremios, ",") {
		if prem != "" {
			i, err := strconv.Atoi(prem)
			if err != nil {
				log.Printf("error parsing premios: %s", err)
				continue
			}
			idPremiostMap[i] = true
		}
	}

	for prem, _ := range idPremiostMap {
		usu.IdPremiosReclamados = append(usu.IdPremiosReclamados, prem)
	}

	premiosMap := make(map[string]bool)

	for _, prem := range strings.Split(premios, ",") {
		if prem != "" {
			premiosMap[prem] = true
		}
	}

	for prem, _ := range premiosMap {
		usu.PremiosReclamados = append(usu.PremiosReclamados, prem)
	}

	idSucursalesMap := make(map[int]bool)

	for _, suc := range strings.Split(idSucursales, ",") {
		if suc != "" {
			i, err := strconv.Atoi(suc)
			if err != nil {
				log.Printf("error parsing sucursales: %s", err)
				continue
			}
			idSucursalesMap[i] = true
		}
	}

	for suc, _ := range idSucursalesMap {
		usu.IdSucursales = append(usu.IdSucursales, suc)
	}

	return nil

}

func (u *usuHistorialPuntosSQLAdapter) TraerInfoFacturas(usu *usuhistorialpuntos.DtoSalida) error {
	var products, categories string
	err := u.facturasdb.QueryRow(`
     with recursive categorias_info(id_categoria,
		nombre,
		id_categoria_padre,
		profundidad,
		id_categorias_total) as (
		select
			cc.id_categoria,
			cc.nombre,
			cc.id_categoria_padre,
			1::int as profundidad,
			cc.id_categoria::text as id_categorias_total
		from
			com_categorias cc
		where
			cc.id_categoria_padre is null
		union all
		select
			cc2.id_categoria,
			cc2.nombre,
			cc2.id_categoria_padre,
			ci.profundidad + 1 as profundidad,
			(ci.id_categorias_total || ',' || cc2.id_categoria::text)
		from
			categorias_info as ci
		inner join com_categorias as cc2 on
			cc2.id_categoria_padre = ci.id_categoria )
		select
			coalesce(string_agg(distinct cp.codigo_item,','),''),
			coalesce(string_agg(ci.id_categorias_total,','), '') as "categorias"
		from
			com_productos cp
		left join categorias_info ci on
			cp.id_categoria = ci.id_categoria
		inner join com_facturas_items cfi on
			cfi.id_producto = cp.id_producto
		inner join com_facturas cf on
			cf.id_factura = cfi.id_factura
		where
			cf.id_historial = $1;`, usu.IdHistorial).
		Scan(&products, &categories)

	if err != nil {
		return fmt.Errorf("error getting extra info from facturas, %v", err)
	}

	catMap := make(map[int]bool)

	for _, cat := range strings.Split(categories, ",") {
		if cat != "" {
			i, err := strconv.Atoi(cat)
			if err != nil {
				log.Printf("error parsing category: %s", err)
				continue
			}
			catMap[i] = true
		}
	}

	for cat, _ := range catMap {
		usu.Categorias = append(usu.Categorias, cat)
	}

	prodMap := make(map[string]bool)

	for _, prod := range strings.Split(products, ",") {
		if prod != "" {
			prodMap[prod] = true
		}
	}

	for prod, _ := range prodMap {
		usu.Productos = append(usu.Productos, prod)
	}

	return nil
}
