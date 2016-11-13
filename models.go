package main

type talkForm struct {
	Name        string `form:"name" binding:"required"`
	Topic       string `form:"topic" binding:"required"`
	Duration    string `form:"duration" binding:"required"`
	Description string `form:"description" binding:"required"`
	Email       bool   `form:"recorded"`
}

type talkFormRes struct {
	Success bool `json:"success"`
}
