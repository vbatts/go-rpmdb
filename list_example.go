// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vbatts/go-rpmdb"
)

func main() {
	nvrs, err := rpmdb.NVRs()
	if err != nil {
		log.Fatal(err)
	}
	for _, nvr := range nvrs {
		hdr, err := rpmdb.Info(nvr)
		if err != nil {
			log.Fatal(err)
		}
		buf, err := json.MarshalIndent(hdr, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	}
}
