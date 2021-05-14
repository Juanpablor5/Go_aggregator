package usu_historial_puntos_service

import (
	usu_historial_puntos_adapters "leal.co/listas-aggregator/src/usu_historial_puntos/adapters"
	usu_historial_puntos "leal.co/listas-aggregator/src/usu_historial_puntos/domain"
	"log"
)

type UsuHistorialPuntosService interface {
	CreateHistorialPuntos(dto usu_historial_puntos.DtoIngreso)
	UpdateHistorialPuntos(dto usu_historial_puntos.DtoIngreso)
}

func NewUsuHistorialPuntosService() UsuHistorialPuntosService {
	return &usuHistorialPuntosService{
		kinesisAdapter: usu_historial_puntos_adapters.NewUsuHistorialPuntosKinesisAdapter(),
		sqlAdapter:     usu_historial_puntos_adapters.NewUsuHistorialPuntosSqlAdapter(),
	}
}

type usuHistorialPuntosService struct {
	kinesisAdapter usu_historial_puntos_adapters.UsuHistorialPuntosKinesisAdapter
	sqlAdapter     usu_historial_puntos_adapters.UsuHistorialPuntosSQLAdapter
}

func (u *usuHistorialPuntosService) CreateHistorialPuntos(dto usu_historial_puntos.DtoIngreso) {
	// buscar faltante
	dtoSalida := dto.ToDtoSalida()

	if err := u.sqlAdapter.TraerInfoExtra(&dtoSalida); err != nil {
		log.Println(err)
	}
	// validar factura en el caso traer info
	traerProductos := dtoSalida.VerificarFactura()

	//validar y en el caso traer info
	if traerProductos {
		if err := u.sqlAdapter.TraerInfoFacturas(&dtoSalida); err != nil {
			log.Println(err)
		}
	}
	//enviar a kinesis otra vez
	u.kinesisAdapter.PutRecord(dtoSalida)
}

func (u *usuHistorialPuntosService) UpdateHistorialPuntos(dto usu_historial_puntos.DtoIngreso) {
	// convertir a dto salido

	dtoSalida := dto.ToDtoSalida()
	if err := u.sqlAdapter.TraerInfoExtra(&dtoSalida); err != nil {
		log.Println(err)
	}

	// validar factura en el caso traer info
	traerProductos := dtoSalida.VerificarFactura()

	//validar y en el caso traer info
	if traerProductos {
		log.Println("Busca productos update")
		if err := u.sqlAdapter.TraerInfoFacturas(&dtoSalida); err != nil {
			log.Println(err)
		}
	}

	log.Printf("%+v", dtoSalida)

	u.kinesisAdapter.PutRecord(dtoSalida)
}
