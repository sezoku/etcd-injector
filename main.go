package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	appName    = "etcd-replacer"
	appVersion = "dummy"
)

type config struct {
	AppName       string
	AppVersion    string
	SrcEndpoints  []string
	DstEndpoints  []string
	SrcDirectory  string
	DstDirectory  string
	RulesFilepath string
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Usage = "recursive copy & replace json value of etcd"
	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "src-endpoints",
			Usage:    "source endpoints of etcd",
			Required: true,
			EnvVars:  []string{"ETCD_SRC_ENDPOINTS"},
		},
		&cli.StringSliceFlag{
			Name:     "dst-endpoints",
			Usage:    "destination directory of etcd",
			Required: true,
			EnvVars:  []string{"ETCD_DST_ENDPOINTS"},
		},
		&cli.StringFlag{
			Name:     "src-directory",
			Aliases:  []string{"s"},
			Usage:    "source directory of etcd",
			Required: true,
			EnvVars:  []string{"ETCD_SRC_DIRECTORY"},
		},
		&cli.StringFlag{
			Name:     "dst-directory",
			Aliases:  []string{"d"},
			Usage:    "destination directory of etcd",
			Required: true,
			EnvVars:  []string{"ETCD_DST_DIRECTORY"},
		},
		&cli.StringFlag{
			Name:     "rules-filepath",
			Aliases:  []string{"f"},
			Usage:    "path of file written replace rules",
			Required: true,
			EnvVars:  []string{"RULES_FILEPATH"},
		},
		// TODO
		// &cli.StringFlag{
		//	Name:     "src-username",
		// },
		// &cli.StringFlag{
		//	Name:     "src-password",
		// },
		// &cli.StringFlag{
		//	Name:     "src-cafile",
		// },
		// &cli.StringFlag{
		//	Name:     "src-certfile",
		// },
		// &cli.StringFlag{
		//	Name:     "src-keyfile",
		// },
		// &cli.StringFlag{
		//	Name:     "dst-username",
		// },
		// &cli.StringFlag{
		//	Name:     "dst-password",
		// },
		// &cli.StringFlag{
		//	Name:     "dst-cafile",
		// },
		// &cli.StringFlag{
		//	Name:     "dst-certfile",
		// },
		// &cli.StringFlag{
		//	Name:     "dst-keyfile",
		// },
	}
	app.Action = func(c *cli.Context) error {
		return Run(config{
			AppName:       appName,
			AppVersion:    appVersion,
			SrcEndpoints:  c.StringSlice("src-endpoints"),
			DstEndpoints:  c.StringSlice("dst-endpoints"),
			SrcDirectory:  c.String("src-directory"),
			DstDirectory:  c.String("dst-directory"),
			RulesFilepath: c.String("rules-filepath"),
		})
	}
	app.EnableBashCompletion = true
	return app
}

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}