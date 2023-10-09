package rabbit

import (
	"context"

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
		nameQueue, // name
		false,     // durable
		false,     // autoDelete - удалить если не используется
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)
	if err != nil {
		log.Errorf("Ошибка создания очереди: %s", err.Error())
	}
	return queue, err
}

// Метод для bindig очереди к exchange
func bindingQueue(rabbitChan *amqp.Channel, nameQueue, nameExchange, routingKey string) error {
	err := rabbitChan.QueueBind(
		nameQueue,    // queue name
		routingKey,   // routing key
		nameExchange, // name exchange
		false,        // noWait
		nil,
	)
	if err != nil {
		log.Errorf("Ошибка binding очереди %s, к exchange %s : %s", nameQueue, nameExchange, err.Error())
	}
	return err
}

// Метод для создания Exchange
func declaringExchange(rabbitChan *amqp.Channel, nameExchange string) error {
	err := rabbitChan.ExchangeDeclare(
		nameExchange, // имя exchange
		"direct",     // тип exchange
		true,         // durable - сохранится ли exchange после перезагрузки Rabbit
		false,        // auto delete - удалить если не используется
		false,        // internal
		false,        // noWait
		nil,          // дополнительные аргументы
	)
	if err != nil {
		log.Errorf("exchange declare failed: %v\n", err)
	}
	return err
}

// Метод для отправки сообщений
func sendingMess(queue amqp.Queue, chanRabbit *amqp.Channel, context context.Context, exchange, mess string) error {
	err := chanRabbit.PublishWithContext(
		context,    // контекст
		exchange,   // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mess),
		})
	if err != nil {
		log.Errorf(log.ErrorKey)
	} else {
		log.Infof("the message: {%s} has been sent", mess)
	}
	return err
}

func exitAndClose(chanRb *amqp.Channel, connRb *amqp.Connection, whoIsIt string) {
	errChan := chanRb.Close()
	if errChan != nil {
		log.Error("channel closing error")
	}
	errConn := connRb.Close()
	if errConn != nil {
		log.Error("connection closing error")
	}

	log.Warningf("The %s was closed", whoIsIt)
}

// Функция создает базового консюмера
func createConsumer(chanRb *amqp.Channel, queueName, consumerName string) (<-chan amqp.Delivery, error) {
	msgs, err := chanRb.Consume(
		queueName,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error("Failed to register a consumer", err)
	}
	return msgs, err
}
