package config

type jwt struct {
	Secret string
}

var jwtConfig jwt

func loadJwt(secret string) {
	jwtConfig.Secret = secret
}

func GetJwtConfig() jwt {
	return jwtConfig
}