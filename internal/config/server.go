package config

func GetPort() string {
	return getenvWithDefaultValue("SERVER_PORT", "8080")
}
