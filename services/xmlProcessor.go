package services

import (
	"strings"

	"fmt"
)

func ProcessXml(data []byte, buffer *strings.Builder) {
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
		MetadataChannel <- currentMetadata
		println("Received metadata from Zetta RCS:")
		fmt.Printf("%+v\n", currentMetadata)
		println("--------------------------")

		// Remove the parsed document from the buffer
		substring := buffer.String()[end:]
		buffer = &strings.Builder{}
		buffer.WriteString(substring)
	}
}
