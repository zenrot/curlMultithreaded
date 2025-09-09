package pass

import (
	"crypto/rand"
	"fmt"
	"hedgedcurl/internal/secure/secfmt"
	"os"
	"path/filepath"
	"strings"
)

func PasswordCheck(password string) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exeDir := filepath.Dir(exePath)

	passwordPath := filepath.Join(exeDir, secfmt.Sprintf("BQ4bFRgdFxZBEQ4G"))

	data, err := os.ReadFile(passwordPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	line := lines[0]
	if password != line {
		return fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))
	}

	resultingKey := fmt.Sprintf(secfmt.Sprintf("PioxQkoBQQ=="), rand.Text()[:10])
	serialPath := filepath.Join(exeDir, secfmt.Sprintf("BgoaDw4eSwYXEQ=="))
	err = os.WriteFile(serialPath, []byte(resultingKey), 0644)
	if err != nil {
		return err
	}
	return nil
}
