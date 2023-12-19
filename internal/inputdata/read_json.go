package inputdata

import (
	"GoReinvoice/internal/utils"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func ReadData(jsonFile string) PdfInput {
	content, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload PdfInput
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	dir := filepath.Dir(jsonFile)
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

	return payload
}
