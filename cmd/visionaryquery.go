package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"visionaryquery/internal/config"
	"visionaryquery/internal/datastore"

	"gopkg.in/yaml.v2"
)

func Run() int {

	// GET SYSTEM PATH
	path, err := os.Getwd()

	if err != nil {
		fmt.Println("ERROR - Unable to determine system path")
		return 1
	}

	// GET CONFIG FILE
	fileName := fmt.Sprintf("%v/config.yml", path)

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

	// Parse flag
	timecode := flag.Bool("timecode", false, "Result will include timecode")
	actor := flag.Bool("actor", false, "Result will include actor")
	character := flag.Bool("character", false, "Result will include character")
	flag.Parse()

	if !*timecode && !*actor && !*character {
		fmt.Println("ERROR - At least 1 result option must be present")
		flag.Usage()
		return 1
	}

	// Connect to DB
	connectionString := cfg.ConnectionURI()

	store, err := datastore.New(connectionString)
	result, err := store.Query(cfg.Dataset, cfg.Studio)

	if err != nil {
		fmt.Printf("ERROR - SQL query failed\n")
		fmt.Println(err)
		return 1
	}

	// Success
	var returnString []string

	if *timecode {
		returnString = append(returnString, result.TimeCode)
	}

	if *actor {
		returnString = append(returnString, result.Actor)
	}

	if *character {
		returnString = append(returnString, result.Character)
	}

	final := strings.Join(returnString[:], ",")
	fmt.Println(final)

	return 0
}

func main() {
	retCode := Run()
	os.Exit(retCode)
}
