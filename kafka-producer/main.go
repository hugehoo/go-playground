package main

import (
	"bufio"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	// Kafka 설정
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Kafka 브로커 서버 주소 설정
	brokers := []string{"localhost:29092"}

	// Kafka 프로듀서 생성
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to create Kafka producer:", err)
	}
	defer producer.Close()

	// 메시지 전송
	topic := "my-topic2"

	// Ctrl+C 시 그레이스풀하게 종료
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		fmt.Println("Kafka producer stopped")
		os.Exit(0)
	}()

	// 터미널 입력 대기
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")

		if message == "exit" {
			break
		}

		// 메시지 생성
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		// 메시지 전송
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("Failed to send message to Kafka:", err)
		} else {
			fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		}
	}
}
