package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

func GetCapital(province string) (string, error) {
	filePath := GetFilePath("data", "capital_cities.json") 
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


func GetFilePath(folder string, filename string) string {
	wd, err := os.Getwd()
	wd = filepath.Dir(wd)
	
	if err != nil {
		panic(err)
	}
	
	return filepath.Join(wd, folder, filename)
}