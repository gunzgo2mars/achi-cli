package main

import (
	"context"
	"log"

	"github.com/gunzgo2mars/achi-cli/internal/core/service"
	"github.com/gunzgo2mars/achi-cli/pkg/validatorz"

	driveRepo "github.com/gunzgo2mars/achi-cli/internal/repository/drive"
)

func main() {

	ctx := context.Background()
	validator := validatorz.New()
	driveRepo := driveRepo.New(ctx, "/usr/local/bin/api-drive-config.json")

	log.SetFlags(0)
	promptSvc := service.New(driveRepo, validator)
	promptSvc.DeployProcess()

}
