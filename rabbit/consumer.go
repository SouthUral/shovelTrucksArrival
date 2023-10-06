package rabbit

import (
	log "github.com/sirupsen/logrus"
)

// Базовый Consumer
func Consumer(URL string, nameQueue string, consumerName string) {
	conn, err := createConnRabbit(URL)
	if err != nil {
		return
	}

	chanRb, errChan := createChannRabbit(conn)
	if errChan != nil {
		return
	}

	// Qos определяет, сколько сообщений или сколько байт сервер попытается сохранить в сети для потребителей,
	// прежде чем получит подтверждение доставки.
	errQos := chanRb.Qos(1, 0, false)
	if errQos != nil {
		log.Errorf("Qos Rabbit failed: %v\n", err)
	}

	queue, errQueue := declaringQueue(chanRb, nameQueue)
	if errQueue != nil {
		return
	}

	msgs, errConsumer := createConsumer(chanRb, queue.Name, consumerName)
	if errConsumer != nil {
		return
	}

	for msg := range msgs {
		log.Infof("Received a message: %s", msg.Body)
		msg.Ack(true)
	}

	defer exitAndClose(chanRb, conn, "Consumer")
}
