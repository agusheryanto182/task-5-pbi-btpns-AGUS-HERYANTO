package models

type Global struct {
	Database Database
	Server   Server
}

type Server struct {
	Host string `env:"SERVER_HOST,required"`
	Port string `env:"SERVER_PORT,required"`
}

type Database struct {
	User     string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Host     string `env:"MYSQL_HOST,required"`
	Port     string `env:"MYSQL_PORT,required"`
	Name     string `env:"MYSQL_DBNAME,required"`
}
