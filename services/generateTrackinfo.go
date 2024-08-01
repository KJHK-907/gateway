package services

import (
	"gateway/models"
	"log"
	"strconv"
	"time"
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
	startTimeLocal, err := time.Parse("1/2/2006 3:04:05 PM", currentMetadata.LogEventCollection.LogEvent.StartTimeLocal)
	if err != nil {
		log.Printf("Error parsing StartTimeLocal: %v", err)
		startTimeLocal = time.Now()
	}
	timestampCST := startTimeLocal.Format("2006-01-02T15:04:05Z")
	trackInfo := models.Trackinfo{
		Track:        currentMetadata.LogEventCollection.LogEvent.Asset[0].Title,
		Album:        currentMetadata.LogEventCollection.LogEvent.Asset[0].Album1,
		Artist:       currentMetadata.LogEventCollection.LogEvent.Asset[0].Artist1,
		Length:       totalLength * 1000,
		TimestampUTC: currentMetadata.LogEventCollection.LogEvent.StartTime,
		TimestampCST: timestampCST,
	}
	return trackInfo
}
