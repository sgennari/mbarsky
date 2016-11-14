package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/smtp"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", "homebrewstr.db")
	if err != nil {
		panic(err)
	}
}

func newTalk(c *gin.Context) {
	var talk talkForm
	if err := c.Bind(&talk); err != nil {
		fmt.Println(err)
		return
	}

	_, err := db.Exec("INSERT INTO talks(name, topic, duration, description) VALUES ($1, $2, $3, $4)", talk.Name, talk.Topic, talk.Duration, talk.Description)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, &talkFormRes{Success: false})
		return
	}

	if talk.Email {
		go emailClub(talk)
	}

	c.JSON(200, &talkFormRes{Success: true})
}

func emailClub(t talkForm) {
	subject := fmt.Sprintf("%s hosted by %s", t.Topic, t.Name)
	body := fmt.Sprintf("%s will be hosting a talk on \"%s\" for about %s minutes.\r\n\r\n Description: %s\r\n", t.Name, t.Topic, t.Duration, t.Description)
	msg := fmt.Sprintf("Subject: [Homebrew] %s\r\n", subject) +
		"From: Homebrewstr <homebrew@src-code.simons-rock>\r\n" +
		"To: CS Club <cs_club@simons-rock.edu>\r\n" +
		fmt.Sprintf("\r\n%s", body)
	smtp.SendMail("localhost:25", nil, "homebrew@src-code.simons-rock", []string{"cs_club@simons-rock.edu"}, []byte(msg))
}

func passFunc(c *gin.Context) {
	c.String(200, "Hit endpoint %s with method %s\n", c.Request.URL, c.Request.Method)
}
