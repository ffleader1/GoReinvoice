package inputdata

import (
	"GoReinvoice/internal/utils"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func ReadData(dataFile string) (PdfInput, error) {
	content, err := os.ReadFile(dataFile)
	if err != nil {
		return PdfInput{}, err
	}

	var payload PdfInput
	if strings.HasSuffix(dataFile, ".json") {
		err = json.Unmarshal(content, &payload)
		if err != nil {
			return PdfInput{}, err
		}
	} else {
		err = yaml.Unmarshal(content, &payload)
		if err != nil {
			return PdfInput{}, err
		}
	}

	dir := filepath.Dir(dataFile)
	if payload.Resource != "" {
		if utils.IsRelativePath(payload.Resource) {
			dir = filepath.Join(dir, payload.Resource)
		} else {
			dir = filepath.Dir(payload.Resource)
		}
	}

	for k, v := range payload.Files {
		if utils.IsRelativePath(v.DataURL) {
			v.DataURL = filepath.Join(dir, v.DataURL)
			payload.Files[k] = v
		}
	}

	for k, v := range payload.Fonts {
		if utils.IsRelativePath(v.DataURL) {
			v.DataURL = filepath.Join(dir, v.DataURL)
			payload.Fonts[k] = v
		}
	}

	return payload, nil
}
