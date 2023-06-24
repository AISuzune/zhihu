package config

type Config struct {
	Logger *Logger `mapstructure:"logger" yaml:"logger"`
}
