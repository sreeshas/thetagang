package main

import (
	"flag"
	"io/ioutil"
	"fmt"
	"encoding/json"
	config "github.com/sreeshas/thetagang/config"
	"github.com/sreeshas/thetagang"
)

var filename = flag.String("config", "config.json", "Location of the config file.")

func main(){
	flag.Parse()
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Print(err)
	}

	var config config.Config
	err = json.Unmarshal(data, &config)
	thetagang.NewCspFinder(config).Execute()

}
