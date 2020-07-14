package peekcli

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Peek - peeks into a file
func Peek(c *cli.Context) error {
	filename := c.Args().First()
	maxLines := c.Int("max")

	if filename == "" {
		log.Fatal("Please enter a filename")
	}

	peek(filename, maxLines)

	return nil
}

func peek(filename string, maxLines int) error {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < maxLines; i++ {
		if scanner.Scan() {
			fmt.Print(scanner.Text())
			if i < maxLines-1 {
				fmt.Print("\n")
			}
		}
	}

	for {
		reader.ReadString('\n')
		if scanner.Scan() {
			fmt.Print(scanner.Text())
		}
	}
}
