package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Config config
)

type config struct {
	Mode string

	FrontendHost string
	FrontendPort string

	CookieKey  string
	CookieName string

	MysqlHost string
	MysqlPort string
	MysqlUser string
	MysqlPass string
	MysqlDb   string
}

func (c *config) MysqlDSN() string {
	return c.MysqlUser + ":" + c.MysqlPass + "@tcp(" + c.MysqlHost + ":" + c.MysqlPort + ")/" + c.MysqlDb + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (c *config) FrontendURL() string {
	if c.FrontendPort == "80" {
		return "http://" + c.FrontendHost
	}
	if c.FrontendPort == "443" {
		return "https://" + c.FrontendHost
	}
	return "http://" + c.FrontendHost + ":" + c.FrontendPort
}

func Getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Getenvint(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		d, err := strconv.Atoi(v)
		if err != nil {
			return def
		}
		return d
	}
	return def
}

func Getenvfloat(key string, def float64) float64 {
	if v := os.Getenv(key); v != "" {
		d, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return def
		}
		return d
	}
	return def
}

func init() {
	MODE := Getenv("MODE", "dev")
	Config.Mode = MODE
	if MODE == "dev" {
		godotenv.Load("development.env")
		Config.FrontendHost = Getenv("FRONTEND_HOST", "localhost")
		Config.FrontendPort = Getenv("FRONTEND_PORT", "5173")
		Config.CookieName = Getenv("COOKIE_NAME", "default")
		Config.CookieKey = Getenv("COOKIE_KEY", "default")
		Config.MysqlHost = Getenv("MYSQL_HOST", "localhost")
		Config.MysqlPort = Getenv("MYSQL_PORT", "3306")
		Config.MysqlUser = Getenv("MYSQL_USER", "root")
		Config.MysqlPass = Getenv("MYSQL_PASSWORD", "password")
		Config.MysqlDb = Getenv("MYSQL_DATABASE", "default")
		log.Println("[ENV] Loaded development.env")
	}
	if MODE == "prod" {
		log.Println("[ENV] Loaded production environment variables")
	}
}
