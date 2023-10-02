package envmanager

// Канал для отправки сообщения AnswerEnv от LoadEnvs
type AnswerEnvCh chan AnswerEnv

// Сообщение, которое вернется из горутины LoadEnvs
type AnswerEnv interface{}

// Структура для хранения переменных окружения, необходимых для подключения к БД Postgres
type EnvPostgres struct {
	Host     string
	Port     string
	Login    string
	Password string
	DBName   string
}

// Структура для хранения переменных окружения, необходимых для подключения к RabbitMQ
type EnvRabbitMQ struct {
	Host     string
	Port     string
	Login    string
	Password string
	VHost    string
}
