package services

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"gateway/models"
)

var (
	recentTrackInfo = make(map[string]time.Time)
	mu              sync.Mutex
)

func ProcessXml(data []byte, buffer *strings.Builder, pool *models.Pool) {
	buffer.Write(data)

	for strings.Contains(buffer.String(), "</ZettaLite>") {
		start := strings.Index(buffer.String(), "<ZettaLite")
		end := strings.Index(buffer.String(), "</ZettaLite>") + len("</ZettaLite>")
		if start == -1 || end == -1 {
			break // In case of malformed XML
		}
		xmlDocument := buffer.String()[start:end]

		// Parse the extracted document
		currentMetadata, err := ParseXml([]byte(xmlDocument))
		if err != nil {
			fmt.Println("Error parsing XML:", err)
			// Optionally, clear the buffer if the document is corrupt and cannot be parsed
			buffer.Reset()
			continue
		}
		trackInfo := GenerateTrackinfo(currentMetadata)
		// Check if trackinfo is empty
		if trackInfo == (models.Trackinfo{}) {
			buffer.Reset()
			continue
		}
		// Check if trackinfo is in the recentTrackInfo map
		mu.Lock()
		if _, ok := recentTrackInfo[trackInfo.Track]; ok {
			mu.Unlock()
			buffer.Reset()
			continue
		}
		recentTrackInfo[trackInfo.Track] = time.Now()
		mu.Unlock()

		pool.Broadcast <- trackInfo
		log.Println("Received metadata from Zetta RCS:")
		fmt.Printf("%+v\n", trackInfo)
		println("--------------------------")

		// Remove the parsed document from the buffer
		buffer.Reset()
	}
}

func CleanupOldEntries() {
	for {
		time.Sleep(time.Hour) // Run cleanup every hour

		mu.Lock()
		for track, timestamp := range recentTrackInfo {
			if time.Since(timestamp) > time.Hour { // Clear entries older than 1 hour
				delete(recentTrackInfo, track)
			}
		}
		mu.Unlock()
	}
}
