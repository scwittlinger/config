package configManager

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// New initializes the struct passed into it as a config
// each field of the struct passed to New() should be named exclusively using uppercase letters and underscores
// there should be an environment variable with the same exact name as each field, which can be cast to the type of that field
func New(config *interface{}) {
	v := reflect.ValueOf(*config)
	fieldCount := v.NumField()

	for i := 0; i < fieldCount; i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			val := reflect.ValueOf(getIntFromEnv(v.Field(i).Type().Name()))
			v.Field(i).Set(val)
		case reflect.String:
			val := reflect.ValueOf(getStringFromEnv(v.Field(i).Type().Name()))
			v.Field(i).Set(val)
		case reflect.Bool:
			val := reflect.ValueOf(getBoolFromEnv(v.Field(i).Type().Name()))
			v.Field(i).Set(val)
		default:
			log.Fatalf("error building config -- %s is not of an acceptable type", v.Field(i).Type().Name())
		}
	}
}

func getBoolFromEnv(key string) bool {
	val := strings.ToLower(os.Getenv(key))
	if val == "true" {
		return true
	} else if val == "false" {
		return false
	}
	log.Fatalf("error building config -- <%s> [key %s] is not an acceptable bool value", val, key)
	return false // needed to satisfy compiler, even though this code is unreachable
}

func getIntFromEnv(key string) int {
	val := os.Getenv(key)
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("error building config -- cannot convert <%s> [key %s] to type int\nOriginal error: %s", val, key, err.Error())
	}
	return intVal
}

func getStringFromEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("error building config -- empty value for key %s", key)
	}
	return val
}
