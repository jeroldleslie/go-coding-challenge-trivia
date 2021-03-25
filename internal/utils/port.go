package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ParseApiPortFromEnv(defaultPort string) string {
	const portNumberEnv = "API_PORT"
	port := os.Getenv(portNumberEnv)
	if len(port) == 0 {
		logrus.Tracef("missing env %+v, so using default port: %+v", portNumberEnv, defaultPort)
		port = defaultPort
	}
	return port
}
