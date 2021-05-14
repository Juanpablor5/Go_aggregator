package usuusuariosports

import (
	"encoding/json"
	usuusuarios "leal.co/listas-aggregator/src/usu_usuarios/domain"
	usuusuariosservice "leal.co/listas-aggregator/src/usu_usuarios/service"
	"log"
)

type UsuUsuariosPorts struct {
	Canal   chan []byte
	service usuusuariosservice.UsuUsuariosService
}

func NewUsuUsuariosPorts() *UsuUsuariosPorts {
	ports := UsuUsuariosPorts{
		Canal:   make(chan []byte),
		service: usuusuariosservice.NewUsuUsuariosService(),
	}

	go ports.ReadEvents()

	return &ports
}

func (u *UsuUsuariosPorts) ReadEvents() {
	for payload := range u.Canal {
		// sacar el dto del payload
		var payloadStruct struct {
			Metadata struct {
				Operation string `json:"operation"`
			} `json:"metadata"`
			Data usuusuarios.DtoIngreso `json:"data"`
		}

		if err := json.Unmarshal(payload, &payloadStruct); err != nil {
			log.Printf("error parsing json payload, %v", err)
			continue
		}

		switch payloadStruct.Metadata.Operation {
		case "update":
			go u.service.UpdateUsuario(payloadStruct.Data)
		case "create":
			go u.service.CreateUsuario(payloadStruct.Data)
		}

	}
}
