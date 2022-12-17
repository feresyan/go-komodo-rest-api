package env

type Environment struct {
	GokomodoDB Database
	SecretKey  Secret
}
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}

type Secret struct {
	Value string
}
