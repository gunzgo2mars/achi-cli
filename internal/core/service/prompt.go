package service

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/gunzgo2mars/achi-cli/internal/core/model"
	driveRepo "github.com/gunzgo2mars/achi-cli/internal/repository/drive"
	"github.com/gunzgo2mars/achi-cli/pkg/menu"
	"github.com/gunzgo2mars/achi-cli/pkg/validatorz"
	"golang.org/x/term"
)

type IPromptService interface {
	DeployProcess()
}

type promptService struct {
	driveRepo driveRepo.IDriveRepository
	validator validatorz.IValidatorz
}

func New(
	driveRepo driveRepo.IDriveRepository,
	validator validatorz.IValidatorz,
) IPromptService {
	return &promptService{
		driveRepo: driveRepo,
		validator: validator,
	}
}

func (s *promptService) DeployProcess() {

	var flagMessage string
	for _, v := range renderFlags() {
		flag.StringVar(&flagMessage, v.Name, v.Description, v.Description)
		flag.Parse()
	}

	if len(os.Args) > 1 {

		// TODO: implement switch case condition

		firstFlag := os.Args[1]

		switch firstFlag {
		case "help", "-h":
			flag.Usage()
			return

		case "init":

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

			var configLayoutData []model.MicroserviceLayout
			var configRootFileData []model.MicroserviceRootFile
			var serviceName string

			fmt.Print("🏗️ service name: ")
			fmt.Scanln(&serviceName)

			serviceNameString := regexp.MustCompile(`\s`)
			isBlankspace := serviceNameString.MatchString(serviceName)

			if isBlankspace {
				log.Fatalf("Error: service name must not cotains any blank space")
			}

			prog := menu.InitProgress("Creating...")

			if err := s.driveRepo.LoadJsonFile("1nWslZeYjwa0oTFWiHl5u_1ISmcpFAxPR", &configLayoutData); err != nil {
				log.Fatalf("Error: %s \n", err.Error())
			}

			if err := s.driveRepo.LoadJsonFile("130__8xGxFfbcRjgwwuiEoj4nJBZzWCIV", &configRootFileData); err != nil {
				log.Fatalf("Error: %s \n", err.Error())
			}

			if choice.MID == 0 {

				s.setupLayoutProcess(
					serviceName,
					configLayoutData,
					configRootFileData,
				)
				prog.Done()
				fmt.Println("\n Done!")
			}

		default:
			flag.Usage()
			return
		}

	}

}

func renderFlags() []model.FlagDetail {

	return []model.FlagDetail{
		{Name: "init", Description: "initializing go project layout and architecture"},
	}

}
