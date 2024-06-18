package models

type Config struct {
	Database struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
	Web struct {
		Port int
		Cert string
		Key  string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Db       int
	}
}
