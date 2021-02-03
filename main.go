package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
	randomstr "github.com/dxede/randomstring"
)

func exitWith(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func generateToken(length int, characters string, separator string, chunks int) string {
	token := randomstr.RandomString(characters, length)

	if separator != "" {
		if chunks == 0 {
			exitWith("Missing parameter: You must specify the size of a chunk along with the separator.")
		}

		tokenLength := length

		if chunks > (tokenLength / 2) {
			exitWith("Chunk size cannot be greater than the length divided by two.")
		}

		rand.Seed(time.Now().Unix())
		chunkSize := chunks

		groupCount := tokenLength / chunkSize
		separatorRune := []rune(separator)[0]
		generatedRunes := []rune(token)

		for i := 0; i < tokenLength; i++ {
			pos := i + 1
			if pos != tokenLength && pos != 0 && pos%groupCount == 0 {
				generatedRunes[i] = separatorRune
			}
		}

		token = string(generatedRunes)
	}

	return token
}

func main() {
	lengthPtr := flag.Int("length", 20, "Length of the generated password.")
	displayPtr := flag.Bool("display", false, "Display the password in console (not recommended).")
	charactersPtr := flag.String("characters", "*", "Types of characters to be used. Example: 'A0!' => uppercase, numeric and special characters")
	separatorPtr := flag.String("sep", "", "Separator to be used in the password, if any.")
	chunksPtr := flag.Int("chunks", 0, "The number of chunks the password is split into.")
	flag.Parse()

	token := generateToken(
		*lengthPtr,
		*charactersPtr,
		*separatorPtr,
		*chunksPtr,
	)

	if *displayPtr || clipboard.Unsupported {
		if clipboard.Unsupported {
			fmt.Println("Access to clipboard is unsupported on this platform.")
		}

		fmt.Println(token)
	} else {
		clipboard.WriteAll(token)
		fmt.Println("Copied to clipboard.")
	}
}
