package envmanager

import "fmt"

// Канал для отправки сообщения AnswerEnv от LoadEnvs
type AnswerEnvCh chan AnswerEnv

// Сообщение, которое вернется из горутины LoadEnvs
type AnswerEnv struct {
	PostgresURL string
	RabbitURL   string
	Error       error
}

// Структура для хранения переменных окружения, необходимых для подключения к БД Postgres
type envStorageStruct struct {
	Host      string
	Port      string
	Login     string
	Password  string
	LastField string
}

// Метод формирует корневой элемент URL для подключения
func (envStorage *envStorageStruct) formationConnString() string {
	return fmt.Sprintf("%s:%s@%s:%s/%s", envStorage.Login, envStorage.Password, envStorage.Host, envStorage.Port, envStorage.LastField)
}

// Метод для формирования URL для подключения к Postgres DB
func (envStorage *envStorageStruct) getPostgresURL() string {
	return fmt.Sprintf("postgresql://%s", envStorage.formationConnString())
}

// Метод для формирования URL для подключения к RabbitMQ
func (envStorage *envStorageStruct) getRabbitURL() string {
	return fmt.Sprintf("amqp://%s", envStorage.formationConnString())
}

// Метод для записи переменных окружения в структуру по наименованию поля
func (envStorage *envStorageStruct) readAndWriteField(NameFiled, ValueField string) {
	switch NameFiled {
	case "Host":
		envStorage.Host = ValueField
	case "Port":
		envStorage.Port = ValueField
	case "Login":
		envStorage.Login = ValueField
	case "Password":
		envStorage.Password = ValueField
	case "DBName", "VHost":
		envStorage.LastField = ValueField
	}
}

// Словарь, который должен содержать имена переменных окружения
type envAlias map[string]string
