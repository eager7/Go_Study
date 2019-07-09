package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
)

func Initialize(addresses []string, groupId string, topics []string) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// init consumer
	consumer, err := cluster.NewConsumer(addresses, groupId, topics, config)
	if err != nil {
		fmt.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupId, err)
		return
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			fmt.Printf("%s:Error: %s\n", groupId, err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			fmt.Printf("%s:Rebalanced: %+v \n", groupId, ntf)
		}
	}()
Loop:
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Println("msg:", msg.Topic, string(msg.Value))
			}
		case <-signals:
			break Loop
		}
	}
}
