package main

import (
	"context"
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b19"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b27"
	ch "github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/urfave/cli/v3"
)

func PerformanceBatch(_ context.Context, cmd *cli.Command) error {
	host := cmd.String("host")
	port := cmd.Int("port")
	address := cmd.StringSlice("address")
	device := cmd.String("device")

	if device != "b27" && device != "b19" {
		panic(fmt.Errorf("invalid device %s", device))
	}
	var client api.ClientV2

	if device == "b27" {
		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &ch.B27Packager{})
		if err != nil {
			return err
		}
		client = b27.NewClient(clientHandler)
	} else {
		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &ch.B19Packager{})
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
}

func ReadGateway(_ context.Context, cmd *cli.Command) error {
	host := cmd.String("host")
	port := cmd.Int("port")
	address := cmd.StringSlice("address")
	device := cmd.String("device")

	if device != "b27" && device != "b19" {
		panic(fmt.Errorf("invalid device %s", device))
	}
	var client api.ClientV2

	if device == "b27" {
		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &ch.B27Packager{})
		if err != nil {
			return err
		}
		client = b27.NewClient(clientHandler)
	} else {
		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &ch.B19Packager{})
		if err != nil {
			return err
		}
		client = b19.NewClient(clientHandler)
	}

	gwInfo, err := client.ReadGateway(address[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(gwInfo)

	return nil
}

func Performance(_ context.Context, cmd *cli.Command) error {
	host := cmd.String("host")
	port := cmd.Int("port")
	address := cmd.String("address")
	device := cmd.String("device")

	if device != "b27" && device != "b19" {
		panic(fmt.Errorf("invalid device %s", device))
	}
	var client api.ClientV2

	if device == "b27" {
		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &ch.B27Packager{})
		if err != nil {
			return err
		}
		client = b27.NewClient(clientHandler)
	} else {
		clientHandler, err := ch.NewTCPClientHandler(fmt.Sprintf("%s:%d", host, port), &ch.B19Packager{})
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
}
