package services

import (
	"encoding/xml"
	"fmt"

	"gateway/models"
)

func ParseXml(data []byte) (models.ZettaLite, error) {
	currentMetadata := models.ZettaLite{}
	err := xml.Unmarshal(data, &currentMetadata)
	if err != nil {
		fmt.Println(err)
		return models.ZettaLite{}, err
	}
	return currentMetadata, nil
}
