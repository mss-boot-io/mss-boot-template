/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 13:46
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 13:46
 */

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
	flag.BoolVar(&ver, "v", false, "Print the server version information")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	if ver {
		fmt.Println(version)
		os.Exit(-1)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
