package main

import (
	"fmt"
	"time"
  	"os"
  	"os/signal"
  	"syscall"

	"github.com/stianeikeland/go-rpio"
)

func main() {

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
    	pin.Low()
    	os.Exit(0)
    }()

	for {
		pin.Toggle()
		time.Sleep(1 * time.Second)
	}
}
