package usu_historial_puntos_ports

import (
	"encoding/json"
	"log"

	usu_historial_puntos "leal.co/listas-aggregator/src/usu_historial_puntos/domain"
	usu_historial_puntos_service "leal.co/listas-aggregator/src/usu_historial_puntos/service"
)

type UsuHistorialPuntosPorts struct {
	Canal   chan []byte
	service usu_historial_puntos_service.UsuHistorialPuntosService
}

func NewUsuHistorialPuntosPorts() *UsuHistorialPuntosPorts {
	ports := UsuHistorialPuntosPorts{
		Canal:   make(chan []byte),
		service: usu_historial_puntos_service.NewUsuHistorialPuntosService(),
	}

	go ports.ReadEvents()

	return &ports
}

func (u *UsuHistorialPuntosPorts) ReadEvents() {
	for payload := range u.Canal {
		// sacar el dto del payload
		var payloadStruct struct {
			Metadata struct {
				Operation string `json:"operation"`
			} `json:"metadata"`
			Data usu_historial_puntos.DtoIngreso `json:"data"`
		}

		if err := json.Unmarshal(payload, &payloadStruct); err != nil {
			log.Printf("error parsing json payload, %v", err)
			continue
		}

		switch payloadStruct.Metadata.Operation {
		case "update":
			go u.service.UpdateHistorialPuntos(payloadStruct.Data)
		case "insert":
			go u.service.CreateHistorialPuntos(payloadStruct.Data)
		}
	}
}
