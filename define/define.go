package define

var (
	ConfigDatabase = map[string]string{
		"host": "127.0.0.1",
		"name": "skeleton",
		"port": "5434",
		"type": "postgres",
		"user": "postgres",
		"pass": "postgres",
	}

	JWTKey = "anfin-jwt"
)
