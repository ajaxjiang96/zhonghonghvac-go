package b19cmd

import (
	"context"
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b19"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	ch "github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/urfave/cli/v3"
)

var B19Cmd = &cli.Command{
	Name:    "b19",
	Aliases: []string{"vrf"},
	Usage:   "control the zhonghong HVAC vrf controller device",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
		&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
		&cli.StringFlag{Name: "address", Aliases: []string{"a"}, Required: true},
	},
	Commands: []*cli.Command{
		StatusCommand,
	},
}

var StatusCommand = &cli.Command{ // status
	Name: "status",
	Action: func(_ context.Context, cmd *cli.Command) error {
		host := cmd.String("host")
		port := cmd.Int("port")
		address := cmd.String("address")
		clientHandler, err := clienthandler.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B19Packager{})
		if err != nil {
			return err
		}
		client := b19.NewClient(clientHandler)
		res, err := client.StatusCheck(address)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var PerformanceCommand = &cli.Command{
	Name:    "performance",
	Aliases: []string{"perf"},
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "device", Aliases: []string{"d"}, Required: true},
	},
	Action: func(_ context.Context, cmd *cli.Command) error {
		host := cmd.String("host")
		port := cmd.Int("port")
		address := cmd.String("address")
		device := cmd.String("device")

		if device != "b27" && device != "b19" {
			panic(fmt.Errorf("invalid device %s", device))
		}
		var client api.ClientV2

		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &clienthandler.B27Packager{})
		if err != nil {
			return err
		}
		client = b19.NewClient(clientHandler)

		performance, err := client.PerformanceCheck(address)
		if err != nil {
			panic(err)
		}

		fmt.Printf("addr: %s, brand: %d, status: %d \n", performance.Addr, performance.ACBrand, performance.Status)
		return nil
	},
}
