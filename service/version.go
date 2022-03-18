package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const version = "{{.version}}"

var ver bool

func init() {
	flag.String("c", "cfg/{{.service}}.yaml",
		"Read configuration from specified `FILE`, supports JSON/YAML/TOML formats")
	flag.BoolVar(&ver, "v", false, "Print the server version information")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	if ver {
		fmt.Println(version)
		os.Exit(-1)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
