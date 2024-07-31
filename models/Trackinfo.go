package models

type Trackinfo struct {
	Track        string  `json:"track"`
	Album        string  `json:"album"`
	Artist       string  `json:"artist"`
	Length       float64 `json:"length"` // Length in milliseconds
	TimestampUTC string  `json:"timestampUTC"`
	TimestampCST string  `json:"timestampCST"`
}
