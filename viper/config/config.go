package config

type Config struct {
	Database     Database     `mapstructure:"database"`
	MessageQueue MessageQueue `mapstructure:"messagequeue"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type MessageQueue struct {
	GooglePubSub GooglePubSub `mapstructure:"google_pubsub"`
}

type GooglePubSub struct {
	ProjectID string `mapstructure:"projectID"`
	TopicID   string `mapstructure:"topicID"`
	SubID     string `mapstructure:"subID"`
}
