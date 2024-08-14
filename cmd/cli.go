package main

import (
	"context"
	"os"

	b19cmd "github.com/Yangsta911/zhonghonghvac-go/cmd/b19"
	b27cmd "github.com/Yangsta911/zhonghonghvac-go/cmd/b27"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}},
		},
		Commands: []*cli.Command{
			b27cmd.B27Cmd,
			b19cmd.B19Cmd,
			{
				Name:    "performance-batch",
				Aliases: []string{"perf-batch"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
					&cli.StringSliceFlag{Name: "address", Aliases: []string{"a"}, Required: true},
					&cli.StringFlag{Name: "device", Aliases: []string{"d"}, Required: true},
				},
				Action: PerformanceBatch,
			},
			{
				Name:    "performance",
				Aliases: []string{"perf"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "192.168.1.220"},
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 4196},
					&cli.StringSliceFlag{Name: "address", Aliases: []string{"a"}, Required: true},
					&cli.StringFlag{Name: "device", Aliases: []string{"d"}, Required: true},
				},
				Action: Performance,
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
				Action: ReadGateway,
			},
		}}
	cmd.Run(context.Background(), os.Args)
}
