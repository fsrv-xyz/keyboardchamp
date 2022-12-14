package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	evdev "github.com/gvalkov/golang-evdev"

	"github.com/fsrv-xyz/keyboardchamp/internal/action"
)

type Event struct {
	Code int
	Type int
}

var states map[int]int

func init() {
	if states == nil {
		states = make(map[int]int)
	}
}

func main() {
	var devicePath, deviceName string

	flag.StringVar(&devicePath, "device.path", "", "/dev/input/event** path to keyboard")
	flag.StringVar(&deviceName, "device.name", "SINO WEALTH Sunreed SKB886S", "name of the keyboard")
	flag.Parse()

	if devicePath == "" {
		devices, _ := evdev.ListInputDevices()
		for _, device := range devices {
			if device.Name == deviceName {
				devicePath = device.Fn
				break
			}
		}
		if devicePath == "" {
			fmt.Println("No device found")
			os.Exit(1)
		}
	}
	fmt.Printf("listening on %+q\n", devicePath)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	device, deviceOpenError := evdev.Open(devicePath)
	if deviceOpenError != nil {
		panic(deviceOpenError)
	}
	deviceGrabError := device.Grab()
	if deviceGrabError != nil {
		panic(deviceGrabError)
	}
	device.SetRepeatRate(150, 300)
	handlerInput := make(chan Event)
	stateChanged := make(chan any)

	go func() {
		for {
			event, readErr := device.ReadOne()
			if readErr != nil {
				panic(readErr)
			}
			if (event.Type == 0 || event.Type == 1) && event.Code != 0 {
				handlerInput <- Event{Code: int(event.Code), Type: int(event.Value)}
			}
		}
	}()

	go func() {
		for range stateChanged {
			for name, actionFactory := range action.RegistryInstance.Checks {
				actionInstance := actionFactory()
				if action.CheckRequirements(states, actionInstance.GetRequirements()) {
					fmt.Printf("executing action %+q\n", name)
					actionInstance.Execute()
					continue
				}
			}
		}
	}()

	go KeyHandler(handlerInput, stateChanged)
	<-sigChan
	defer fmt.Println(device.Release())
}

func KeyHandler(input chan Event, changed chan any) {
	for event := range input {
		fmt.Printf("%+v\n", event)
		states[event.Code] = event.Type
		changed <- nil
	}
}
