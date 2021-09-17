package config

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type Config struct {
	SERVICE_NAME string
}

func New() *Config {
	config := &Config{}
	v := reflect.ValueOf(*c)
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
	return config
}

func getBoolFromEnv(key string) bool {
	val := strings.ToLower(os.GetEnv(key))
	if val == "true" {
		return true
	} else if val == "false" {
		return false
	}
	log.Fatalf("error building config -- <%s> [key %s] is not an acceptable bool value", val, key)
}

func getIntFromEnv(key string) int {
	val := os.GetEnv(key)
	intVal, err := strconv.AtoI(val)
	if err != nil {
		log.Fatalf("error building config -- cannot convert <%s> [key %s] to type int\nOriginal error: %s", val, key, err.Error())
	}
	return intVal
}

func getStringFromEnv(key string) string {
	val := os.GetEnv(key)
	if val == "" {
		log.Fatalf("error building config -- empty value for key %s", key)
	}
	return val
}
