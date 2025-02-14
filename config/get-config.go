package config

import (
	"encoding/json"
	"log"
	"os"
)

func GetConfig(key string) string {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	value, exists := config[key]
	if !exists {
		log.Fatalf("key %s not found in config", key)
	}

	// **تبدیل مقدار به string**
	strValue, ok := value.(string)
	if !ok {
		log.Fatalf("key %s is not a string", key)
	}

	return strValue
}
