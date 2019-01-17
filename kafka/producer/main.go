package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	log.Println("测试kafka生产者")

	config := sarama.NewConfig()

	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll

	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	//设置随机分配订阅者
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	//设置kafka版本
	config.Version = sarama.V2_1_0_0

	//新建异步生产者
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.AsyncClose()

	//循环判断哪个通道发送来数据
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				// log.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				log.Println("err: ", fail)
			}
		}
	}(producer)

	//生产消息
	var value string
	for {
		time.Sleep(1000 * time.Millisecond)

		value = "message " + time.Now().Format("2006/01/02 15:04:05")

		//新建消息体
		msg := &sarama.ProducerMessage{
			Topic: "test",
		}

		//赋值
		msg.Value = sarama.ByteEncoder(value)

		//使用通道发送
		producer.Input() <- msg

	}
}
