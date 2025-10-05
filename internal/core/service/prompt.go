package service

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/gunzgo2mars/achi-cli/internal/core/model"
)

type IPromptService interface {
	DeployProcess()
}

type promptService struct {
	message string
}

func New(message string) IPromptService {
	return &promptService{
		message: message,
	}
}

func (s *promptService) DeployProcess() {

	var flagMessage string
	for _, v := range renderFlags() {
		flag.StringVar(&flagMessage, v.Name, v.Description, v.Description)
		flag.Parse()
	}

	if len(os.Args) > 1 {

		if os.Args[1] == "-h" || os.Args[1] == "help" {
			flag.Usage()
			return
		}

		if os.Args[1] == "init" {
			log.Println("Application initialized.")
			log.Print("Please type a project name and press Enter:")
			reader := bufio.NewReader(os.Stdin)
			projectName, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading input: %v", err)
			}

			log.Printf("Project name: %s \n", projectName)
			os.Exit(0)

		}

	}

}

func renderFlags() []model.FlagDetail {

	return []model.FlagDetail{
		{Name: "gitname", Description: "Ask for project github repository."},
	}

}
