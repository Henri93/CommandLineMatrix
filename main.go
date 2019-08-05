package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gookit/color"
)

func main() {
	//TODO make width depend on size of terminal
	const WIDTH = 200
	const RESET = 40
	const DELAY = time.Second / 75

	//Handle CTRL+C Interrupt
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()

	var resetCount = RESET
	var seed = GenerateSeed(WIDTH)

	//Run until Interrupt
	for {
		color.Green.Println(CreateTransitionLevel(seed))
		resetCount--
		if resetCount < 0 {
			seed = CreateTransitionLevel(seed)
			resetCount = RESET
		}
		time.Sleep(DELAY)
	}

}

func GenerateSeed(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		if rand.Float32() < 0.65 {
			//add random spaces
			bytes[i] = byte(32)
		} else {
			bytes[i] = byte(33 + rand.Intn(90))
		}
	}
	return string(bytes)
}

func CreateTransitionLevel(base string) string {
	bytes := make([]byte, len(base))
	for i := 0; i < len(base); i++ {
		if i != 0 && base[i] == byte(32) && base[i-1] == byte(32) {
			//add space if not on boundary of space gaps
			bytes[i] = byte(32)
		} else if i > 1 && i < (len(base)-1) && base[i] == byte(32) && base[i-1] == byte(32) && base[i+1] == byte(32) {
			bytes[i] = byte(33 + rand.Intn(90))
		} else {
			if rand.Float32() < 0.25 {
				//add random spaces
				bytes[i] = byte(32)
			} else {
				bytes[i] = byte(33 + rand.Intn(90))
			}
		}
	}
	return string(bytes)
}
