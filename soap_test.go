package soap

import (
	"encoding/json"
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
		SavePushInfornationResult: "test",
		StrErrMsg:                 "test",
	}

	en := NewEnvelope(WebMethodRes)
	response, err := en.WriteEnvelope()
	if err != nil {
		t.Error(err)
	}
	// t.Log(string(response))
	needsNode := make(map[string]interface{})
	needsNode["ns"] = "http://tempuri.org/"
	needsNode["Child"] = []string{"SavePushInfornationResult", "strErrMsg"}
	needsNode["Root"] = "SavePushInfornationResponse"
	res, err := ReadEnvelope(response, needsNode)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(res))
	}
	var resmap map[string]string
	err = json.Unmarshal(res, &resmap)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resmap["SavePushInfornationResult"], resmap["strErrMsg"])
	}

}
