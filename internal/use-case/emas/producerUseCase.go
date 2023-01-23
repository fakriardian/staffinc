package emas

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	pro "github.com/fakriardian/staffinc/internal/delivery/kafka/producer"
	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func (eu *emasUseCase) ProducerUpdateHarga(request constant.InputHargaRequest) (model.Harga, error) {
	connectKafka := pro.Connection()
	topic := os.Getenv("KAfKA_INPUT_HARGA_TOPIC")
	topicConfig := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
	err := connectKafka.CreateTopics(topicConfig)

	if err != nil {
		return model.Harga{}, errors.New("kafka not ready")
	}

	producer := pro.NewProducer()
	marsheledProduct, _ := json.Marshal(request)
	pro.Produce([]byte(uuid.NewString()), marsheledProduct, topic, producer)

	return model.Harga{}, nil
}

func (eu *emasUseCase) validationTopUp(request constant.TopUpRequest) (bool, int64, error) {
	validationHarga, err := eu.hargaRepo.CheckHarga()
	if err != nil {
		return false, 0, err
	}

	if validationHarga.HargaTopUp != request.HargaTopUp {
		return false, 0, errors.New("harga topup berbeda")
	}

	request.HargaBuyBack = validationHarga.HargaBuyBack

	stringGram := fmt.Sprintf("%g", request.Gram)
	parts := strings.Split(stringGram, ".")

	if len(parts[1]) > 3 {
		return false, 0, errors.New("harus 3 angka di belakang koma")
	}

	// if request.Gram >= 1 {
	// 	return true, nil
	// }

	// if math.Mod(float64(request.Gram), 0.001) != 0 {
	// 	return false, errors.New("gram tidak sesuai")
	// }

	return true, validationHarga.HargaBuyBack, nil
}

func (eu *emasUseCase) ProducerTopUp(request constant.TopUpRequest) (model.Transaction, error) {
	_, hargaBuyBack, err := eu.validationTopUp(request)
	if err != nil {
		return model.Transaction{}, err
	}

	request.HargaBuyBack = hargaBuyBack

	connectKafka := pro.Connection()
	topic := os.Getenv("KAfKA_TOPUP_TOPIC")
	topicConfig := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
	err = connectKafka.CreateTopics(topicConfig)

	if err != nil {
		return model.Transaction{}, errors.New("kafka not ready")
	}

	producer := pro.NewProducer()
	marsheledProduct, _ := json.Marshal(request)
	pro.Produce([]byte(uuid.NewString()), marsheledProduct, topic, producer)

	return model.Transaction{}, nil

}

func (eu *emasUseCase) validationBuyBack(request constant.BuyBackRequest) (bool, int64, error) {
	validationHarga, err := eu.hargaRepo.CheckHarga()
	if err != nil {
		return false, 0, err
	}

	getSaldo, err := eu.rekeningRepo.CheckSaldo(request.NoRek)
	if err != nil {
		return false, 0, err
	}

	if validationHarga.HargaBuyBack != request.HargaBuyBack {
		return false, 0, errors.New("harga buyback berbeda")
	}

	if request.Gram > getSaldo.Saldo {
		return false, 0, errors.New("saldo tidak cukup")
	}

	request.HargaTopUp = validationHarga.HargaTopUp

	return true, validationHarga.HargaTopUp, nil
}

func (eu *emasUseCase) ProducerBuyBack(request constant.BuyBackRequest) (model.Transaction, error) {
	_, hargaTopup, err := eu.validationBuyBack(request)
	if err != nil {
		return model.Transaction{}, err
	}

	request.HargaTopUp = hargaTopup

	connectKafka := pro.Connection()
	topic := os.Getenv("KAfKA_BUYBACK_TOPIC")
	topicConfig := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
	err = connectKafka.CreateTopics(topicConfig)

	if err != nil {
		return model.Transaction{}, errors.New("kafka not ready")
	}

	producer := pro.NewProducer()
	marsheledProduct, _ := json.Marshal(request)
	pro.Produce([]byte(uuid.NewString()), marsheledProduct, topic, producer)

	return model.Transaction{}, nil
}
