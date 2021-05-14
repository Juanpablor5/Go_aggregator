package usu_usuarios_comercios_ports

import (
	"encoding/json"
	"log"

	usu_usuarios_comercios "leal.co/listas-aggregator/src/usu_usuarios_comercios/domain"
	usu_usuarios_comercios_service "leal.co/listas-aggregator/src/usu_usuarios_comercios/service"
)

type UsuUsuariosComerciosPorts struct {
	Canal   chan []byte
	service usu_usuarios_comercios_service.UsuUsuariosComerciosService
}

func NewUsuUsuariosComerciosPorts() *UsuUsuariosComerciosPorts {
	ports := UsuUsuariosComerciosPorts{
		Canal:   make(chan []byte),
		service: usu_usuarios_comercios_service.NewUsuUsuariosComerciosService(),
	}

	go ports.ReadEvents()

	return &ports
}

func (u *UsuUsuariosComerciosPorts) ReadEvents() {
	for payload := range u.Canal {
		// sacar el dto del payload
		var payloadStruct struct {
			Metadata struct {
				Operation string `json:"operation"`
			} `json:"metadata"`
			Data usu_usuarios_comercios.DtoIngreso `json:"data"`
		}

		if err := json.Unmarshal(payload, &payloadStruct); err != nil {
			log.Printf("error parsing json payload, %v", err)
			continue
		}

		switch payloadStruct.Metadata.Operation {
		case "update":
			go u.service.UpdateUsuUsuariosComercios(payloadStruct.Data)
		case "insert":
			go u.service.CreateUsuUsuariosComercios(payloadStruct.Data)
		}
	}
}
