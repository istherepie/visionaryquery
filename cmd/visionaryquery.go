package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"visionaryquery/internal/config"
	"visionaryquery/internal/datastore"

	"gopkg.in/yaml.v2"
)

func Run() int {

	// GET SYSTEM PATH
	exe, err := os.Executable()

	if err != nil {
		fmt.Println("ERROR - Unable to determine system path")
		return 1
	}

	configPath := path.Dir(exe)

	defaultConfigPath := fmt.Sprintf("%v/config.yml", configPath)

	// Parse flag
	configFile := flag.String("config", defaultConfigPath, "Config file location (unix path)")
	timecode := flag.Bool("timecode", false, "Result will include timecode")
	actor := flag.Bool("actor", false, "Result will include actor")
	character := flag.Bool("character", false, "Result will include character")
	flag.Parse()

	if !*timecode && !*actor && !*character {
		flag.Usage()
		return 0
	}

	// GET CONFIG FILE
	// TODO: Perhaps there should be a list of default locations
	// Example: /usr/local/etc/namespace/config.yml etc.

	var fileName string

	if config.FileExists(*configFile) {
		fileName = *configFile
	} else {
		fileName = defaultConfigPath
	}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("ERROR - Unable to load config file (%s)\n", err)
		return 1
	}

	// Parse config
	var cfg config.Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		fmt.Printf("ERROR - parsing configuration file (%s)\n", err)
		return 1
	}

	// Connect to DB
	connectionString := cfg.ConnectionURI()

	store, err := datastore.New(connectionString)

	// TODO: The constructor should set the table and schema
	store.Schema = cfg.Schema
	store.Table = cfg.Table

	result, err := store.Query(cfg.Dataset, cfg.Studio)

	if err != nil {
		fmt.Printf("ERROR - SQL query failed\n")
		fmt.Println(err)
		return 1
	}

	// Success
	var returnString []string

	if *actor {
		returnString = append(returnString, result.Actor)
	}

	if *character {
		returnString = append(returnString, result.Character)
	}

	if *timecode {
		returnString = append(returnString, result.TimeCode)
	}

	final := strings.Join(returnString[:], ",")
	fmt.Println(final)

	return 0
}

func main() {
	retCode := Run()
	os.Exit(retCode)
}
