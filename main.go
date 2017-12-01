package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var version = flag.Bool("version", false, "print version information")
var list = flag.Bool("list", false, "list mode")
var test = flag.Bool("test", false, "test mode")
var host = flag.String("host", "", "host mode")
var apiEndpoint = "http://5a1bc601c3630f0012b241f4.mockapi.io/api/v1"

type inventoryALL interface{}

type resource struct {
	Hosts    []string               `json:"hosts"`
	Children []string               `json:"children"`
	Vars     map[string]interface{} `json:"vars"`
}

type inventory struct {
	ID             string   `json:"id"`
	CreatedAt      int      `json:"createdAt"`
	Hostname       string   `json:"hostname"`
	IPAddrPublic   string   `json:"ipaddrPublic"`
	IPAddrPrivate  string   `json:"ipaddrPrivate"`
	Tasks          []string `json:"tasks"`
	PrivateNetwork bool     `json:"private_networking"`
}

type meta struct {
	HostVars map[string]string `json:"hostvars"`
}

func main() {
	flag.Parse()

	if *version == true {
		fmt.Printf("%s version %s\n", os.Args[0], versionInfo())
		return
	}

	if !*list && *host == "" && !*test {
		fmt.Fprint(os.Stderr, "Either --host or --list must be specified")
		os.Exit(1)
	}

	if *list {
		res := gatherResources("inventory")
		all := buildInventory(res)
		os.Exit(cliList(os.Stdout, os.Stderr, &all))
	}

	if *test {
		byt := []byte(`{"hosts": [], "vars": {}}`)
		// byt.hosts = "sdaf"

		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)
	}
}
