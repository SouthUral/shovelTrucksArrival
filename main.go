package main

import (
	"os"
	"time"

	env "github.com/SouthUral/shovelTrucksArrival/envmanager"

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
	start := time.Now()

	chEnv := env.InitLoadEnvs()
	answer := <-chEnv

	timeStart := time.Since(start).Microseconds()

	log.Info(timeStart)
	log.Info(answer.PostgresURL)
	log.Info(answer.RabbitURL)
	if answer.Error != nil {
		log.Error(answer.Error)
	}
}
