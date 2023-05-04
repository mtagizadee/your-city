package config

type server struct {
  Host string
  Port string
}

var serverConfig server

func loadServer(port string, host string) {
  serverConfig.Host = host
  serverConfig.Port = port
}

func GetServerConfig() server {
  return serverConfig
}
