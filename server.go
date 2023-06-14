package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserBody struct {
	Id   int64  `json:"id" xml:"id" form:"id" query:"id"`
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Age  int64  `json:"age" xml:"age" form:"age" query:"age"`
}

var userDb []*UserBody

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", saveUser)
	e.GET("/users", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}

func saveUser(c echo.Context) error {
	u := new(UserBody)

	if err := c.Bind(u); err != nil {
		return err
	}

	userDb = append(userDb, u)

	return c.JSON(http.StatusCreated, userDb)
}

func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, userDb)
}

func updateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Error during conversion")
		return nil
	}

	userBody := new(UserBody)

	if err := c.Bind(userBody); err != nil {
		return err
	}

	userDb = UpdateUserById(userDb, id, userBody)

	return c.JSON(http.StatusCreated, userDb)
}

func deleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Error during conversion")
		return nil
	}

	userDb = RemoveUserById(userDb, id)

	return c.NoContent(http.StatusNoContent)
}

func UpdateUserById(userList []*UserBody, id int, userBody *UserBody) []*UserBody {
	_, index := FindIndexOfUserById(userList, id)
	userList[index] = userBody
	return userList
}

func RemoveUserById(userList []*UserBody, id int) []*UserBody {
	found, index := FindIndexOfUserById(userList, id)
	if found {
		return append(userList[:index], userList[index+1:]...)
	}
	return userList
}

func FindIndexOfUserById(userList []*UserBody, id int) (bool, int) {
	index := -1
	found := false
	for i := range userList {
		if userList[i].Id == int64(id) {
			index = i
			found = true
			break
		}
	}
	return found, index
}
