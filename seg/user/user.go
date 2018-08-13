package main

import (
	"log"
	"os"
	user2 "os/user"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {}

type User struct{}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &User{}
	return segment
}

func (*User) Render(env sys.Environment) []sys.Chunk {
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
		if !exists {
			if logName == username {
				return nil
			}
		}
	}
	hue := 300.0*sys.HashToFloat64([]byte(username)) + 30.0
	return []sys.Chunk{
		{
			Text: username,
			Fg:   colorful.Hsv(hue, sys.Config.FgSaturationLow, sys.Config.FgValue),
		},
	}
}
