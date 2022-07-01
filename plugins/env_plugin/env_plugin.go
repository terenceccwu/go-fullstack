package env_plugin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var ENV map[string]string

func Register() {
	err := godotenv.Load(".env")

	ENV = LoadEnv()

	if err != nil {
		log.Infof("Error loading .env file")
	}
}

func LoadEnv() map[string]string {
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}
	environment := getenvironment(os.Environ(), func(item string) (key, val string) {
		splits := strings.Split(item, "=")
		key = splits[0]
		val = splits[1]
		return
	})

	return environment
}

// return empty string if env not exists
func GetEnv(key string, defaultValue ...string) string {
	if value, exists := ENV[key]; exists {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

func GetEnvInt(key string, defaultValue ...string) int {
	str := GetEnv(key, defaultValue...)
	if str == "" {
		return 0
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("unable to parse ENV: %s. %v", key, err))
	}

	return value
}

func GetEnvTimeSecond(key string, defaultValue ...string) time.Duration {
	value := GetEnvInt(key, defaultValue...)
	return time.Duration(value) * time.Second
}
