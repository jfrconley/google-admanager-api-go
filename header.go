package admanager

import "encoding/xml"

// RequestHeader is the SOAP request header required by the Google Ad Manager API.
// The XML namespace is set dynamically based on the API version.
type RequestHeader struct {
	XMLName         xml.Name
	NetworkCode     string `xml:"networkCode,omitempty"`
	ApplicationName string `xml:"applicationName,omitempty"`
}

// NewRequestHeader creates a RequestHeader with the XML namespace set to the
// given API version (e.g. "v202505").
func NewRequestHeader(version, networkCode, applicationName string) *RequestHeader {
	return &RequestHeader{
		XMLName: xml.Name{
			Space: "https://www.google.com/apis/ads/publisher/" + version,
			Local: "RequestHeader",
		},
		NetworkCode:     networkCode,
		ApplicationName: applicationName,
	}
}
