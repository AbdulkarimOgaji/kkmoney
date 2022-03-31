package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/stretchr/testify/require"
)

type PayloadTypeUser struct {
	user models.UserStruct
}
type PayloadTypeUsers struct {
	user models.UserStruct
}

type ResponseOne struct {
	message string
	status  bool
	Err     bool `bson:"error" json:"error"`
	payload PayloadTypeUser
}
type ResponseMany struct {
	message string
	status  bool
	Err     bool `bson:"error" json:"error"`
	payload PayloadTypeUsers
}

func getResBodyOne(resp *http.Response, t *testing.T) ResponseOne {
	defer resp.Body.Close()
	var r ResponseOne
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(bodyBytes, &r)
	if err != nil {
		t.Fatal("failed to unmarshal response body", err)
	}
	return r
}
func getResBodyMany(resp *http.Response, t *testing.T) ResponseMany {
	defer resp.Body.Close()
	var r ResponseMany
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(bodyBytes, &r)
	if err != nil {
		t.Fatal("failed to unmarshal response body", err)
	}
	return r
}

func TestCreateUser(t *testing.T) {
	jsonBody := `{
	"firstName": "Abba",
	"lastName": "Sadiq",
	"otherName": "Yunusa",
	"email": "exapmle@udusok.edu.ng",
	"phoneNum": "07039666042",
	"otherNum": "",
	"gender": "M",
	"address": "6 Edmund Crescent NIMR yaba, lagos, Nigeria",
	"kinName": "Yunusa Ogaji",
	"kinNumber": "08055603698",
	"kinRelationship": "Father",
	"passwordHash": "yusufOgaji"
	}`
	reqBody := bytes.NewBuffer([]byte(jsonBody))
	resp, err := http.Post("http://localhost:8000/api/v1/createUser", "application/json", reqBody)
	if err != nil {
		t.Fatal("Failed to send request: ", err)
	}
	require.Equal(t, 201, resp.StatusCode, "You should get a 201 response")

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
	require.Equal(t, 200, response.StatusCode, "The status code should be 200")

}

func TestUpdateUser(t *testing.T) {
	jsonBody := `{
	"userId": 1,
	"firstName": "Abdulkarim",
	"lastName": "Ogaji",
	"otherName": "Yunusa",
	"email": "abdulkarimogaji002@gmail.com",
	"phoneNum": "08166629550",
	"otherNum": "",
	"gender": "M",
	"address": "6 Edmund Crescent NIMR yaba, lagos, Nigeria",
	"kinName": "Ahmed Ogaji",
	"kinNumber": "08036281855",
	"kinRelationship": "Brother"
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
	resBody := getResBodyOne(resp, t)
	switch resp.StatusCode {
	case 200:
		require.Equal(t, "Abdulkarim", resBody.payload.user.FirstName, "FIrstname should have updated to Fatima")
	case http.StatusBadRequest:
		require.Equal(t, "Request caused no alterations to the database", resBody.message, "The rows should not have been affected")
	default:
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode, "The error should be an internal server error")
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8000/api/v1/deleteUser/1", nil)
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
