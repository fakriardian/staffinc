package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fakriardian/staffinc/internal/model/constant"
)

func (h *handler) ConsumerUpdateHarga() error {
	topic := os.Getenv("KAfKA_INPUT_HARGA_TOPIC")
	reader := GetKafkaReader(topic)

	defer reader.Close()

	fmt.Println("start consuming harga ... !!")
	temp := constant.InputHargaRequest{}

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(msg.Value, &temp)
		if err != nil {
			fmt.Println(err)
		}

		hargaData, err := h.emasUseCase.ConsumerUpdateHarga(temp)

		message, _ := json.Marshal(hargaData)

		if err != nil {
			return err
		}

		fmt.Printf("get new message %s\n", string(message))
		time.Sleep(1000 * time.Millisecond)
	}
}
