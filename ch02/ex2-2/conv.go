package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	oneF = 0.3048
	oneP = 2.20462
)

func print(in string, out *bufio.Writer) {
	t, err := strconv.ParseFloat(in, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conv args failed: %v\n", err)
		os.Exit(1)
	}
	out.WriteString(fmt.Sprintf("%s = %s\n",
		Feet(t).String(), Feet(t).ToMeter().String()))
	out.WriteString(fmt.Sprintf("%s = %s\n",
		Pound(t).String(), Pound(t).ToKilogram().String()))
}

func (f Feet) ToMeter() Meter { return Meter(f * oneF) }
func (f Feet) String() string { return fmt.Sprintf("%.3gft", f) }

func (m Meter) String() string { return fmt.Sprintf("%.3gm", m) }

func (p Pound) ToKilogram() Kilogram { return Kilogram(p * oneP) }
func (p Pound) String() string       { return fmt.Sprintf("%.3glbs", p) }

func (kg Kilogram) String() string { return fmt.Sprintf("%.3gkg", kg) }

func main() {
	result := bufio.NewWriter(os.Stdout)
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			print(arg, result)
		}
		result.Flush()
		return
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		print(sc.Text(), result)
		result.Flush()
	}
}
