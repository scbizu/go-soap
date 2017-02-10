package soap

import (
	"encoding/xml"
	"testing"
)

func TestWriteEnvelop(t *testing.T) {
	WebMethod := struct {
		XMLName  xml.Name `xml:"http://tempuri.org/ SavePushInfornation"`
		UserName string   `xml:"userName"`
		StrPwd   string   `xml:"strPwd"`
		StrJSON  string   `xml:"strJson"`
	}{
		UserName: "123",
		StrPwd:   "123",
		StrJSON:  "123",
	}

	en := NewEnvelope(WebMethod)
	res, err := en.WriteEnvelope()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}

func TestReadEnvelop(t *testing.T) {

	WebMethodRes := struct {
		XMLName                   xml.Name `xml:"http://tempuri.org/ SavePushInfornationResponse"`
		SavePushInfornationResult string   `xml:"SavePushInfornationResult"`
		StrErrMsg                 string   `xml:"strErrMsg"`
	}{
		SavePushInfornationResult: "123",
		StrErrMsg:                 "123",
	}

	en := NewEnvelope(WebMethodRes)
	response, err := en.WriteEnvelope()
	if err != nil {
		t.Error(err)
	}
	err = en.ReadEnvelope(response)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(en.Body.Data)
	}

	err = xml.Unmarshal([]byte(en.Body.Data), &WebMethodRes)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(WebMethodRes.SavePushInfornationResult, WebMethodRes.StrErrMsg)
	}

}
