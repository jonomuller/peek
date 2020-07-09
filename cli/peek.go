package peekcli

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var defaultLinesToPeek = 10

// Peek - peeks into a file
func Peek(c *cli.Context) error {
	filename := c.Args().First()

	if filename == "" {
		log.Fatal("Please enter a filename")
	}

	peek(filename)

	return nil
}

func peek(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < defaultLinesToPeek; i++ {
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
