package service

import (
	"flag"
	"log"
	"os"

	"github.com/gunzgo2mars/achi-cli/internal/core/model"
	"github.com/gunzgo2mars/achi-cli/pkg/menu"
	"golang.org/x/term"
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

		// if os.Args[1] == "-h" || os.Args[1] == "help" {
		// 	flag.Usage()
		// 	return
		// }

		// if os.Args[1] == "init" {
		// 	log.Println("Application initialized.")
		// 	log.Print("Please type a project name and press Enter:")
		// 	reader := bufio.NewReader(os.Stdin)
		// 	projectName, err := reader.ReadString('\n')
		// 	if err != nil {
		// 		log.Fatalf("Error reading input: %v", err)
		// 	}

		// 	log.Printf("Project name: %s \n", projectName)
		// 	os.Exit(0)

		// }

		// TODO: implement switch case condition

		firstFlag := os.Args[1]

		switch firstFlag {
		case "help", "-h":
			flag.Usage()
			return

		case "init":

			log.Println("⚔ Achilles - go project layout generator ⚔️")
			fd := int(os.Stdin.Fd())

			if !term.IsTerminal(fd) {
				log.Fatal("Application must be run in a proper terminal to use arrow-key selection.")
			}

			previousTermState, err := term.MakeRaw(fd)
			if err != nil {
				log.Fatalf("Failed to put terminal into raw mode: %v", err)
			}

			defer term.Restore(fd, previousTermState)

			choice, err := menu.HandleSelection(previousTermState, fd)
			if err != nil {
				log.Printf("Error: %s \n", err.Error())
			}

			log.Printf("Choice: %s with MID:%d \n", choice.MenuName, choice.MID)

		default:
			flag.Usage()
			return
		}

	}

}

func renderFlags() []model.FlagDetail {

	return []model.FlagDetail{
		{Name: "gitname", Description: "Ask for project github repository."},
	}

}
