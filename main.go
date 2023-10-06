package main

import (
	"os"
	"time"

	env "github.com/SouthUral/shovelTrucksArrival/envmanager"
	rb "github.com/SouthUral/shovelTrucksArrival/rabbit"

	log "github.com/sirupsen/logrus"
)

func init() {
	// логи в формате JSON, по умолчанию формат ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// логи идут на стандартный вывод, их можно перенаправить в файл
	log.SetOutput(os.Stdout)

	// установка уровня логирования
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("Запуск сервиса")
	// start := time.Now()

	chEnv := env.InitLoadEnvs()
	answer := <-chEnv

	go rb.Publisher(answer.RabbitURL, "TestExchange", "TestQueue")
	go rb.Consumer(answer.RabbitURL, "TestQueue", "Basic_consumer")

	// timeStart := time.Since(start).Microseconds()

	// log.Info(timeStart)

	if answer.Error != nil {
		log.Error(answer.Error)
	}

	time.Sleep(time.Duration(100) * time.Second)
}
