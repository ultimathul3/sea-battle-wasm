package config

type HttpServer struct {
	Host string
	Port uint16
}

type Config struct {
	HttpServer          HttpServer
	DevelopmentNickname string
}

func New() *Config {
	return &Config{
		HttpServer{
			Host: "http://localhost",
			Port: 8082,
		},
		"test",
	}
}
