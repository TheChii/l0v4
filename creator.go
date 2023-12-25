package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"syscall"
)

var myMap = make(map[string]string)

func main() {
	// Setup signal handler for interrupt signal (Ctrl+C)
	setupSignalHandler()

	fmt.Println("Enter key-value pairs (press Ctrl+C to save and exit):")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter key: ")
		scanner.Scan()
		key := scanner.Text()

		fmt.Print("Enter value: ")
		scanner.Scan()
		value := scanner.Text()

		myMap[key] = value

		fmt.Println("Map updated.")
	}

	// The program will never reach here due to the signal handler.
}

func setupSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		<-sigChan
		saveMapToFile()
		os.Exit(0)
	}()
}

func saveMapToFile() {
	file, err := os.Create("output_map.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "package main")
	fmt.Fprintln(file, "\nvar myHardcodedMap = map[string]string{")

	for key, value := range myMap {
		fmt.Fprintf(file, "\t\"%s\": \"%s\",\n", key, value)
	}

	fmt.Fprintln(file, "}")
	fmt.Println("Map has been saved to 'output_map.go'.")
}
