package drive

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type IDriveRepository interface {
	LoadJsonFile(fileId string, config any) error
}

type driveRepository struct {
	driveService *drive.Service
}

var _ IDriveRepository = (*driveRepository)(nil)

func New(ctx context.Context, credentialFile string) IDriveRepository {

	driveSvc, err := drive.NewService(ctx, option.WithCredentialsFile(credentialFile))
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}

	return &driveRepository{
		driveService: driveSvc,
	}
}

func (r *driveRepository) LoadJsonFile(fileId string, config any) error {

	resp, err := r.driveService.Files.Get(fileId).Download()
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}

	return json.Unmarshal(data, config)
}
