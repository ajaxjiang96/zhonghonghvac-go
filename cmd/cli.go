package main

import (
	"fmt"
	"os"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b19"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b27"
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
					&cli.StringFlag{Name: "address", Aliases: []string{"a"}, Required: true},
				},

				Subcommands: []*cli.Command{
					{ // performance
						Name: "performance",
						Action: func(cCtx *cli.Context) error {
							host := cCtx.String("host")
							port := cCtx.Int("port")
							address := cCtx.String("address")

							clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
							if err != nil {
								return err
							}
							client := b27.NewClient(clientHandler)
							performance, err := client.PerformanceCheck(address)
							if err != nil {
								panic(err)
							}

							fmt.Printf("addr: %s, brand: %d, status: %d \n", performance.Addr, performance.ACBrand, performance.Status)
							return nil
						},
					},
					{ // status
						Name: "status",
						Action: func(cCtx *cli.Context) error {
							host := cCtx.String("host")
							port := cCtx.Int("port")
							address := cCtx.String("address")

							clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
							if err != nil {
								return err
							}
							client := b27.NewClient(clientHandler)

							status, err := client.StatusCheck(address)
							if err != nil {
								return err
							}

							fmt.Println(status)
							return nil
						},
					},
					{ // control
						Name: "control",
						Args: true,
						Flags: []cli.Flag{
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
							address := cCtx.String("address")
							temperature := cCtx.Uint("temperature")
							onoff := cCtx.Bool("onoff")
							fanspeed := cCtx.Int("fanspeed")
							mode := cCtx.Int("mode")
							direction := cCtx.Int("direction")

							fmt.Printf("host: %s, port: %d, address: %s , temperature: %d, onoff: %t, fanspeed: %d, mode: %d, direction: %d\n", host, port, address, temperature, onoff, fanspeed, mode, direction)
							clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
							if err != nil {
								return err
							}
							client := b27.NewClient(clientHandler)
							command := protocol.ACControlRequest{}

							if cCtx.IsSet("onoff") {
								command.On = &onoff
							}
							if cCtx.IsSet("temperature") {
								command.Temp = &temperature
							}
							if cCtx.IsSet("fanspeed") {
								fanSpeedEnum := protocol.FanSpeed(byte(fanspeed))
								command.FanSpeed = &fanSpeedEnum
							}
							if cCtx.IsSet("mode") {
								modeEnum := protocol.ACMode(byte(mode))
								command.Mode = &modeEnum
							}
							if cCtx.IsSet("direction") {
								dirEnum := protocol.ACWindDir(direction)
								command.Direction = &dirEnum
							}

							res, err := client.Control(address, command)
							if err != nil {
								return err
							}
							fmt.Sprintln(res)
							return nil
						},
					},
				},
			},
			{
				Name:    "b19",
				Aliases: []string{"vrf"},
				Usage:   "control the zhonghong HVAC vrf controller device",
				Subcommands: []*cli.Command{
					{ // status
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
			{
				Name:    "performance",
				Aliases: []string{"perf"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
					&cli.StringFlag{Name: "address", Aliases: []string{"a"}, Required: true},
					&cli.StringFlag{Name: "device", Aliases: []string{"d"}, Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					host := cCtx.String("host")
					port := cCtx.Int("port")
					address := cCtx.String("address")
					device := cCtx.String("device")

					if device != "b27" && device != "b19" {
						panic(fmt.Errorf("invalid device %s", device))
					}
					var client api.ClientV2

					if device == "b27" {
						clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
						if err != nil {
							return err
						}
						client = b27.NewClient(clientHandler)
					} else {
						clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B19Packager{})
						if err != nil {
							return err
						}
						client = b19.NewClient(clientHandler)
					}

					performance, err := client.PerformanceCheck(address)
					if err != nil {
						panic(err)
					}

					fmt.Printf("addr: %s, brand: %d, status: %d \n", performance.Addr, performance.ACBrand, performance.Status)
					return nil
				},
			},
			{
				Name:    "performance-batch",
				Aliases: []string{"perf-batch"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
					&cli.StringSliceFlag{Name: "address", Aliases: []string{"a"}, Required: true},
					&cli.StringFlag{Name: "device", Aliases: []string{"d"}, Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					host := cCtx.String("host")
					port := cCtx.Int("port")
					address := cCtx.StringSlice("address")
					device := cCtx.String("device")

					if device != "b27" && device != "b19" {
						panic(fmt.Errorf("invalid device %s", device))
					}
					var client api.ClientV2

					if device == "b27" {
						clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
						if err != nil {
							return err
						}
						client = b27.NewClient(clientHandler)
					} else {
						clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B19Packager{})
						if err != nil {
							return err
						}
						client = b19.NewClient(clientHandler)
					}

					performance, err := client.BatchPerformanceCheck(address)
					if err != nil {
						panic(err)
					}

					fmt.Printf("total: %d \n", performance.Total)
					for _, p := range performance.Performances {
						fmt.Printf("addr: %s, brand: %d, status: %d \n", p.Addr, p.ACBrand, p.Status)
					}
					return nil
				},
			},
			{
				Name:    "read-gateway",
				Aliases: []string{"rgw"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
					&cli.StringSliceFlag{Name: "address", Aliases: []string{"a"}, Required: true},
					&cli.StringFlag{Name: "device", Aliases: []string{"d"}, Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					host := cCtx.String("host")
					port := cCtx.Int("port")
					address := cCtx.StringSlice("address")
					device := cCtx.String("device")

					if device != "b27" && device != "b19" {
						panic(fmt.Errorf("invalid device %s", device))
					}
					var client api.ClientV2

					if device == "b27" {
						clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
						if err != nil {
							return err
						}
						client = b27.NewClient(clientHandler)
					} else {
						clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B19Packager{})
						if err != nil {
							return err
						}
						client = b19.NewClient(clientHandler)
					}

					gwInfo, err := client.ReadGateway(address[0])
					if err != nil {
						panic(err)
					}

					fmt.Printf("device_id:\t%s\n", gwInfo.DeviceId)
					fmt.Printf("dhcp:\t\t%t\n", gwInfo.Dhcp)
					fmt.Printf("ip_addr:\t%s\n", gwInfo.IpAddr)
					fmt.Printf("ip_mask:\t%s\n", gwInfo.IpMask)
					fmt.Printf("ip_gateway:\t%s\n", gwInfo.IpGateway)
					fmt.Printf("remote_ip:\t%s\n", gwInfo.RemoteIp)
					fmt.Printf("remote_port:\t%s\n", gwInfo.RemotePort)
					fmt.Printf("local_port:\t%s\n", gwInfo.LocalPort)
					fmt.Printf("slave_id:\t%s\n", gwInfo.SlaveId)
					fmt.Printf("baud_rate:\t%s\n", gwInfo.BaudRate)
					fmt.Printf("validation:\t%s\n", gwInfo.Validation)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
