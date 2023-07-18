package config

type HttpServer struct {
	Host string
	Port uint16
}

type Config struct {
	HttpServer HttpServer
}

func New() *Config {
	return &Config{
		HttpServer{
			Host: "localhost",
			Port: 8080,
		},
	}
}
