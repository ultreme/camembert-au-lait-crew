package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/camembertaulaitcrew/moi-j-aime-generator"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	length := 20
	var err error
	if len(os.Args) > 1 {
		length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Parse error: %q: %v", os.Args[1], err)
		}
	}

	for i := 0; i < length; i++ {
		fmt.Println(moijaime.Generate())
	}
}
