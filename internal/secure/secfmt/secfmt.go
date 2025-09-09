package secfmt

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

var key = []byte("uohforeroevrerbeovob")

func encrypt(input string) string {
	data := []byte(input)
	for i := range data {
		data[i] ^= key[i%len(key)]
	}
	return base64.StdEncoding.EncodeToString(data)
}

func decrypt(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	for i := range decoded {
		decoded[i] ^= key[i%len(key)]
	}
	return string(decoded), nil
}

func processFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		enc := encrypt(line)
		lines = append(lines, fmt.Sprintf("%s | %s", line, enc))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, l := range lines {
		_, _ = out.WriteString(l + "\n")
	}

	return nil
}

func Printf(enc string) (int, error) {
	plain, err := decrypt(enc)
	if err != nil {
		return 0, err
	}
	return fmt.Printf(plain)
}

func Println(enc string) (int, error) {
	plain, err := decrypt(enc)
	if err != nil {
		return 0, err
	}
	return fmt.Println(plain)
}

func Sprintf(enc string) string {
	plain, err := decrypt(enc)
	if err != nil {
		return ""
	}
	return plain
}

func Errorf(enc string, args ...interface{}) (error, error) {
	plain, err := decrypt(enc)
	if err != nil {
		return nil, err
	}
	return fmt.Errorf(plain, args...), nil
}

func Start() {
	err := processFile(Sprintf("EQ4cB0EGHQY="))
	if err != nil {
		Println("pfG57r/KtcO/36bCXw==")
		return
	}
	Println("pcu41r/LtclPtcii1KPitd+m3rLLvuq236LY")

}
