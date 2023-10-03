package envmanager

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Функция формирует канал для ответа и запускает горутину LoadEnvs.
// Возвращает канал для ответа от LoadEnvs
func InitLoadEnvs() AnswerEnvCh {
	answerCh := make(AnswerEnvCh)
	go loadEnvs(answerCh)
	return answerCh
}

// Функция для загрузки переменной окружения по ключу.
// Если переменная не загружена, то функция вернет пустую строку и ошибку.
func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if exists {
		return value, nil
	}
	err := fmt.Errorf("переменная %s не загружена", key)
	log.Warning(err.Error())
	return "", err
}

// Функция загружает переменные окружения для модулей Postgres и Rabbit.
// В случае если какие-то переменные не загружены, будет сформирована ошибка
func loadEnvs(answerCh AnswerEnvCh) {
	pgEnv, errPg := loadPostgresEnv()
	rbEnv, errRb := loadRabbitEnv()

	responce := AnswerEnv{
		PostgresURL: pgEnv.getPostgresURL(),
		RabbitURL:   rbEnv.getRabbitURL(),
	}

	if errPg != nil || errRb != nil {
		responce.Error = fmt.Errorf("не загружены env %s, %s", errPg, errRb)
	}

	answerCh <- responce
}

func loadingEnvVar(fieldsAndAliasEnv envAlias) (envStorageStruct, error) {
	envStorage := envStorageStruct{}
	errorFields := make([]string, 0)

	for fieldName, envAlias := range fieldsAndAliasEnv {
		loadedEnv, err := getEnv(envAlias)
		if err != nil {
			errorFields = append(errorFields, envAlias)
		}
		envStorage.readAndWriteField(fieldName, loadedEnv)
	}

	if len(errorFields) > 0 {
		errRes := fmt.Errorf("%s", strings.Join(errorFields, ", "))
		return envStorage, errRes
	}
	return envStorage, nil
}

func loadPostgresEnv() (envStorageStruct, error) {
	envsAlias := envAlias{
		"Host":     "ASD_POSTGRES_HOST",
		"Port":     "ASD_POSTGRES_PORT",
		"Login":    "ASD_POSTGRES_LOGIN",
		"Password": "ASD_POSTGRES_PASSWORD",
		"DBName":   "ASD_POSTGRES_DBNAME",
	}
	loadedEnvs, err := loadingEnvVar(envsAlias)
	return loadedEnvs, err
}

func loadRabbitEnv() (envStorageStruct, error) {
	envsAlias := envAlias{
		"Host":     "ASD_RMQ_HOST",
		"Port":     "ASD_RMQ_PORT",
		"Login":    "ASD_RMQ_LOGIN",
		"Password": "ASD_RMQ_PASSWORD",
		"VHost":    "ASD_RMQ_HOST",
	}
	loadedEnvs, err := loadingEnvVar(envsAlias)
	return loadedEnvs, err
}
