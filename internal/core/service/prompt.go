package service

import (
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

	if len(os.Args) < 1 {
		log.Print("Error: no flag")
		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "-h" || os.Args[1] == "help" {
		for _, v := range renderFlags() {
			flag.StringVar(&flagMessage, v.Name, v.Description, v.Description)
			flag.Parse()
		}
		flag.Usage()
		return
	}

}

func renderFlags() []model.FlagDetail {

	return []model.FlagDetail{
		{Name: "gitname", Description: "Ask for project github repository."},
	}

}
