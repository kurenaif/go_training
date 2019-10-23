package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type DepList struct {
	Dir  string
	Name string
	Deps []string
}

func splitJson(data []byte) [][]byte {
	// split json (go list -json hoge fuga のjsonがjsonではないため)
	var jsons [][]byte
	depth := 0
	prev := 0
	flag := false
	for i := 0; i < len(data); i++ {
		if data[i] == byte('{') {
			flag = true
			depth++
		}
		if data[i] == byte('}') {
			depth--
		}
		if depth == 0 && flag {
			// fmt.Println("-----")
			// fmt.Println(string(data[prev : i+1]))
			jsons = append(jsons, data[prev:i+1])
			prev = i + 1
			flag = false
		}
	}
	return jsons
}

func getDeps(packages []string) []DepList {
	args := append([]string{"list", "-json"}, packages...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	jsons := splitJson(out)

	var deps []DepList
	for _, jsonData := range jsons {
		var list DepList
		if err := json.Unmarshal(jsonData, &list); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(list.Dir)
		// fmt.Println(list.Deps)
		deps = append(deps, list)
	}
	return deps
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: main package...")
		os.Exit(1)
	}
	packages := getDeps(os.Args[1:])
	packagesAll := getDeps([]string{"..."})

	workSpacePackageSet := make(map[string]string)
	for _, pack := range packagesAll {
		workSpacePackageSet[pack.Name] = pack.Dir
	}

	for _, pack := range packages {
		for _, dep := range pack.Deps {
			if path, ok := workSpacePackageSet[dep]; ok {
				fmt.Println(pack.Dir, "=>", path)
			}
		}
	}
}
