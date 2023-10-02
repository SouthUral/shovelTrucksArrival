package envmanager

// Канал для отправки сообщения AnswerEnv от LoadEnvs
type AnswerEnvCh chan AnswerEnv

// Сообщение, которое вернется из горутины LoadEnvs
type AnswerEnv interface{}
