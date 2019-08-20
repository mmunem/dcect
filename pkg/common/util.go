package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RequireEnv validates and returns an environment variable's value if it exists
func RequireEnv(env string) string {
	val, ok := os.LookupEnv(env)
	if !ok {
		panic(fmt.Errorf("missing required env var %v", env))
	}
	return val
}

func RequireEnvInt(env string) int {
	val := RequireEnv(env)
	intVal, err := strconv.Atoi(val)

	if err != nil {
		panic(fmt.Errorf("unable to parse env var %s as int: %s", env, val))
	}

	return intVal
}

func RequireEnvStringSlice(env string, sep string) []string {
	val := RequireEnv(env)
	list := strings.Split(val, sep)

	// Remove empty values
	cleanList := []string{}
	for _, value := range list {
		if value != "" {
			cleanList = append(cleanList, value)
		}
	}

	return cleanList
}

func RequireEnvFloatSlice(env string, sep string) []float64 {
	strList := RequireEnvStringSlice(env, sep)

	floatList := []float64{}
	for _, val := range strList {
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(fmt.Errorf("env var %s contains invalid float64 value: %s", env, val))
		}
		floatList = append(floatList, floatVal)
	}

	return floatList
}
