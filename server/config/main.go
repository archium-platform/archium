package config

type Config struct {
	WSPort string
	WSPath string
}

var DefaultConfig = Config{
	WSPort: ":443",
	WSPath: "/ws",
}
