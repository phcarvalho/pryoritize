package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const defaultText = "Use `pry -p \"some new priority\"` to change your priority"

func main() {
	newPriority := flag.String("p", "", "change the priority to be shown")
	flag.Parse()

	text := *newPriority
	if text != "" {
		changePriority(*newPriority)
	} else {
		text = readFile()
	}

	printText(os.Stdout, text)
}

func printText(w io.Writer, text string) {
	fmt.Fprintf(w, "\n > %s\n\n", text)
}

func getFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/.pry", home)
}

func changePriority(priority string) {
	filepath := getFilePath()
	err := os.WriteFile(filepath, []byte(priority), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() string {
	filepath := getFilePath()

	file, err := os.Open(filepath)
	if err != nil {
		err := os.WriteFile(filepath, []byte(defaultText), 0644)
		if err != nil {
			log.Fatal(err)
		}
		return defaultText
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(buf)
}
