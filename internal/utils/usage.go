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
	fmt.Println(string(Cyan) + "  1. From the cmd folder:" + string(Reset))
	fmt.Println(string(Yellow) + "  go run . ../stations/<network_map> <start_station> <end_station> <number_of_trains>" + string(Reset))
	fmt.Println(string(Cyan) + "  2. From the root folder:" + string(Reset))
	fmt.Println(string(Yellow) + "  go run cmd/main.go stations/<network_map> <start_station> <end_station> <number_of_trains>" + string(Reset))
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
	fmt.Println(string(Green) + "Running the Program:" + string(Reset))
	fmt.Println(string(Cyan) + "  1. Navigate to the project root directory" + string(Reset))
	fmt.Println(string(Cyan) + "  2. Run the following command:" + string(Reset))
	fmt.Println(string(Yellow) + "     go run cmd/main.go stations/network.map waterloo st_pancras 4" + string(Reset))
	fmt.Println()
	fmt.Println(string(Green) + "Displaying Help:" + string(Reset))
	fmt.Println(string(Cyan) + "  To show this help message:" + string(Reset))
	fmt.Println(string(Cyan) + "  1. From the root folder:" + string(Reset))
	fmt.Println(string(Yellow) + "     go run cmd/main.go -h" + string(Reset))
	fmt.Println(string(Cyan) + "  2. From the cmd folder:" + string(Reset))
	fmt.Println(string(Yellow) + "     go run . -h" + string(Reset))
	fmt.Println()
	fmt.Println(string(Green) + "Testing:" + string(Reset))
	fmt.Println(string(Cyan) + "  To run the tests:" + string(Reset))
	fmt.Println(string(Cyan) + "  1. Navigate to the project root directory" + string(Reset))
	fmt.Println(string(Cyan) + "  2. Run one of the following commands:" + string(Reset))
	fmt.Println(string(Yellow) + "     go test ./tests                 " + string(Reset) + "Run all tests with default output")
	fmt.Println(string(Yellow) + "     go test ./tests -v              " + string(Reset) + "Run all tests with verbose output")
}
