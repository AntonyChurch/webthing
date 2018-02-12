package main

import (
	"fmt"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

var pin embd.DigitalPin

func main() {
	embd.InitGPIO()
	defer embd.CloseGPIO()

	pin, err := embd.NewDigitalPin(4)

	pin.SetDirection(embd.Out)

	if err != nil {
		fmt.Println(err)
		return
	}

	// for {
	// 	pin.Write(embd.High)
	// 	time.Sleep(time.Second)
	// 	fmt.Println("On")
	// 	pin.Write(embd.Low)
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Off")
	// }

	SetupAndListen()
}
