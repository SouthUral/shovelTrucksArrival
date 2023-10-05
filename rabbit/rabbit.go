package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"

	log "github.com/sirupsen/logrus"
)

// метод для создания коннекта
func createConnRabbit(URL string) (*amqp.Connection, error) {
	сonn, err := amqp.Dial(URL)
	if err != nil {
		log.Error(err.Error())
	}
	return сonn, err
}

// метод для создания канала
func createChannRabbit(conn *amqp.Connection) (*amqp.Channel, error) {
	rabbitChan, err := conn.Channel()
	if err != nil {
		log.Error(err.Error())
	}
	return rabbitChan, err
}

// метод для создания очереди
func declaringQueue(rabbitChan *amqp.Channel, nameQueue string) (amqp.Queue, error) {
	queue, err := rabbitChan.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("Ошибка создания очереди: %s", err.Error())
	}
	return queue, err
}

// Метод для отправки сообщений
func sendingMess(queue amqp.Queue, chanRabbit *amqp.Channel, mess string) {

}
