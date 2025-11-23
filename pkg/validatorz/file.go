package validatorz

import (
	"fmt"
	"path/filepath"
)

func (v *Instance) ValidateFileExt(path string, t string) error {
	if ext := filepath.Ext(path); ext != t {
		return fmt.Errorf("require %s file extension", t)
	}
	return nil
}
