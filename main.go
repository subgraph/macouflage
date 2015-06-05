package main

import (
	"os"
	"fmt"
	"log"
	"github.com/codegangsta/cli"

)

func main() {
	app := cli.NewApp()
	app.Name = "macouflage"
	app.Usage ="macouflage is a MAC address anonymization tool"
	app.Version = "0.1"
	app.Author = "David McKinney"
	app.Email = "mckinney@subgraph.com"
	// BUG: Help template does not show subcommands by default, supply own template
	app.Commands = []cli.Command {
		{
			Name: "show",
			Aliases: []string{"s"},
			Usage: "Print the MAC address and exit",
			Action: show,
		},
		{
			Name: "list",
			Usage: "Print known vendors",
			Action: list,
			Subcommands: []cli.Command{
				{
					Name: "popular",
					Usage: "Print known popular vendors",
					Action: listPopular,
				},
			},
		},
		{
			Name: "search",
			Usage: "Search vendor names",
			Action: search,
		},
	}
	app.Run(os.Args)
}

func show(c *cli.Context)  {
	result, err := getMacInfo(c.Args().First())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func list(c *cli.Context) {
	results, err := listVendors("", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func listPopular(c *cli.Context) {
	results, err := listVendors("", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func search(c *cli.Context) {
	results, err := listVendors(c.Args().First(), false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}
