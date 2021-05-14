package usuusuariosadapters

import (
	"encoding/json"
	"fmt"

	producer "github.com/mitooos/kinesis-producer"

	kinesisproducer "leal.co/listas-aggregator/src/common/kinesis-producer"
	usuusuarios "leal.co/listas-aggregator/src/usu_usuarios/domain"
)

type UsuUsuariosKinsisAdapter interface {
	PutRecord(usuusuarios.DtoSalida) error
}

func NewUsuUsuariosKinesisAdapter() UsuUsuariosKinsisAdapter {
	return &usuUsuariosKinesisAdapter{
		kinesisProducer: kinesisproducer.GetProducer(),
	}
}

type usuUsuariosKinesisAdapter struct {
	kinesisProducer *producer.Producer
}

func (u *usuUsuariosKinesisAdapter) PutRecord(dto usuusuarios.DtoSalida) error {
	objSalida := struct {
		Event string                `json:"event"`
		Data  usuusuarios.DtoSalida `json:"data"`
	}{
		Event: "cliente_actualizado",
		Data:  dto,
	}

	usuarioActualzado, err := json.Marshal(objSalida)

	fmt.Println(string(usuarioActualzado))

	if err != nil {
		return fmt.Errorf("error putting dto to kinesis, %v", err)
	}

	err = u.kinesisProducer.Put(usuarioActualzado, dto.Uid)

	if err != nil {
		return fmt.Errorf("error putting dto to kinesis, %v", err)
	}

	return nil
}
