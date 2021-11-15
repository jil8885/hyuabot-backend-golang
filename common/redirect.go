package common

import "os"

func GetPrimaryServer() string{
	isPrimary, keyExists := os.LookupEnv("is_primary")
	primaryServer, addressExists := os.LookupEnv("primary_server")
	if keyExists && addressExists && isPrimary == "False" {
		return primaryServer
	} else {
		return ""
	}
}
