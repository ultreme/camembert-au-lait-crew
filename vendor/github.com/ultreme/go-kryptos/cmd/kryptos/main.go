package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ultreme/go-kryptos"
)

var decryptFlag bool

func init() {
	flag.BoolVar(&decryptFlag, "decrypt", false, "decrypt")
}

func cryptDecrypt(input string, decrypt bool) string {
	if decrypt {
		return kryptos.Decrypt(input)
	} else {
		return kryptos.Encrypt(input)
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println(cryptDecrypt(strings.Join(flag.Args(), " "), decryptFlag))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(cryptDecrypt(line, decryptFlag))
		}
	}
}
