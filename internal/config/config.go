package config

type Config struct{
	Addr string
	Assets string
}

func New() Config {
	return Config{Addr: ":8080", Assets: "web/public/assets"}
}