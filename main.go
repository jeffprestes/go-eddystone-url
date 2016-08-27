package main

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/suapapa/go_eddystone"
)

const (
	advTypeAllUUID16           = 0x03
	advTypeServiceData16       = 0x16
	advFlagGeneralDiscoverable = 0x02
	advFlagLEOnly              = 0x04
)

func main() {
	f, err := eddystone.MakeURLFrame("http://bit.ly/1Ld89Mg", -4)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)

	advPkt := &gatt.AdvPacket{}
	advPkt.AppendFlags(advFlagGeneralDiscoverable | advFlagLEOnly)
	advPkt.AppendField(advTypeAllUUID16, eddystone.SvcUUIDBytes)
	advPkt.AppendField(advTypeServiceData16, append(eddystone.SvcUUIDBytes, f...))

	fmt.Println(advPkt.Len(), advPkt)

	device, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	device.Handle(
		gatt.CentralConnected(
			func(c gatt.Central) {
				fmt.Println("Connect: ", c.ID())
			}),
		gatt.CentralDisconnected(
			func(c gatt.Central) {
				fmt.Println("Disconnect: ", c.ID())
			}),
	)

	onStateChanged := func(device gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			device.Advertise(advPkt)
		default:
			log.Println(s)
		}
	}

	device.Init(onStateChanged)
	select {}
}
