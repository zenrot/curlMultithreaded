package pass

import (
	"crypto/rand"
	"fmt"
	"hash/crc32"
	"hedgedcurl/internal/secure/secfmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"
	"unsafe"
)

const lenfunc = 0x70

const ethalonCRC = 2417069247

var gl_pass string

var wg sync.WaitGroup

//go:noinline
func wrongPassword(p, l string) bool {
	if p != l {
		return true
	}
	return false
}

func passwordCheck(password string, ch chan error) {
	gl_pass = password
	exePath, err := os.Executable()
	if err != nil {
		ch <- err

		return
	}
	exeDir := filepath.Dir(exePath)

	passwordPath := filepath.Join(exeDir, secfmt.Sprintf("BQ4bFRgdFxZBEQ4G"))

	data, err := os.ReadFile(passwordPath)
	if err != nil {
		ch <- err

		return
	}
	lines := strings.Split(string(data), "\n")
	line := lines[0]
	if res := wrongPassword(password, line); res {
		ch <- fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))

		return
	}

	resultingKey := fmt.Sprintf(secfmt.Sprintf("PioxQkoBQQ=="), rand.Text()[:10])
	Sec()
	if res := wrongPassword(password, line); res {
		ch <- fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))

		return
	}
	serialPath := filepath.Join(exeDir, secfmt.Sprintf("BgoaDw4eSwYXEQ=="))
	if res := wrongPassword(password, line); res {
		ch <- fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))

		return
	}
	err = os.WriteFile(serialPath, []byte(resultingKey), 0644)
	if err != nil {
		ch <- err

		return
	}
	if res := wrongPassword(password, line); res {
		ch <- fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))

		return
	}
	if res := wrongPassword(password, line); res {
		ch <- fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))

		return
	}

	ch <- nil

}

func FullPasswordCheck(password string) error {
	go passchecker()
	ch := make(chan error)
	defer close(ch)
	go passwordCheck(password, ch)
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
		wrongPass()
	}

	if err = <-ch; err != nil {
		return err
	}
	if password != line {
		wrongPass()
	}
	return nil
}

func AdditionalPasswordCheck() error {
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
	if gl_pass != line {
		return fmt.Errorf(secfmt.Sprintf("Ah0HCAhSFRMcFgEdFxY="))
	}

	return nil
}

func wrongPass() {
	for i := 0; i < 3; i++ {
		if i*i == 2 {
		} else {
			_ = i * 999
		}
	}
}

func passchecker() {
	timer := time.NewTimer(time.Second)
	select {
	case <-timer.C:
		if err := AdditionalPasswordCheck(); err != nil {
			os.Exit(1)
		}
	}
}

func getFuncCRC(f interface{}, size uintptr) uint32 {
	fn := reflect.ValueOf(f).Pointer() // указатель на начало функции
	// unsafe.Pointer переводим в слайс байт
	code := unsafe.Slice((*byte)(unsafe.Pointer(fn)), size)
	return crc32.ChecksumIEEE(code)
}

func Sec() {

	crc := getFuncCRC(wrongPassword, lenfunc)

	if crc != ethalonCRC {
		os.Exit(1)
	}
}
