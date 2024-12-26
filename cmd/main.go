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
		usage()
		return
	}

	input := os.Args[1]
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		usage()
		return
	}

	flagSet := flag.NewFlagSet("b2h", flag.ExitOnError)
	flagSet.Usage = usage

	unit := flagSet.String("u", "MB", "")
	precision := flagSet.Int("p", 2, "")

	flagSet.Parse(os.Args[2:])

	multiplier, ok := multipliers[*unit]
	if !ok {
		fmt.Printf("Error: '%s' is not an available unit.\n", *unit)
		return
	}

	var (
		result  = value / math.Pow(FACTOR, multiplier)
		rounded string
	)

	format := fmt.Sprintf("%%.%df", *precision)
	rounded = fmt.Sprintf(format, result)

	fmt.Printf("%s %s\n", rounded, *unit)
}

func usage() {
	fmt.Println(`Usage of b2h <numOfBytes>:
  -u string
        unit of conversion (KB, MB, ...) (default "MB")
  -p int
        output precision to control decimal places (default 2)`)
}
