package models

import "encoding/xml"

type ZettaLite struct {
	XMLName            xml.Name           `xml:"ZettaLite"`
	Version            string             `xml:"Version,attr"`
	LogEventCollection LogEventCollection `xml:"LogEventCollection"`
}

type LogEventCollection struct {
	XMLName  xml.Name `xml:"LogEventCollection"`
	LogEvent LogEvent `xml:"LogEvent"`
}

type LogEvent struct {
	XMLName        xml.Name `xml:"LogEvent"`
	Type           string   `xml:"Type,attr"`
	ScheduledTime  string   `xml:"ScheduledTime,attr"`
	StartTime      string   `xml:"StartTime,attr"`
	StartTimeLocal string   `xml:"StartTimeLocal,attr"`
	Chain          string   `xml:"Chain,attr"`
	Status         string   `xml:"Status.attr"`
	Duration       string   `xml:"Duration,attr"`
	EditCode       string   `xml:"EditCode,attr"`
	LastStarted    string   `xml:"LastStarted,attr"`
	ZettaId        string   `xml:"ZettaId,attr"`
	Asset          []Asset  `xml:"Asset"`
}

type Asset struct {
	XMLName       xml.Name `xml:"Asset"`
	Type          string   `xml:"Type,attr"`
	AssetTypeName string   `xml:"AssetTypeName,attr"`
	Title         string   `xml:"Title,attr"`
	Comment       string   `xml:"Comment,attr"`
	ZettaId       string   `xml:"ZettaId,attr"`
	ThirdPartyId  string   `xml:"ThirdPartyId,attr"`
	File          string   `xml:"File,attr"`
	TotalLength   string   `xml:"TotalLength,attr"`
	Artist1       string   `xml:"Artist1,attr"`
	Album1        string   `xml:"Album1,attr"`
	Category      string   `xml:"Category,attr"`
	TrimIn        string   `xml:"TrimIn,attr"`
	TrimOut       string   `xml:"TrimOut,attr"`
}
