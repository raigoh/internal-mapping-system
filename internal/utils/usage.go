package utils

import "fmt"

func PrintUsage() {
	// Define ANSI color codes
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
		Cyan   = "\033[36m"
	)

	fmt.Println(string(Green) + "Usage:" + string(Reset))
	fmt.Println(string(Yellow) + "  go run . <network_map> <start_station> <end_station> <number_of_trains>" + string(Reset))
	fmt.Println()
	fmt.Println(string(Green) + "Arguments:" + string(Reset))
	fmt.Println(string(Cyan) + "  <network_map>      " + string(Reset) + "Path to the network map file")
	fmt.Println(string(Cyan) + "  <start_station>    " + string(Reset) + "Name of the start station")
	fmt.Println(string(Cyan) + "  <end_station>      " + string(Reset) + "Name of the end station")
	fmt.Println(string(Cyan) + "  <number_of_trains> " + string(Reset) + "Number of trains (positive integer)")
	fmt.Println()
	fmt.Println(string(Green) + "Flags:" + string(Reset))
	fmt.Println(string(Cyan) + "  -h, --help         " + string(Reset) + "Show this help message")
	fmt.Println()
	fmt.Println(string(Green) + "Testing:" + string(Reset))
	fmt.Println(string(Cyan) + "  To run the tests, use one of the following commands:" + string(Reset))
	fmt.Println(string(Cyan) + "    go test          " + string(Reset) + "Run the tests with default output")
	fmt.Println(string(Cyan) + "    go test -v       " + string(Reset) + "Run the tests with detailed output")
}
