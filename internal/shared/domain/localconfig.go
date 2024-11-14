package domain

type Configuration struct {
	AllowedOrigins []string
	AllowedMethods []string
	Env            string
	PortHTTP       uint
	JWTSecret      string
	Timezone       string
	Database       Database
	OTEL           OTEL
}

type Database struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     uint
	Name     string
	SSLMode  string
}

type OTEL struct {
	CollectorEndpoint string
}
