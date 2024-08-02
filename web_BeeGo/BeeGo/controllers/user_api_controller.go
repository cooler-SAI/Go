package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

type UserAPIController struct {
	web.Controller
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// In-memory user storage for example purposes
var users = make(map[string]User)

func (c *UserAPIController) GetAll() {
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	c.Data["json"] = userList
	c.ServeJSON()
}

func (c *UserAPIController) Get() {
	userID := c.Ctx.Input.Param(":id")
	if user, ok := users[userID]; ok {
		c.Data["json"] = user
	} else {
		c.Data["json"] = map[string]string{"error": "User not found"}
	}
	c.ServeJSON()
}

func (c *UserAPIController) Create() {
	var user User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	users[user.ID] = user
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserAPIController) Update() {
	userID := c.Ctx.Input.Param(":id")
	var user User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if _, ok := users[userID]; ok {
		users[userID] = user
		c.Data["json"] = user
	} else {
		c.Data["json"] = map[string]string{"error": "User not found"}
	}
	c.ServeJSON()
}

func (c *UserAPIController) Delete() {
	userID := c.Ctx.Input.Param(":id")
	if _, ok := users[userID]; ok {
		delete(users, userID)
		c.Data["json"] = map[string]string{"result": "Deleted"}
	} else {
		c.Data["json"] = map[string]string{"error": "User not found"}
	}
	c.ServeJSON()
}
