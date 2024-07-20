package models

type Trackinfo struct {
	Track        string `json:"track"`
	Album        string `json:"album"`
	Artist       string `json:"artist"`
	Length       int    `json:"length"` // Length in milliseconds
	TimestampUTC string `json:"timestampUTC"`
	TimestampCST string `json:"timestampCST"`
}
