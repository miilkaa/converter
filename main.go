package main

import (
	"flag"
	"fmt"

	"github.com/miilkaa/converter/internal/converter"
)

func main() {
	mode := flag.String("mode", "", "Conversion mode: json2env or env2json")
	jsonFile := flag.String("json", "config.json", "Path to the JSON file")
	envName := flag.String("env", ".env", "Name of the output env file")

	flag.Parse()

	fmt.Println("Mode: ", *mode)
	fmt.Println("JSON File: ", *jsonFile)
	fmt.Println("Env File: ", *envName)

	switch *mode {
	case "env2json":
		err := converter.ConvertEnvToJSON(*envName, *jsonFile)
		if err != nil {
			panic(err)
		}
	case "json2env":
		err := converter.ConvertJSONtoEnv(*jsonFile, *envName)
		if err != nil {
			panic(err)
		}
	default:
		panic("Invalid mode. Use 'json2env' or 'env2json'")
	}

}
