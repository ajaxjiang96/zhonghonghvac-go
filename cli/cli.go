package main

import (
	"fmt"
	"os"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
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
			}, {
				Name:    "b19",
				Aliases: []string{"vrf"},
				Usage:   "control the zhonghong HVAC vrf controller device",
				Subcommands: []*cli.Command{
					{
						Name: "status",
						Args: true,
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
							&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
							&cli.IntFlag{Name: "address", Aliases: []string{"a"}, Required: true},
						},

						Action: func(cCtx *cli.Context) error {
							fmt.Println(cCtx.Args())
							host := cCtx.String("host")
							port := cCtx.Int("port")
							address := cCtx.Int("address")

							fmt.Printf("host: %s, port: %d, address: %d \n", host, port, address)
							clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B19Packager{})
							if err != nil {
								return err
							}
							client := client.NewB19Client(clientHandler)

							res, err := client.AllACStatus()
							if err != nil {
								return err
							}
							command := res.Data[0]
							if command != 0xFF {
								panic(fmt.Errorf("invalid command in response %x", command))
							}
							fmt.Printf("res: %x \n", res)
							numDevices := uint(res.Data[1])
							if numDevices == 0 {
								panic(fmt.Errorf("gateway not ready yet"))
							}
							fmt.Printf("numDevices: %d \n", numDevices)
							devicesInfo := res.Data[2:]
							for i := 0; i < int(numDevices); i++ {
								device := devicesInfo[i*10 : i*10+10]
								fmt.Printf("device %d-%d:\ton: %x, temp: %d, mode: %x, fanspeed: %x, room_temp: %d, errorcode: %x, direction: %x, is_slave: %x \n", device[0], device[1], device[2], device[3], protocol.ACMode(device[4]), protocol.FanSpeed(device[5]), device[6], device[7], device[8], device[9])
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
