package encoding

import (
	"encoding/json"
	_ "fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	var yamlData YAMLData
	fileInput, err := os.ReadFile(j.FileInput)
	if err != nil {
		panic("не удалось прочитать файл")
	}
	err = json.Unmarshal(fileInput, &yamlData.DockerCompose)
	if err != nil {
		panic("не удалось десериализировать")
	}

	fileOutput, err := os.Create(j.FileOutput)
	if err != nil {
		panic("не удалось создать файл")
	}
	defer fileOutput.Close()
	bytesOfYaml, err := yaml.Marshal(yamlData.DockerCompose)
	if err != nil {
		panic("не удалось преобразовать файл в json")
	}
	fileOutput.Write(bytesOfYaml)
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var jsonData JSONData
	fileInput, err := os.ReadFile(y.FileInput)
	if err != nil {
		panic("не удалось прочитать файл")
	}
	err = yaml.Unmarshal(fileInput, &jsonData.DockerCompose)
	if err != nil {
		panic("не удалось десериализировать")
	}

	fileOutput, err := os.Create(y.FileOutput)
	if err != nil {
		panic("не удалось создать файл")
	}
	defer fileOutput.Close()
	bytesOfJson, err := json.Marshal(jsonData.DockerCompose)
	if err != nil {
		panic("не удалось преобразовать файл в json")
	}
	fileOutput.Write(bytesOfJson)
	return nil
}
