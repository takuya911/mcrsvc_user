package conf

import (
	"os"
)

type config struct {
	DB struct {
		Host      string
		Port      string
		Name      string
		User      string
		Pass      string
		Collation string
	}
}

var C config

const (
	defaultDBHost      = "db"
	defaultDBPort      = "3306"
	defaultDBName      = "mysql-db"
	defaultDBUser      = "root"
	defaultDBPass      = "password"
	defaultDBCollation = "utf8mb4_unicode_ci"
)

func init() {
	if C.DB.Host = os.Getenv("DB_HOST"); C.DB.Host == "" {
		C.DB.Host = defaultDBHost
	}
	if C.DB.Port = os.Getenv("DB_PORT"); C.DB.Port == "" {
		C.DB.Port = defaultDBPort
	}
	if C.DB.Name = os.Getenv("DB_NAME"); C.DB.Name == "" {
		C.DB.Name = defaultDBName
	}
	if C.DB.User = os.Getenv("DB_USER"); C.DB.User == "" {
		C.DB.User = defaultDBUser
	}
	if C.DB.Pass = os.Getenv("DB_PASS"); C.DB.Pass == "" {
		C.DB.Pass = defaultDBPass
	}
	if C.DB.Collation = os.Getenv("DB_COLLATION"); C.DB.Collation == "" {
		C.DB.Collation = defaultDBCollation
	}

}
