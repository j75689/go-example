package main

import (
	"encoding/json"
	"log"
	"strings"
	"viper-example/config"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var configuration config.Config
	// default value, if not set, env can't binding (viper v1.7.0)
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 3306)
	v.SetDefault("database.username", "root")
	v.SetDefault("database.password", "123456")
	v.SetDefault("messagequeue.google_pubsub.projectID", "")
	v.SetDefault("messagequeue.google_pubsub.topicID", "")
	v.SetDefault("messagequeue.google_pubsub.subID", "")

	v.ReadInConfig()
	v.Unmarshal(&configuration)

	configStr, err := json.Marshal(configuration)
	if err != nil {
		log.Fatalf("unable to json marshal config, %v", err)
	}
	log.Printf("%s", configStr)
}
