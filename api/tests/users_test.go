package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/stretchr/testify/require"
)

type Response struct {
	message string
	status  bool
	Err     bool `bson:"error" json:"error"`
	payload interface{}
}

func getResBody(resp *http.Response, t *testing.T) Response {
	var r Response
	var body []byte
	resp.Body.Read(body)
	err := json.Unmarshal(body, &r)
	if err != nil {
		t.Fatal("failed to unmarshal response body", err)
	}
	return r
}

func TestGetUsers(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/getUsers")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
}

func TestGetUserById(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/getUsers/1")
	if err != nil {
		t.Fatal(err)
	}
	resBody := getResBody(response, t)
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
	require.Equal(t, "1", resBody.payload.(models.UserStruct).UserId)
}

func TestUpdateUser(t *testing.T) {
	jsonBody := `{
		"userId": 1,
		"firstName": "Abdulkarim",
		"lastName": "Ogaji"
	}`
	reqBody := bytes.NewBuffer([]byte(jsonBody))
	req, err := http.NewRequest("PUT", "http://localhost:8000/api/v1/editUser/1", reqBody)
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
		require.Equal(t, "Abdulkarim", resBody.payload.(models.UserStruct).FirstName, "FIrstname should have updated to Abdulkarim")
	case http.StatusBadRequest:
		require.Equal(t, "Request caused no alterations to the database", resBody.message, "The rows should not have been affected")
	default:
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode, "The error should be an internal server error")
	}
}

func TestDeleteUser(t *testing.T) {
	response, err := http.Get("http://localhost:8000/api/v1/deleteUser/1")
	if err != nil {
		t.Fatal(err)
	}
	resBody := getResBody(response, t)
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")
	require.Equal(t, "user deleted successfully", resBody.message)
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
