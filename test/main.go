package main

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

var (
	users = map[int]*User{}
	count = 1
)

func main() {
	e := echo.New()

	e.POST("/users", func(c echo.Context) error {
		u := &User{
			ID:	count,
		}
		if err := c.Bind(u); err != nil {
			return err
		}
		users[u.ID] = u
		count++
		return c.JSON(http.StatusCreated, u)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, users[id])
	})

	e.PATCH("/users/:id", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		id, _ := strconv.Atoi(c.Param("id"))
		users[id].Name = u.Name
		return c.JSON(http.StatusOK, users[id])
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		delete(users, id)
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":2205"))
}
