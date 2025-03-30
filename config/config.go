package config

type Config struct {
	Port    string
	TempDir string
}

var GlobalConfig = Config{
	Port:    "3000",
	TempDir: "./tmp/",
}
