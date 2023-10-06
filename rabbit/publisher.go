package rabbit

import (
	"context"
	"fmt"
	"time"
)

func Publisher(URL string, nameExchange, nameQueue string) {

	conn, err := createConnRabbit(URL)
	if err != nil {
		return
	}

	chanRb, errChan := createChannRabbit(conn)
	if errChan != nil {
		return
	}

	queue, errQueue := declaringQueue(chanRb, nameQueue)
	if errQueue != nil {
		return
	}

	errExchange := declaringExchange(chanRb, nameExchange)
	if errExchange != nil {
		return
	}

	context, _ := context.WithTimeout(context.Background(), 30*time.Second)

	for counter := 0; counter < 10; {
		mess := fmt.Sprintf("Message № %d", counter)
		err := sendingMess(queue, chanRb, context, nameExchange, mess)
		if err != nil {
			return
		}
		counter++
	}

	defer exitAndClose(chanRb, conn, "Publisher")
}
