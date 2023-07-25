package config

type HttpServer struct {
	Host string
}

type Config struct {
	HttpServer          HttpServer
	DevelopmentNickname string
}

func New() *Config {
	return &Config{
		HttpServer{
			Host: "https://ultimathul3.ru/sea-server",
		},
		"test",
	}
}
