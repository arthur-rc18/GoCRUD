package models

import (
	"api/zeus/database"
	"api/zeus/user"
	"errors"
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

func GetAllUsers(ctx *gin.Context) ([]user.User, error) {

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	db := database.ConnectionPostgres()
	defer db.Close()

	sqlRow, err := db.Query("select * from users order by id asc")
	if err != nil {
		log.Println(err.Error())
	}

	users := user.User{}
	userSlice := []user.User{}

	for sqlRow.Next() {
		var id int
		var nome, email, senha, telefone string

		err = sqlRow.Scan(&id, &nome, &email, &senha, &telefone)
		if err != nil {
			log.Println(err.Error())
			return userSlice, err

		}

		users.Id = id
		users.Nome = nome
		users.Email = email
		users.Senha = senha
		users.Telefone = telefone

		userSlice = append(userSlice, users)
	}

	return userSlice, nil

}

func GetUserByID(ctx *gin.Context) ([]user.User, error) {

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	users := user.User{}
	userSlice := []user.User{}

	db := database.ConnectionPostgres()
	defer db.Close()

	userID := ctx.Param("id")
	var id int
	var nome, email, senha, telefone string

	err := db.QueryRow("select * from users where id=$1", userID).Scan(&id, &nome, &email, &senha, &telefone)
	users.Id = id
	users.Nome = nome
	users.Email = email
	users.Senha = senha
	users.Telefone = telefone
	userSlice = append(userSlice, users)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, err
	} else if err != nil && err.Error() != "sql: no rows in result set" {
		log.Fatal(err)
		return userSlice, err
	}

	return userSlice, nil

}

func CreateUser(ctx *gin.Context) (user.User, error) {

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	db := database.ConnectionPostgres()
	defer db.Close()

	users := user.User{}

	ctx.BindJSON(&users)

	userQuery, err := db.Prepare("insert into users(nome, email, senha, telefone) values($1, $2, $3, $4)")
	if err != nil {
		log.Println(err.Error())
		return users, err
	}
	defer userQuery.Close()

	if users.Nome == "" || users.Email == "" || users.Senha == "" || users.Telefone == "" {
		return users, errors.New("Null value")
	}

	userQuery.Exec(users.Nome, users.Email, users.Senha, users.Telefone)

	return users, nil

}

func UpdateUserByID(ctx *gin.Context) (user.User, error) {

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	db := database.ConnectionPostgres()
	defer db.Close()

	users := user.User{}
	ctx.BindJSON(&users)
	updateID := ctx.Param("id")

	updateUser, err := db.Prepare("update users set nome=$1, email=$2, senha=$3, telefone=$4 where id=$5")
	defer updateUser.Close()
	if err != nil {
		log.Println(err.Error())
		return users, err
	}

	if users.Nome == "" || users.Email == "" || users.Senha == "" || users.Telefone == "" {
		return users, errors.New("Null value")
	}

	updateUser.Exec(users.Nome, users.Email, users.Senha, users.Telefone, updateID)

	return users, nil
}

func DeleteUserByID(ctx *gin.Context) error {

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	postgresClient := database.ConnectionPostgres()

	deleteID := ctx.Param("id")

	userDeleted, err := postgresClient.Prepare("delete from users where id = $1")
	defer userDeleted.Close()

	if err != nil {
		log.Println(err.Error())
		return err

	}

	sqlResult, err := userDeleted.Exec(deleteID)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if idNonExistant, _ := sqlResult.RowsAffected(); idNonExistant == 0 {
		return errors.New("Non existant ID")
	}

	defer postgresClient.Close()

	return nil

}

func PartiallyUpdatingUser(ctx *gin.Context) (user.User, error) {

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	db := database.ConnectionPostgres()
	defer db.Close()

	users := user.User{}
	ctx.BindJSON(&users)
	updateID := ctx.Param("id")

	updateUser, err := db.Prepare("update users set nome=$1, email=$2, senha=$3, telefone=$4 where id=$5")
	defer updateUser.Close()
	if err != nil {
		log.Println(err.Error())
		return users, err
	}

	updateUser.Exec(users.Nome, users.Email, users.Senha, users.Telefone, updateID)

	return users, nil

}
