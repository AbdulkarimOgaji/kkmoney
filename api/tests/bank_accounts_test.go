package tests

import (
	"bytes"
	"net/http"
	"testing"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/stretchr/testify/require"
)

func TestGetAccounts(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/getAccts")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
}

func TestGetAcctById(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/getaccts/1")
	if err != nil {
		t.Fatal(err)
	}
	resBody := getResBody(response, t)
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
	require.Equal(t, "1", resBody.payload.(models.AcctStruct).AcctId)
}

func TestUpdateAcct(t *testing.T) {
	jsonBody := `{
		"acctType": "S",
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
	resBody := getResBody(resp, t)
	switch resp.StatusCode {
	case 200:
		require.Equal(t, 'S', resBody.payload.(models.AcctStruct).AcctType, "Account type should have been updated to S")
	case http.StatusBadRequest:
		require.Equal(t, "Request caused no alterations to the database", resBody.message, "The rows should not have been affected")
	default:
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode, "The error should be an internal server error")
	}
}

func TestDeleteAcct(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/deleteAcct/1")
	if err != nil {
		t.Fatal(err)
	}
	resBody := getResBody(response, t)
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
	require.Equal(t, "account deleted successfully", resBody.message)
}

// func TestCreateUser(t *testing.T) {
// 	jsonBody := `{
// 		"userId": 1,
// 		"firstName": "Abdulkarim",
// 		"lastName": "Ogaji"
// 	}`
// 	reqBody := bytes.NewBuffer([]byte(jsonBody))
// 	resp, err := http.Post("http://localhost:8000/api/v1/createUser", "application/json", reqBody)
// }
