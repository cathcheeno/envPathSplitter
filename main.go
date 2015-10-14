package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "eps"
	app.Usage = "list up env paths with split style."

	// setup flags.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "path, p",
			Value: "PATH",
			Usage: "specify the env name.",
		},
		cli.StringFlag{
			Name: "splitter, s",
			Value: ":",
			Usage: "specify the path splitter.",
		},
	}

	// action settings.
	app.Action = func (c *cli.Context) {
		path := c.String("path")
		splitter := c.String("splitter")

		listUpValuesWithSplitter(path, splitter)
	}

	app.Run(os.Args)
}

func listUpValuesWithSplitter(path string, splitter string) {
	fmt.Println("---------------------")
	fmt.Printf("target env: %s\n", path)
	fmt.Printf("splitter: %s\n", splitter)
	fmt.Println("---------------------")
	context := getContextWithPath(path)
	values := strings.Split(context, splitter)
	for _, value := range values {
		fmt.Println(value)
	}
	fmt.Println("---------------------")
}

func getContextWithPath(path string) string {
	var retStr string
	envs := os.Environ()
	for _, val := range envs {
		splits := strings.SplitN(val, "=", 2)
		key := splits[0]
		value := splits[1]
		if key == path {
			retStr = value
			break
		}
	}

	return retStr
}
