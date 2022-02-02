package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TestSuite struct {
	Tests struct {
		Env struct {
			Dev struct {
				EndpointInfo string `json:"endpointInfo"`
				EndpointURL  string `json:"endpointURL"`
			} `json:"dev"`
		} `json:"env"`
	} `json:"tests"`
}

func marshalTests() TestSuite {
	var t TestSuite
	f, err := os.Open("src/tests/tests.json")
	if err != nil {
		log.Fatal(err)
	}
	byte_value, error := ioutil.ReadAll(f)
	if error != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byte_value, &t)
	return t
}

func Test(t TestSuite) int {

	resp, err := http.Get(t.Tests.Env.Dev.EndpointURL)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.StatusCode
}

func ReturnResults() int {
	return Test(marshalTests())
}
