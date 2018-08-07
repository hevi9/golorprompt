package main

import (
	"github.com/lucasb-eyer/go-colorful"
	user2 "os/user"
	"log"
	"os"
)

func init() {
	SegRegister("user", "Show user",
		func() Segment { return &User{} })
}

type User struct{}

func (*User) Render() []Chunk {
	user, err := user2.Current()
	if err != nil {
		log.Print("Error: Cannot get current user: %s", err)
		return nil
	}
	username := user.Username
	//log.Printf("user.Username=%s", user.Username)
	// Don't show user if login user
	logName, exists := os.LookupEnv("LOGNAME")
	if exists {
		_, exists := os.LookupEnv("SUDO_USER")
		if ! exists {
			if logName == username {
				return nil
			}
		}
	}
	hue := 300.0*hashToFloat64([]byte(username)) + 30.0
	return []Chunk{
		{
			text: username,
			fg:   colorful.Hsv(hue, config.FgSaturationLow, config.FgValue),
		},
	}
}
