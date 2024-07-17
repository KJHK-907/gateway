package services

import (
	"testing"
)

func TestParseXml(t *testing.T) {
	var tests = []struct {
		name            string
		input           []byte
		expectedError   bool
		expectedVersion string
	}{
		{"Valid XML", []byte(`
			<ZettaLite Version="0.0.0">
				<LogEventCollection>
					<LogEvent Type="Music" ScheduledTime="2023-04-01T12:00:00Z" StartTime="2023-04-01T12:00:00Z" StartTimeLocal="2023-04-01T08:00:00-04:00" Chain="1" Status="Completed" Duration="180" EditCode="" LastStarted="2023-04-01T12:00:00Z" ZettaId="123">
						<Asset Type="Song" AssetTypeName="Music" Title="Sample Song" Comment="" ZettaId="456" ThirdPartyId="789" File="sample.mp3" TotalLength="180" Artist1="Sample Artist" Album1="Sample Album" Category="Pop" TrimIn="0" TrimOut="180"/>
					</LogEvent>
				</LogEventCollection>
			</ZettaLite>
			`), false, "0.0.0"},
		{"Invalid XML", []byte("<metadata>"), true, ""},
		{"Empty XML", []byte(""), true, ""},
		{"Invalid XML 2", []byte(`
			<ZettaLite Version="0.0.0">
			`), true, ""},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentMetadata, err := ParseXml(tt.input)
			if (err != nil) != tt.expectedError {
				t.Errorf("parseXml() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if !tt.expectedError && currentMetadata.Version != tt.expectedVersion {
				t.Errorf("Expected Version '%s', got '%s'", tt.expectedVersion, currentMetadata.Version)
			}
		})
	}
}
