package b27cmd

import (
	"context"
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b27"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
	"github.com/urfave/cli/v3"
)

var B27Cmd = &cli.Command{
	Name:    "b27",
	Aliases: []string{"superman"},
	Usage:   "control the zhonghong HVAC device",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
		&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
		&cli.StringFlag{Name: "address", Aliases: []string{"a"}, Required: true},
	},
	Commands: []*cli.Command{
		PerformanceCommand,
		StatusCommand,
		ControlCommand,
	},
}

var PerformanceCommand = &cli.Command{ // performance
	Name: "performance",
	Action: func(_ context.Context, cmd *cli.Command) error {
		host := cmd.String("host")
		port := cmd.Int("port")
		address := cmd.String("address")

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
}

var StatusCommand = &cli.Command{ // status
	Name: "status",
	Action: func(_ context.Context, cmd *cli.Command) error {
		host := cmd.String("host")
		port := cmd.Int("port")
		address := cmd.String("address")

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
}

var ControlCommand = &cli.Command{ // control
	Name: "control",
	Flags: []cli.Flag{
		&cli.IntFlag{Name: "temperature", Aliases: []string{"temp", "t"}},
		&cli.BoolFlag{Name: "onoff", Aliases: []string{"o"}, Usage: "1: on, 0: off"},
		&cli.IntFlag{Name: "fanspeed", Aliases: []string{"f"}, Value: 0, Usage: "0: auto, 1: low, 2: medium, 3: high"},
		&cli.IntFlag{Name: "mode", Aliases: []string{"m"}, Value: 2, Usage: "1: heat, 2: cool, 4: vent, 8: dehumidify"},
		&cli.IntFlag{Name: "direction", Aliases: []string{"d"}, Value: 0, Usage: "0: no direction, 1-7: direction, 0xff: auto"},
	},
	Action: func(_ context.Context, cmd *cli.Command) error {
		host := cmd.String("host")
		port := cmd.Int("port")
		address := cmd.String("address")
		temperature := cmd.Int("temperature")
		onOff := cmd.Bool("onoff")
		fanSpeed := cmd.Int("fanspeed")
		mode := cmd.Int("mode")
		direction := cmd.Int("direction")
		verbose := cmd.Bool("verbose")
		clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
		if err != nil {
			return err
		}
		client := b27.NewClient(clientHandler)
		command := protocol.ACControlRequest{}

		if cmd.IsSet("onoff") {
			if verbose {
				fmt.Printf("setting onoff: %v\n", onOff)
			}
			command.On = &onOff
		}
		if cmd.IsSet("temperature") {
			if verbose {
				fmt.Printf("setting temperature: %d\n", temperature)
			}
			temp := int(temperature)
			command.Temp = &temp
		}
		if cmd.IsSet("fanspeed") {
			if verbose {
				fmt.Printf("setting fanspeed: %d\n", fanSpeed)
			}
			fanSpeedEnum := protocol.FanSpeed(byte(fanSpeed))
			command.FanSpeed = &fanSpeedEnum
		}
		if cmd.IsSet("mode") {
			if verbose {
				fmt.Printf("setting mode: %d\n", mode)
			}
			modeEnum := protocol.ACMode(byte(mode))
			command.Mode = &modeEnum
		}
		if cmd.IsSet("direction") {
			if verbose {
				fmt.Printf("setting direction: %d\n", direction)
			}
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
}
