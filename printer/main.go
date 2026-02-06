package main

import (
	"fmt"
	"log"
	"os"
)

// Create a Printer interface with Print(msg string).
// Implement ConsolePrinter and FilePrinter.

type Printer interface {
	Print(msg string)
}

type ConsolePrinter struct {
}

func (cp *ConsolePrinter) Print(msg string) {
	fmt.Println(msg)
}

type FilePrinter struct {
}

func (fp *FilePrinter) Print(msg string) {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write([]byte(msg))
}

func main() {
	var cprinter Printer
	cprinter = &ConsolePrinter{}

	cprinter.Print("This is Console log")

	var filePrinter Printer
	filePrinter = &FilePrinter{}

	filePrinter.Print("This will be saved in a file called output.txt \n")
}
