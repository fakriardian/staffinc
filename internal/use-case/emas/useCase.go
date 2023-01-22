package emas

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	con "github.com/fakriardian/staffinc/internal/delivery/kafka/consumer"
	pro "github.com/fakriardian/staffinc/internal/delivery/kafka/producer"
	"github.com/fakriardian/staffinc/internal/model"
	"github.com/fakriardian/staffinc/internal/model/constant"
	"github.com/fakriardian/staffinc/internal/repository/harga"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type emasUseCase struct {
	hargaRepo harga.Repository
}

func GetUseCase(hargaRepo harga.Repository) Usecase {
	return &emasUseCase{
		hargaRepo: hargaRepo,
	}
}

func (eu *emasUseCase) UpdateHarga(request constant.InputHargaRequest) (model.Harga, error) {
	connectKafka := pro.Connection()
	topic := "input-harga"
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

func (eu *emasUseCase) ConsumerUpdateHarga() error {
	topic := "input-harga"
	reader := con.GetKafkaReader(topic)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	temp := model.Harga{}

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(msg.Value, &temp)
		if err != nil {
			fmt.Println(err)
		}

		isExistingData, err := eu.hargaRepo.IsExisting()
		if err != nil {
			return err
		}

		if isExistingData != "" {
			eu.hargaRepo.DeleteExisting(isExistingData)
		}

		hargaData, err := eu.hargaRepo.AddHarga(model.Harga{
			AdminID:      temp.AdminID,
			HargaTopUp:   temp.HargaTopUp,
			HargaBuyBack: temp.HargaBuyBack,
		})

		message, _ := json.Marshal(hargaData)

		if err != nil {
			return err
		}

		fmt.Printf("get new message %s\n", string(message))
		time.Sleep(1000 * time.Millisecond)
	}
}
