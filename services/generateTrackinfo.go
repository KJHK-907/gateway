package services

import (
	"gateway/models"
	"log"
	"strconv"
)

func GenerateTrackinfo(currentMetadata models.ZettaLite) models.Trackinfo {

	if len(currentMetadata.LogEventCollection.LogEvent.Asset) == 0 {
		log.Println("Error: Asset slice is empty")
		return models.Trackinfo{}
	}
	totalLength, err := strconv.ParseFloat(currentMetadata.LogEventCollection.LogEvent.Asset[0].TotalLength, 64)
	if err != nil {
		log.Printf("Error converting TotalLength to int: %v", err)
		totalLength = 0.0
	}
	trackInfo := models.Trackinfo{
		Track:        currentMetadata.LogEventCollection.LogEvent.Asset[0].Title,
		Type:         currentMetadata.LogEventCollection.LogEvent.Type,
		Album:        currentMetadata.LogEventCollection.LogEvent.Asset[0].Album1,
		Artist:       currentMetadata.LogEventCollection.LogEvent.Asset[0].Artist1,
		Length:       totalLength * 1000,
		TimestampUTC: currentMetadata.LogEventCollection.LogEvent.StartTime,
	}
	return trackInfo
}
