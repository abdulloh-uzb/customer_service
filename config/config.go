package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment        string
	PostgresHost       string
	PostgresPort       int
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	LogLevel           string
	RPCPort            string
	PostServiceHost    string
	PostServicePort    int
	RankingServiceHost string
	RankingServicePort int
	KafkaHost          string
	KafkaPort          string
	KafkaTopic         string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "customerdb_abdulloh"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "abdulloh"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "abdulloh"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9000"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post-service"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 8000))

	c.RankingServiceHost = cast.ToString(getOrReturnDefault("RANKING_SERVICE_HOST", "reyting-service"))
	c.RankingServicePort = cast.ToInt(getOrReturnDefault("RANKING_SERVICE_HOST", 1111))

	c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "kafka"))
	c.KafkaPort = cast.ToString(getOrReturnDefault("KAFKA_PORT", "9092"))
	c.KafkaTopic = cast.ToString(getOrReturnDefault("KAFKA_TOPIC", "postTopic"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
