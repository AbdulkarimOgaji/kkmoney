package tests

import (
	"bytes"
	"net/http"
	"testing"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/stretchr/testify/require"
)

type PayloadTypeAcct struct {
	account models.AcctStruct
}
type PayloadTypeAccts struct {
	acctounts []models.AcctStruct
}

func TestCreateAcct(t *testing.T) {
	jsonBody := `{
	"userId": 2,
	"acctType": "C"
	}`
	reqBody := bytes.NewBuffer([]byte(jsonBody))
	resp, err := http.Post("http://localhost:8000/api/v1/createAcct", "application/json", reqBody)
	if err != nil {
		t.Fatal("Failed to send request: ", err)
	}
	require.Equal(t, 201, resp.StatusCode, "You should get a 201 response")

}

func TestGetAccounts(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/getAccts")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
}

func TestGetAcctById(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/getAccts/1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
}

func TestUpdateAcct(t *testing.T) {
	jsonBody := `{
		"acctType": "C"
	}`
	reqBody := bytes.NewBuffer([]byte(jsonBody))
	req, err := http.NewRequest("PUT", "http://localhost:8000/api/v1/editAcct/1", reqBody)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal("Failed to prepare request", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 200, resp.StatusCode, "The update should have been a success")
}

func TestDeleteAcct(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8000/api/v1/deleteAcct/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
}
