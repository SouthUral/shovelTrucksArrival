package main

import (
	"os"

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

	chEnv := env.InitLoadEnvs()
	answer := <-chEnv
	log.Info(answer)

}
