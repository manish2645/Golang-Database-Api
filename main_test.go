package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"bytes"
	"testing"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {

	route := gin.Default()
	route.POST("/users", saveUser)
	
	user := User{
			ID:    "552",
			Name:  "Manish",
			Email: "manish@xenonstack.com",
			Age:18,
			Phone:"8787637060",
			State:"Meghalaya",
			City:"Shillong",
			Zip:"793002",
			Country:"India",
		}

	jsonValue, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	if err!= nil {
        fmt.Println("Error creating request:", err)
        return
    }
	
	recorder := httptest.NewRecorder()

	route.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestGetUser(t *testing.T) {

	route := gin.Default()
	route.GET("/users", getUsers)

	expectedUsers := []User{
		{
			ID:    "552",
			Name:  "Manish",
			Email: "manish@xenonstack.com",
			Age:18,
			Phone:"8787637060",
			State:"Meghalaya",
			City:"Shillong",
			Zip:"793002",
			Country:"India",
		},
	}

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil{
		fmt.Println("Error creating request",err)
		return
	}
	
	res := httptest.NewRecorder()

	route.ServeHTTP(res, req)

	// assert.Equal(t, http.StatusOK, res.Code)
	var resBody map[string][]User
	_ = json.Unmarshal(res.Body.Bytes(), &resBody)
	assert.Equal(t, expectedUsers, resBody["users"])
}

func TestMain(t *testing.T) {
	go main()
}
