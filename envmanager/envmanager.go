package envmanager

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Функция формирует канал для ответа и запускает горутину LoadEnvs.
// Возвращает канал для ответа от LoadEnvs
func InitLoadEnvs() AnswerEnvCh {
	answerCh := make(AnswerEnvCh)
	go loadEnvs(answerCh)
	return answerCh
}

// Функция для загрузки переменной окружения по ключу
func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	log.Warningf("Переменная %s не загружена", key)
	return ""
}

// Функция загружает переменные окружения для модулей Postgres и Rabbit.
// В случае если какие-то переменные не загружены, будет сформирована ошибка
func loadEnvs(answerCh AnswerEnvCh) {

}

func loadingVariablesIntoStruct[filledStruct EnvPostgres | EnvRabbitMQ](envs envAlias) filledStruct {
	for key, value := range envs {

	}
}

func loadPostgresEnv() {

}

func loadRabbitEnv() {

}
