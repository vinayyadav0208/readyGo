package main

import (
	"fmt"
	"readyGo/cmd"
)

func main() {

	fmt.Println("Hello Muruga")

	cmd.Execute()

	/*log.Println("Loading all templates into in-memory")

	configFile := "configs/config.yaml"

	tm, err := template.New("templates")

	if err != nil {
		log.Fatal("error occured loading templates..---->", err)
	}

	log.Println("Generating files and dependencies based on config file.")

	templateConfig := "configs/template_config.json"

	tc, err := configure.New(&templateConfig)

	fmt.Println(tc, err)

	tg, err := generate.New(&configFile, tm, tc)

	if err != nil {
		log.Fatal("seems , things went wrong.. -->", err)
	}

	err = tg.GenerateAll("http_mongo")
	if err != nil {
		log.Println("seems , things went wrong.Rolling back all generated files -->", err)
		err = tg.RmDir()
		if err != nil {
			log.Println("Unable to remove files. Please remove created directory manually", err)
		}
	}*/

}
