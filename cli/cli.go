package main

import (
	"fmt"
	"os"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "b27",
				Aliases: []string{"superman"},
				Usage:   "control the zhonghong HVAC device",
				Args:    true,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
					&cli.IntFlag{Name: "address", Aliases: []string{"a"}, Required: true},
					&cli.IntFlag{Name: "temperature", Aliases: []string{"temp", "t"}},
					&cli.IntFlag{Name: "onoff", Aliases: []string{"o"}, Usage: "1: on, 0: off"},
					&cli.IntFlag{Name: "fanspeed", Aliases: []string{"f"}, Value: 0, Usage: "0: auto, 1: low, 2: medium, 3: high"},
					&cli.IntFlag{Name: "mode", Aliases: []string{"m"}, Value: 2, Usage: "1: heat, 2: cool, 4: vent, 8: dehumidify"},
					&cli.IntFlag{Name: "direction", Aliases: []string{"d"}, Value: 0, Usage: "0: no direction, 1-7: direction, 0xff: auto"},
				},

				Action: func(cCtx *cli.Context) error {
					fmt.Println(cCtx.Args())
					host := cCtx.String("host")
					port := cCtx.Int("port")
					address := cCtx.Int("address")
					temperature := cCtx.Int("temperature")
					onoff := cCtx.Bool("onoff")
					fanspeed := cCtx.Int("fanspeed")
					mode := cCtx.Int("mode")
					direction := cCtx.Int("direction")

					fmt.Printf("host: %s, port: %d, address: %d, temperature: %d, onoff: %t, fanspeed: %d, mode: %d, direction: %d\n", host, port, address, temperature, onoff, fanspeed, mode, direction)
					clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
					if err != nil {
						return err
					}
					client := client.NewB27Client(clientHandler)

					// TODO: the composition of command bytes should be done in the client package
					command := make([]byte, 5)
					if onoff {
						command[0] = 0x01
					} else {
						command[0] = 0x00
					}
					command[1] = byte(temperature)
					command[2] = byte(mode)
					command[3] = byte(fanspeed)
					command[4] = byte(direction)

					res, err := client.Control(byte(address), command...)
					if err != nil {
						return err
					}
					if res.Data[0] != 0x01 {
						return err
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
