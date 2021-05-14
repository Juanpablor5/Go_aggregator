package usuusuarioscomerciosadapters

import (
	"encoding/json"
	"fmt"

	producer "github.com/mitooos/kinesis-producer"

	kinesisproducer "leal.co/listas-aggregator/src/common/kinesis-producer"
	usu_usuarios_comercios "leal.co/listas-aggregator/src/usu_usuarios_comercios/domain"
)

type UsuUsuariosComerciosKinesisAdapter interface {
	PutRecord(usu_usuarios_comercios.DtoSalida) error
}

func NewUsuUsuariosComerciosKinesisAdapter() UsuUsuariosComerciosKinesisAdapter {
	return &usuUsuariosComerciosKinesisAdapter{
		kinesisProducer: kinesisproducer.GetProducer(),
	}
}

type usuUsuariosComerciosKinesisAdapter struct {
	kinesisProducer *producer.Producer
}

func (u *usuUsuariosComerciosKinesisAdapter) PutRecord(dto usu_usuarios_comercios.DtoSalida) error {
	objSalida := struct {
		Event string                           `json:"event"`
		Data  usu_usuarios_comercios.DtoSalida `json:"data"`
	}{
		Event: "usuario_comercio_actualizado",
		Data:  dto,
	}

	bytesSalida, err := json.Marshal(objSalida)

	fmt.Println(string(bytesSalida))
	if err != nil {
		return fmt.Errorf("error putting dto to kinesis, %v", err)
	}
	err = u.kinesisProducer.Put(bytesSalida, dto.Uid)
	if err != nil {
		return fmt.Errorf("error putting dto to kinesis, %v", err)
	}

	return nil
}
