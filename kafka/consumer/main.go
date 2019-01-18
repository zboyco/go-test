package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func main() {
	go NewConsumer("one")
	go NewConsumer("two")
	go NewConsumer("three")
	NewConsumer("four")
}

func NewConsumer(consumerId string) {
	config := cluster.NewConfig()
	// 返回错误
	config.Consumer.Return.Errors = true
	// offsets提交间隔
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	// 默认从最新的offsets开始获取
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Group.Return.Notifications = true

	brokers := []string{"localhost:9092"}
	// kafka地址
	topics := []string{"test"}
	// 话题

	// 第二个参数是groupId
	consumer, err := cluster.NewConsumer(brokers, "consumer-group1", topics, config)
	if err != nil {
		log.Panicln(err)
	}
	defer consumer.Close()

	// 系统消息
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 接收错误
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// 打印一些rebalance的信息
	go func() {
		for ntf := range consumer.Notifications() {
			log.Println("Rebalanced: ", ntf)
		}
	}()

	// 消费消息
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				log.Printf("%s : %s/%d/%d\t%s\t%s\n", consumerId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // 提交offset
			}
		case <-signals:
			return
		}
	}
}
