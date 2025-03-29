package converter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ConvertJSONtoEnv(filepath, envName string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var (
		data    map[string]interface{}
		envVars []string
	)
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	for key, value := range data {
		upperKey := strings.ToUpper(key)

		envVar := fmt.Sprintf("%s=%v", upperKey, value)
		envVars = append(envVars, envVar)
	}
	if !strings.HasSuffix(envName, ".env") {
		envName += ".env"
	}

	file, err = os.Create(envName)
	if err != nil {
		return fmt.Errorf("failed to create env file: %w", err)
	}
	defer file.Close()

	for _, envVar := range envVars {
		if _, err := file.WriteString(envVar + "\n"); err != nil {
			return fmt.Errorf("failed to write to env file: %w", err)
		}
	}
	fmt.Printf("Environment variables written to %v\n", envName)

	return nil
}

func ConvertEnvToJSON(envFilePath, jsonFilePath string) error {
	file, err := os.Open(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to open env file: %w", err)
	}
	defer file.Close()

	data := make(map[string]interface{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid line in env file: %s", line)
		}

		key := strings.ToLower(parts[0])
		value := parts[1]
		data[key] = value
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read env file: %w", err)
	}

	jsonFile, err := os.Create(jsonFilePath)
	if err != nil {
		return fmt.Errorf("failed to create JSON file: %w", err)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	fmt.Printf("JSON data written to %v\n", jsonFilePath)
	return nil
}
