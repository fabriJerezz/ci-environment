package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"log"
)

func GetCapital(province string) (string, error) {
	filePath := GetFilePath("data/capital_cities.json") 
	jsonContent, err := os.ReadFile(filePath)
	
	if err != nil {
		return "", fmt.Errorf("Error al leer el archivo de ciudades capitales %w", err)
	}

	var capitalCities map[string]string

	err = json.Unmarshal(jsonContent, &capitalCities)
	if err != nil {
		return "", fmt.Errorf("Error al deserializar el archivo de ciudades capitales %w", err)
	}

	simplifiedProvince := strings.ToUpper(province)
	simplifiedProvince = RemoveAccentMarks(simplifiedProvince)

	capitalCity := capitalCities[simplifiedProvince]
	if capitalCity == "" {
		answer := ""
		return answer, nil
	}

	return capitalCity, nil
}

func RemoveAccentMarks(str string) string {
	str = strings.ReplaceAll(str, "Á", "A")

	str = strings.ReplaceAll(str, "É", "E")

	str = strings.ReplaceAll(str, "Í", "I")

	str = strings.ReplaceAll(str, "Ó", "O")

	str = strings.ReplaceAll(str, "Ú", "U")
	return str
}


func GetFilePath(pathInsideTheProject string) string {
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Couldn't obtain the current file path")
	}
	projectFilePath := filepath.Dir(currentFilePath)
	projectFilePath = filepath.Dir(projectFilePath)
	targetFilePath := filepath.Join(projectFilePath, pathInsideTheProject)
	return targetFilePath
}