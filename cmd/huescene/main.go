package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/amimof/huego"

	"github.com/bruce/huescene/internal/huescene"
)

var configpath string
var docmdlistlights bool
var userkey string

func main() {
	flag.StringVar(&configpath, "config", "", "Path to the huescene YAML configuration")
	flag.StringVar(&configpath, "c", "", "Path to the huescene YAML configuration")
	flag.StringVar(&userkey, "key", "", "Hue Bridge user key")
	flag.BoolVar(&docmdlistlights, "list-lights", false, "List the names of the available lights")
	flag.Parse()

	if configpath == "" {
		log.Fatal("No --config path given")
	}

	bridge, err := huego.Discover()
	if err != nil {
		log.Panic(err)
	}

	cfg, err := huescene.ReadConfig(configpath)
	if err != nil {
		log.Panic(err)
	}

	key := findKey(*cfg)
	if key == "" {
		cmdCreateUser(*bridge, *cfg)
	} else {
		authd := huego.New(bridge.Host, key)
		if docmdlistlights {
			cmdListLights(*authd, *cfg)
		} else {
			cmdSetScene(*authd, *cfg)
		}
	}
}

func cmdListLights(bridge huego.Bridge, cfg huescene.Config) {
	ls, err := bridge.GetLights()
	if err != nil {
		log.Panic(err)
	}

	sort.Slice(ls, func(i, j int) bool { return ls[i].Name < ls[j].Name })

	for _, l := range ls {
		fmt.Println(l.Name)
	}
}

func cmdCreateUser(bridge huego.Bridge, cfg huescene.Config) {
	key, err := bridge.CreateUser(cfg.Username)
	if err != nil {
		log.Println("Did you forget to press to Hue Bridge Link button?")
		log.Panic(err)
	}

	log.Printf("Hue Bridge user \"%s\" created.\nPlease add the following to your configuration:\n\nkey: %s\n", cfg.Username, key)
	os.Exit(0)
}

func cmdSetScene(bridge huego.Bridge, cfg huescene.Config) {
	err := huescene.SetScene(cfg, bridge, flag.Args()[0])
	if err != nil {
		log.Println("An error occurred.")
		log.Fatal(err)
	}
}

func findKey(cfg huescene.Config) string {
	if cfg.Key != "" {
		return cfg.Key
	} else if userkey != "" {
		return userkey
	}
	return ""
}
