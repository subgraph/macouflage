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
	app.Flags = []cli.Flag {
			cli.StringFlag{
				Name: "i, interface",
				Usage: "Target device (required)",
			},
			cli.BoolFlag{
				Name: "b, bia",
				Usage: "Pretend to be a burned-in-address",
			},
	}
	// BUG: Help template does not show subcommands by default, supply own template
	app.Commands = []cli.Command {
		{
			Name: "show",
			Usage: "Print the MAC address and exit",
			Action: show,
		},
		{
			Name: "ending",
			Usage: "Don't change the vendor bytes (generate last three bytes: XX:XX:XX:??:??:??)",
			Action: ending,
		},
		{
			Name: "another",
			Usage: "Set random vendor MAC of the same kind",
			Action: another,
		},
		{
			Name: "any",
			Usage: "Set random vendor MAC of any kind",
			Action: any,
		},
		{
			Name: "permanent",
			Usage: "Reset to original, permanent hardware MAC",
			Action: permanent,
		},
		{
			Name: "random",
			Usage: "Set fully random MAC",
			Action: random,
		},
		{
			Name: "popular",
			Usage: "Set a MAC from the popular vendors list",
			Action: popular,
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
		{
			Name: "mac",
			Usage: "Set the MAC XX:XX:XX:XX:XX:XX",
			Action: mac,
		},
	}
	app.Run(os.Args)
}

func show(c *cli.Context)  {
	iface := c.GlobalString("i")
	if iface != "" {
		result, err := getCurrentMacInfo(iface)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
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

func ending(c *cli.Context) {
	iface := c.GlobalString("i")
	if iface != "" {
		err := spoofMacEnding(iface)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}

func another(c *cli.Context) {
	iface := c.GlobalString("i")
	if iface != "" {
		err := spoofMacAnother(iface)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}

func any(c *cli.Context) {
	iface := c.GlobalString("i")
	if iface != "" {
		err := spoofMacAny(iface)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}

func permanent(c *cli.Context) {
	iface := c.GlobalString("i")
	if iface != "" {
		err := revertMac(iface)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}

func random(c *cli.Context) {
	iface := c.GlobalString("i")
	if iface != "" {
		err := spoofMacRandom(iface, c.GlobalBool("b"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}

func popular(c *cli.Context) {
	iface := c.GlobalString("i")
	if iface != "" {
		err := spoofMacPopular(iface)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}

func mac(c *cli.Context) {
	if c.Args().First() == "" {
		log.Fatal("No MAC address argument specified")
	}
	iface := c.GlobalString("i")
	if iface != "" && c.Args().First() != "" {
		err := spoofMac(iface, c.Args().First())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("No target device provided via -i, --interface argument")
	}
}
