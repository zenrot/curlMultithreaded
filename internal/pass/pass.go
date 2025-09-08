package pass

import (
	"crypto/rand"
	"fmt"
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

	passwordPath := filepath.Join(exeDir, "password.txt")

	data, err := os.ReadFile(passwordPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	line := lines[0]
	if password != line {
		return fmt.Errorf("wrong password")
	}

	resultingKey := fmt.Sprintf("KEY$%s$", rand.Text()[:10])
	serialPath := filepath.Join(exeDir, "serial.txt")
	err = os.WriteFile(serialPath, []byte(resultingKey), 0644)
	if err != nil {
		return err
	}
	return nil
}
