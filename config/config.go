package config

const (
	// RabbitMQ connection (docker-compose uses 'rabbitmq' hostname)
	RabbitURL    = "amqp://guest:guest@localhost:5672/"
	ExchangeMain = "app_exchange"        // direct exchange for normal deliver
	ExchangeDead = "dead_exchange"       // dead-letter exchange
	ExchangeDelayed = "delayed_exchange" // x-delayed-message exchange for retries
	MainQueue    = "main_queue"
	DeadQueue    = "dead_queue"
	RoutingKey   = "task"                // routing key for main_queue
	MaxRetries   = 3                     // max retry attempts
)
