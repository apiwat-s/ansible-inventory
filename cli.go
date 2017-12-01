package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

func appendUniq(strs []string, item string) []string {
	if len(strs) == 0 {
		strs = append(strs, item)
		return strs
	}
	sort.Strings(strs)
	i := sort.SearchStrings(strs, item)
	if i == len(strs) || (i < len(strs) && strs[i] != item) {
		strs = append(strs, item)
	}
	return strs
}

func gatherResources(url string) []inventory {
	urlGet := apiEndpoint + "/" + url

	res, err := http.Get(urlGet)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	var result []inventory
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func buildInventory(ivtr []inventory) inventoryALL {
	outputGroups := make(map[string]inventoryALL)

	resource := resource{Hosts: make([]string, 0), Children: make([]string, 0), Vars: make(map[string]interface{})}
	tasks := make(map[string][]string)

	for _, res := range ivtr {
		// place in list of all to resources
		resource.Hosts = appendUniq(resource.Hosts, res.IPAddrPublic)

		if len(res.Tasks) != 0 {
			for _, task := range res.Tasks {
				// place in list of resource tasks
				tasks[task] = appendUniq(tasks[task], res.IPAddrPublic)
				outputGroups[task] = tasks[task]
			}
		}

	}

	outputGroups["all"] = resource
	// for k, v := range types {
	// 	if old, exists := outputGroups[k]; exists {
	// 		fmt.Fprintf(os.Stderr, "types overwriting already existing output key %s, old: %v, new: %v", k, old, v)
	// 	}
	// 	outputGroups[k] = v
	// }
	return outputGroups
}

func cliList(stdout io.Writer, stderr io.Writer, hostList inventoryALL) int {
	return output(stdout, stderr, hostList)
}

func cliHost(stdout io.Writer, stderr io.Writer, hostList inventoryALL) int {
	return output(stdout, stderr, hostList)
}

func output(stdout io.Writer, stderr io.Writer, whatever interface{}) int {
	b, err := json.Marshal(whatever)
	if err != nil {
		fmt.Fprintf(stderr, "Error encoding JSON: %s\n", err)
		return 1
	}

	_, err = stdout.Write(b)
	if err != nil {
		fmt.Fprintf(stderr, "Error writing JSON: %s\n", err)
		return 1
	}

	return 0
}
