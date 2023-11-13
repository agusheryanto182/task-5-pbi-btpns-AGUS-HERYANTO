package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int()

	database.InitDB()

	api := "/api/v1"
	r := router.RouteInit()

	randUsername := fmt.Sprintf("Hafif Nur Muhammad %d", randNum)
	randEmail := fmt.Sprintf("hafifcyber%d@gmail.com", randNum)
	newUser := models.User{
		Username: randUsername,
		Email:    randEmail,
		Password: "rootuser",
	}
	jsonValue, _ := json.Marshal(newUser)

	req, _ := http.NewRequest("POST", api+"/users/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
