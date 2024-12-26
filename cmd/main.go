package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

const FACTOR = 1024

var multipliers = map[string]float64{
	"KB": 1.0,
	"MB": 2.0,
	"GB": 3.0,
	"TB": 4.0,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: b2h <numOfBytes> -t [KB|MB|GB|TB]")
		return
	}

	input := os.Args[1]
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number\n", input)
		return
	}

	flagSet := flag.NewFlagSet("b2h", flag.ExitOnError)
	unit := flagSet.String("u", "MB", "unit of conversion (KB, MB, ...)")

	flagSet.Parse(os.Args[2:])

	multiplier, ok := multipliers[*unit]
	if !ok {
		fmt.Printf("Error: '%s' is not an available unit\n", *unit)
		return
	}

	var (
		result  = value / math.Pow(FACTOR, multiplier)
		rounded string
	)

	if result < 1 {
		rounded = fmt.Sprintf("%.10f", result)
	} else {
		rounded = fmt.Sprintf("%.2f", result)
	}

	fmt.Printf("%s %s\n", rounded, *unit)
}
