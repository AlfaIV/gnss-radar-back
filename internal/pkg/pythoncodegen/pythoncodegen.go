package pythoncodegen

import (
	"bytes"
	"html/template"
	"os"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
)

func GenerateCode(config model.PythonGenConfig) (string, error) {
	tmpl, err := template.New("python_code").Parse(model.ImportsTemplate +
		model.ConstantsTemplate +
		model.RequestHandlerTemplate +
		model.SignalProcessorTemplate +
		model.SdrHandlerTemplate +
		model.DataProcessorTemplate +
		model.MainFunctionTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, config)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func SaveCodeToFile(config model.PythonGenConfig, filename string) (*os.File, error) {

	code, err := GenerateCode(config)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(code)
	if err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}
