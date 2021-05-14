package usuhistorialpuntosadapters

import (
	"encoding/json"
	"fmt"

	producer "github.com/mitooos/kinesis-producer"

	kinesisproducer "leal.co/listas-aggregator/src/common/kinesis-producer"
	usu_historial_puntos "leal.co/listas-aggregator/src/usu_historial_puntos/domain"
)

type UsuHistorialPuntosKinesisAdapter interface {
	PutRecord(usu_historial_puntos.DtoSalida) error
}

func NewUsuHistorialPuntosKinesisAdapter() UsuHistorialPuntosKinesisAdapter {
	return &usuHistorialPuntosKinesisAdapter{
		kinesisProducer: kinesisproducer.GetProducer(),
	}
}

type usuHistorialPuntosKinesisAdapter struct {
	kinesisProducer *producer.Producer
}

func (u *usuHistorialPuntosKinesisAdapter) PutRecord(dto usu_historial_puntos.DtoSalida) error {
	//Dto -> data, Event -> name event
	objSalida := struct {
		Event string                         `json:"event"`
		Data  usu_historial_puntos.DtoSalida `json:"data"`
	}{
		Event: "puntos_cargados",
		Data:  dto,
	}
	usuarioActualizado, err := json.Marshal(objSalida)
	fmt.Println(string(usuarioActualizado))
	if err != nil {
		return fmt.Errorf("error putting dto to kinesis, %v", err)
	}
	err = u.kinesisProducer.Put(usuarioActualizado, dto.Uid)
	if err != nil {
		return fmt.Errorf("error putting dto to kinesis, %v", err)
	}

	return nil
}
