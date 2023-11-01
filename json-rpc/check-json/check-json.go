package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/pretty"

	openrpc_document "github.com/ParspooyeshFanavar/meta-schema"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: " + os.Args[0] + " .../ibsng-docs/json-rpc")
	}
	dirPath := os.Args[1]
	branches := []string{"E", "D", "C"}
	for _, branch := range branches {
		branchDir := filepath.Join(dirPath, branch)
		entries, err := os.ReadDir(branchDir)
		if err != nil {
			panic(err)
		}
		for _, entry := range entries {
			fpath := filepath.Join(branchDir, entry.Name())
			if strings.HasSuffix(fpath, ".std.json") {
				continue
			}
			ext := filepath.Ext(fpath)
			if ext != ".json" {
				continue
			}
			fpath_nox := fpath[:len(fpath)-5]
			jsonB, err := os.ReadFile(fpath)
			if err != nil {
				panic(err)
			}
			doc := &openrpc_document.OpenrpcDocument{}
			err = json.Unmarshal(jsonB, doc)
			if err != nil {
				panic(err)
			}
			jsonB_2, err := json.Marshal(doc)
			if err != nil {
				panic(err)
			}
			jsonB_3 := pretty.PrettyOptions(jsonB_2, &pretty.Options{
				Width:    80,
				Prefix:   "",
				Indent:   "  ",
				SortKeys: false,
			})
			stdJsonPath := fpath_nox + ".std.json"
			err = os.WriteFile(stdJsonPath, jsonB_3, 0o644)
			if err != nil {
				panic(err)
			}
			fmt.Println(stdJsonPath)
		}
	}
}
