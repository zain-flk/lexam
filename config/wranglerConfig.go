package config

import (
	"encoding/json"
	"os"
)

// GenerateWranglerConfig creates and saves the wrangler.json file
func GenerateWranglerConfig(filePath, name string) error {
	config := map[string]interface{}{
		"$schema":            "node_modules/wrangler/config-schema.json",
		"name":               name,
		"main":               "src/index.ts",
		"compatibility_date": "2025-03-13",
		"observability": map[string]bool{
			"enabled": true,
		},
	}

	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}
