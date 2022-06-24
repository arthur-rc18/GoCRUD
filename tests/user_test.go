package tests

import (
	"api/zeus/database"
	"api/zeus/handlers"
	"api/zeus/user"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestUsers struct {
	Id       int
	Nome     string
	Email    string
	Senha    string
	Telefone string
}

var validUsers = []TestUsers{
	{Nome: "alice", Email: "alicea@gmail.com", Senha: "123456", Telefone: "219717418"},
	{Nome: "paulo", Email: "paulo@gmail.com", Senha: "123456", Telefone: "21992716541"},
	{Nome: "arthur", Email: "atr2@gmail.com", Senha: "123456", Telefone: "51987197546"},
}

var ID int

// func TestGetAllUsers(t *testing.T) {

// 	var ctx *gin.Context
// 	got, err := models.GetAllUsers(ctx)
// 	t.Parallel()

// 	for _, users := range validUsers {
// 		t.Run("GetUsersTest", func(t *testing.T) {
// 			assert.Equal(t, err, nil)
// 			assert.Equal(t, got, users)
// 		})
// 	}
// }

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateUserMock() {

	user := user.User{Id: 0, Nome: "User test", Email: "user@gmail.com", Senha: "123456", Telefone: "34979401823"}
	db := database.ConnectionPostgres()
	defer db.Close()

	userQuery, err := db.Prepare("insert into users(id, nome, email, senha, telefone) values($1, $2, $3, $4, $5)")
	if err != nil {
		log.Println(err.Error())
	}
	defer userQuery.Close()

	userQuery.Exec(user.Id, user.Nome, user.Email, user.Senha, user.Telefone)

}

func DeleteUserMock() {
	postgresClient := database.ConnectionPostgres()
	defer postgresClient.Close()

	userDeleted, err := postgresClient.Query("delete from users where id = 0")
	defer userDeleted.Close()

	if err != nil {
		log.Println(err.Error())
		return

	}

}

func TestGetAllUsers(t *testing.T) {

	database.ConnectionPostgres()
	CreateUserMock()
	defer DeleteUserMock()
	t.Parallel()

	routes := RoutesSetup()
	routes.GET("/users", handlers.GetAllUsers)
	req, _ := http.NewRequest("GET", "/users", nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}

func TestGetUserByID(t *testing.T) {

	database.ConnectionPostgres()
	CreateUserMock()
	defer DeleteUserMock()
	t.Parallel()

	routes := RoutesSetup()
	routes.GET("/users/:id", handlers.GetUserByID)
	path := "/users/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	var users user.User
	json.Unmarshal(res.Body.Bytes(), &users)
	fmt.Println(&users)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestDeletUser(t *testing.T) {

	database.ConnectionPostgres()
	CreateUserMock()
	t.Parallel()

	routes := RoutesSetup()
	routes.DELETE("/users/:id", handlers.DeleteUserByID)
	path := "/users/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestCreateUser(t *testing.T) {

	database.ConnectionPostgres()
	CreateUserMock()
	defer DeleteUserMock()
	t.Parallel()

	routes := RoutesSetup()
	routes.POST("/users/", handlers.CreateUser)
	req, _ := http.NewRequest("POST", "/users", nil)
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}

func TestUpdateUserByID(t *testing.T) {

	database.ConnectionPostgres()
	CreateUserMock()
	defer DeleteUserMock()
	t.Parallel()

	user := user.User{Id: 0, Nome: "User test", Email: "user@gmail.com", Senha: "123456", Telefone: "34979401823"}

	routes := RoutesSetup()
	routes.PUT("/users/:id", handlers.UpdateUserByID)
	jsonValue, _ := json.Marshal(&user)
	path := "/users/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(jsonValue))
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestPartiallyUpdateUserByID(t *testing.T) {

	database.ConnectionPostgres()
	CreateUserMock()
	defer DeleteUserMock()
	t.Parallel()

	user := user.User{Id: 0, Nome: "User test", Email: "user@gmail.com", Senha: "123456", Telefone: "34979401823"}

	routes := RoutesSetup()
	routes.PATCH("/users/:id", handlers.PartiallyUpdatingUser)
	jsonValue, _ := json.Marshal(&user)
	path := "/users/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonValue))
	res := httptest.NewRecorder()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
