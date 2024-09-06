package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKeyAuthEmpty(t *testing.T) {
	var headers = make(http.Header)
	headers.Add("Key", "Value")
	msg, err := GetAPIKey(headers)
	if msg != "" || err == nil {
		t.Fatalf(`GetAPIKey(headers) with no auth key resulted in %q, %v. Want "" and error`, msg, err)
	}
}

func TestGetAPIKeyMalformed(t *testing.T) {
	var headerWithWrongLength = make(http.Header)
	headerWithWrongLength.Add("Authorization", "MYREALYLONGAPIKEYBUTWITHOUT")

	var headerWithBadApiKeyPlacement = make(http.Header)
	headerWithBadApiKeyPlacement.Add("Authorization", "")

	tests := []struct {
		name          string
		header        http.Header
		exptected     string
		expectedError bool
	}{
		{"no api key prefix", headerWithWrongLength, "", true},
		{"wrong api key placement", headerWithBadApiKeyPlacement, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAPIKey(tt.header)
			if (err != nil) != tt.expectedError {
				t.Errorf("exptected error")
			}
			if result != tt.exptected {
				t.Errorf("expected: %s, got: %s", tt.exptected, result)
			}
		})
	}
}

func TestGetAPIKeyGood(t *testing.T) {
	var apiKey = "MYAPIKEYAMKNVONOIJF)#!*R)!#HR!H#SKDLFSKDF"
	var header = make(http.Header)
	header.Add("Authorization", fmt.Sprintf("ApiKey %s", apiKey))
	result, err := GetAPIKey(header)
	if result != apiKey {
		t.Errorf("Expected: %s, got: %s", apiKey, result)
	}
	if err != nil {
		t.Errorf("Exptected no error, got: %v", err)
	}
}
