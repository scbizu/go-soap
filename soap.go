package soap

import (
	"bytes"
	"encoding/json"
	"encoding/xml"

	xmlx "github.com/jteeuwen/go-pkg-xmlx"
)

//Envelope defines the soap envelope column ,
type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	// Soapenc       string   `xml:"xmlns:soapenc,attr"`
	Xsd string `xml:"xmlns:xsd,attr"`
	// EncodingStyle string   `xml:"soap:encodingStyle,attr"`
	SoapNS string `xml:"xmlns:soap,attr"`
	Body   Body
}

//Body defines the customsized object .
type Body struct {
	XMLName xml.Name `xml:"soap:Body"`
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
		SoapNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:   Body{Data: string(msg)},
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
func ReadEnvelope(responseBody []byte, NeedNodes map[string]interface{}) ([]byte, error) {
	doc := xmlx.New()

	if err := doc.LoadString(string(responseBody), nil); err != nil {
		return nil, err
	}

	res := make(map[string]string)
	rootNode := doc.SelectNode(NeedNodes["ns"].(string), NeedNodes["Root"].(string))
	for _, neednode := range NeedNodes["Child"].([]string) {

		s := rootNode.S(NeedNodes["ns"].(string), neednode)
		res[neednode] = s
	}
	resjson, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	return resjson, nil
}
