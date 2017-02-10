package soap

import (
	"bytes"
	"encoding/xml"
	"log"
)

//Envelope defines the soap envelope column ,
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	// Soapenc       string   `xml:"xmlns:soapenc,attr"`
	Xsd string `xml:"xmlns:xsd,attr"`
	// EncodingStyle string   `xml:"soap:encodingStyle,attr"`
	Soap string `xml:"xmlns:soap,attr"`
	Body Body
}

//Body defines the customsized object .
type Body struct {
	XMLName xml.Name `xml:"Body"`
	Data    string   `xml:",innerxml"`
}

//NewEnvelope init a soap structure.
func NewEnvelope(data interface{}) Envelope {
	msg, err := xml.Marshal(data)
	if err != nil {
		panic(err)
	}
	return Envelope{
		Xsi: "http://www.w3.org/2001/XMLSchema-instance",
		// Soapenc:       "http://schemas.xmlsoap.org/soap/encoding/",
		Xsd: "http://www.w3.org/2001/XMLSchema",
		// EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Soap: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: Body{Data: string(msg)},
	}
}

//WriteEnvelope write the envelope to the buffer .
func (env *Envelope) WriteEnvelope() ([]byte, error) {
	msg, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	buf.Write(msg)
	return buf.Bytes(), nil
}

//ReadEnvelope decode the soap structure .
func (env *Envelope) ReadEnvelope(responseBody []byte) error {

	err := xml.Unmarshal([]byte(responseBody), &env)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
