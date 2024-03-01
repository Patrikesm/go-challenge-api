package main

import (
	"bytes"
	"challenge-week-one/controllers"
	"challenge-week-one/database"
	"challenge-week-one/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	return r
}

func CreateMockDeclaration() {
	declaration := models.Declaration{
		Author:    "Mock author",
		Image:     "Mock image",
		Testimony: "Mock Testimony",
	}

	database.DB.Create(&declaration)

	ID = int(declaration.ID)

	fmt.Println("Criado aqui", ID)
}

func DeleteMockDelcaration() {
	var declaration models.Declaration

	fmt.Println("Deletado aqui", ID)

	database.DB.Delete(&declaration, ID)
}

func TestGetAllFunction(t *testing.T) {
	database.ConnectDB()

	CreateMockDeclaration()

	r := SetupRouter()

	r.GET("/depoimentos", controllers.GetAll)

	req, _ := http.NewRequest("GET", "/depoimentos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockDelcaration()
}

func TestGetRandomDeclaration(t *testing.T) {
	database.ConnectDB()
	CreateMockDeclaration()

	r := SetupRouter()
	r.GET("/depoimentos-home", controllers.GetRandomDeclaration)

	req, _ := http.NewRequest("GET", "/depoimentos-home", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockDelcaration()
}

func TestUpdateDeclaration(t *testing.T) {
	var declaration = models.Declaration{
		Author:    "Change Mock Declaration",
		Image:     "Change Mock Image",
		Testimony: "Change Mock Testimony",
	}

	database.ConnectDB()

	CreateMockDeclaration()

	r := SetupRouter()
	r.PATCH("/depoimentos/:id", controllers.Update)

	encondedDeclaration, _ := json.Marshal(declaration)

	req, _ := http.NewRequest("PATCH", "/depoimentos/"+strconv.Itoa(ID), bytes.NewBuffer(encondedDeclaration))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	fmt.Println("depois de atualizar", ID)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockDelcaration()
}

func TestDeleteMockDeclaration(t *testing.T) {
	database.ConnectDB()
	CreateMockDeclaration()

	r := SetupRouter()
	r.DELETE("/depoimentos/:id", controllers.Delete)

	req, _ := http.NewRequest("DELETE", "/depoimentos/"+strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockDelcaration()
}

func TestCreateMockDeclaration(t *testing.T) {
	database.ConnectDB()

	var declaration = models.Declaration{
		Author: "Create Mock Author",
	}

	encondedDeclaration, _ := json.Marshal(declaration)

	r := SetupRouter()
	r.POST("/depoimentos", controllers.Create)

	req, _ := http.NewRequest("POST", "/depoimentos", bytes.NewBuffer(encondedDeclaration))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var body models.Declaration

	json.Unmarshal(res.Body.Bytes(), &body)

	ID = int(body.ID)

	fmt.Println(ID)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockDelcaration()
}
