package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/LeeDark/go-education/trainee-tester/test-http-server/cdbeasy"
)

// TODO: get path from ENV
const cdbDirPath = "/Users/lee/code/horisen/mnp-server/mnpserver2/test/work/etc/horisen/mnpserver/data"

// CdbSourceResponse is struct for /cdbsource endpoint response
type CdbSourceResponse struct {
	Number string `json:"number"`
	Mcc    string `json:"mcc"`
	Mnc    string `json:"mnc"`
	Ported bool   `json:"ported"`
}

func cdbSourceHandler(w http.ResponseWriter, req *http.Request) {
	// request: /cdbsource?name=akton&number=123456789
	name := req.URL.Query().Get("name")
	number := req.URL.Query().Get("number")

	// TODO: logic
	response, err := cdbSourceLogic(name, number)
	if err != nil {
		sendJSON(http.StatusInternalServerError, w, nil)
	}

	// response: JSON {"number": "123456789", "mcc": MCC, "mnc": MNC, "ported": isPorted}

	sendJSON(http.StatusOK, w, response)
}

func cdbSourceLogic(name, number string) (response *CdbSourceResponse, err error) {
	switch name {
	case "akton":
		response, err = cdbSourceSearch("akton.cdb", number)
	default:
		response = &CdbSourceResponse{
			Number: number,
		}
	}

	return
}

func cdbSourceSearch(cdbFile, number string) (response *CdbSourceResponse, err error) {
	response = &CdbSourceResponse{
		Number: number,
		Ported: false,
	}

	cdbFilePath := cdbDirPath + "/" + cdbFile

	// +12,15:385992026020->mcc=219&mnc=010
	rec, err := cdbeasy.FindOne(cdbFilePath, number)
	if err != nil || rec == "" {
		// TODO: error handling
		return
	}

	log.Println(rec)

	recMap, err := url.ParseQuery(rec)
	if err != nil {
		// TODO: error handling
		return
	}

	log.Println(recMap)

	// mcc=219&mnc=010
	// TODO: error handling
	response.Mcc = recMap["mcc"][0]
	response.Mnc = recMap["mnc"][0]
	response.Ported = true

	return
}
