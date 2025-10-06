package main

import (
	"log"

	"github.com/gunzgo2mars/achi-cli/internal/core/service"
)

func main() {

	log.SetFlags(0)

	promptSvc := service.New("test")
	promptSvc.DeployProcess()

}
