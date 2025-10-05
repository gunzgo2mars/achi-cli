package main

import "github.com/gunzgo2mars/achi-cli/internal/core/service"

func main() {

	promptSvc := service.New("test")
	promptSvc.DeployProcess()

}
