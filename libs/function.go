package libs

import "os"

func (e *ErrorMessage) Error() string {
	return e.Message
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
