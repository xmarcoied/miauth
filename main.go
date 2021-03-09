package main

import (
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xmarcoied/miauth/cmd"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
}
func main() {
	os.Exit(cmd.RunServer())
}
