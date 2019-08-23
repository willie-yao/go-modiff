package main

import (
	"fmt"
	"os"

	"github.com/saschagrunert/ccli"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/saschagrunert/go-modiff/internal/modiff"
)

func main() {
	// Init the logging facade
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	logrus.SetLevel(logrus.DebugLevel)

	// Enable to modules
	os.Setenv("GO111MODULE", "on")

	app := ccli.NewApp()
	app.Name = "go-modiff"
	app.Version = "0.3.0"
	app.Author = "Sascha Grunert"
	app.Email = "mail@saschagrunert.de"
	app.Usage = "Command line tool for diffing go module " +
		"dependency changes between versions"
	app.UsageText = app.Usage
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  modiff.RepositoryArg + ", r",
			Usage: "repository to be used, like: github.com/owner/repo",
		},
		cli.StringFlag{
			Name:  modiff.FromArg + ", f",
			Value: "master",
			Usage: "the start of the comparison, any valid git rev",
		},
		cli.StringFlag{
			Name:  modiff.ToArg + ", t",
			Value: "master",
			Usage: "the end of the comparison, any valid git rev",
		},
	}
	app.Commands = []cli.Command{{
		Name:    "docs",
		Aliases: []string{"d"},
		Action:  docs,
		Usage: "generate the markdown or man page documentation " +
			"and print it to stdout",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "markdown",
				Usage: "print the markdown version",
			},
			cli.BoolFlag{
				Name:  "man",
				Usage: "print the man version",
			},
		},
	}}
	app.Action = func(c *cli.Context) error {
		res, err := modiff.Run(c)
		if err != nil {
			return err
		}
		logrus.Info("Done, the result will be printed to `stdout`")
		fmt.Print(res)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func docs(c *cli.Context) (err error) {
	res := ""
	if c.Bool("markdown") {
		res, err = c.App.ToMarkdown()
	} else if c.Bool("man") {
		res, err = c.App.ToMan()
	}
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", res)
	return nil
}
