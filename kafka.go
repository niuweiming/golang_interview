package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

// 生产者
func Consumer() {
	// 配置Kafka生产者
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// 连接到Kafka代理
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer producer.Close()

	// 创建消息
	msg := &sarama.ProducerMessage{
		Topic: "example-topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// 发送消息
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// 打印发送结果
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

// 消费者

func Product() {
	// 配置Kafka消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 连接到Kafka代理
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer consumer.Close()

	// 订阅主题
	partitionConsumer, err := consumer.ConsumePartition("example-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consuming partition: %v", err)
	}
	defer partitionConsumer.Close()

	// 处理消息和错误
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("Received error: %v\n", err)
		case <-signals:
			fmt.Println("Interrupt is detected")
			return
		}
	}
}
