package source

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Print logo
func (Xc *Checker) Logo() {
	logo := `
	___` + z + `__  __` + r + `____` + z + `____________` + r + `___` + z + `_____` + r + `
	__  ` + z + `/ / /` + r + `__  ` + z + `____/__  __/` + r + `_  ` + z + `____/
	` + r + `_  ` + z + `/ / /` + r + `__  ` + z + `/_   ` + r + `__  ` + z + `/  ` + r + `_` + z + `  /     
	/ /_/ / ` + r + `_  ` + z + `__/   ` + r + `_` + z + `  /   / /___   
	\____/  /_/      /_/    \____/   

	[` + r + `Ultra-Fast-Token-Checker` + z + `]		
	[` + r + `https://github.com/Tainted06/Go-Token-Checker` + z + `]

	Made by YABOI, Modified by Tainted
	
	` + r + `Starting..

`
	fmt.Print(logo)
}

// Clear console
func (Xc *Checker) Cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Reading tokens file
func (Xc *Checker) Read_tokens(files string) ([]string, error) {
	file, err := os.Open(files)
	Xc.Errs(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Writing to file
func (Xc *Checker) OpnFile(file string, token string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0660)
	Xc.Errs(err)
	defer f.Close()
	_, ers := f.WriteString(token + "\n")
	Xc.Errs(ers)
}

// Handling errors
func (Xc *Checker) Errs(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Formatting token
func (Xc *Checker) FormatToken(token string) string {
	if strings.Contains(token, ":") {
		splitToken := strings.Split(token, ":")
		for j := 1; j < len(splitToken); j++ {
			if len(splitToken[0]) < len(splitToken[j]) {
				splitToken[0] = splitToken[j]
			}
		}
		return splitToken[0]
	} else {
		return token
	}
}
