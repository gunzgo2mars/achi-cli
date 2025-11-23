package service

import (
	"fmt"
	"os"

	"github.com/gunzgo2mars/achi-cli/internal/core/model"
)

func (s *promptService) setupLayoutProcess(
	rootDir string,
	layoutConfig []model.MicroserviceLayout,
	rootFileConfig []model.MicroserviceRootFile,
) error {

	if err := os.MkdirAll(rootDir, 0755); err != nil {
		return err
	}

	for _, value := range layoutConfig {
		if len(value.SubDir) != 0 {
			for _, subDirValue := range value.SubDir {
				if err := os.MkdirAll(fmt.Sprintf("%s%s%s", rootDir, value.DirName, subDirValue.DirName), 0755); err != nil {
					return err
				}
			}
		}

		if err := os.MkdirAll(fmt.Sprintf("%s%s", rootDir, value.DirName), 0755); err != nil {
			return err
		}
	}

	for _, rootFileValue := range rootFileConfig {
		file, err := os.OpenFile(fmt.Sprintf("%s/%s", rootDir, rootFileValue.FileName), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
		if err != nil {
			return err
		}
		defer file.Close()

		if rootFileValue.FileName == "README.md" {
			rootFileValue.Content = fmt.Sprintf(rootFileValue.Content, rootDir)
		}

		_, err = file.WriteString(rootFileValue.Content)
		if err != nil {
			return err
		}
	}

	return nil
}
