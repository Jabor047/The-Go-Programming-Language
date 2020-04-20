// general conersion converts temperature, length and weight
package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kilo float64
type Pounds float64
type Meters float64
type Feet float64

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		k := Kilo(t)
		p := Pounds(t)
		m := Meters(t)
		fe := Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FtoC(f), c, CtoF(c))
		fmt.Printf("%s = %s, %s = %s\n", k, KtoP(k), p, PtoK(p))
		fmt.Printf("%s = %s, %s = %s\n",m, MtoFe(m), fe, FetoM(fe))
	}
	arg := os.Args[1:]
	if len(arg) == 0 {
		input := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the number to be converted")
		text, _ := input.ReadString(' ')
		t, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		k := Kilo(t)
		p := Pounds(t)
		m := Meters(t)
		fe := Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FtoC(f), c, CtoF(c))
		fmt.Printf("%s = %s, %s = %s\n", k, KtoP(k), p, PtoK(p))
		fmt.Printf("%s = %s, %s = %s\n",m, MtoFe(m), fe, FetoM(fe))
	}
}

func CtoF(c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5 + 32)
}
func FtoC(f Fahrenheit) Celsius{
	return Celsius((f - 32) * 5 / 9)
}
func KtoP(k Kilo) Pounds{
	return Pounds(k * 2.204)
}
func PtoK(p Pounds) Kilo{
	return Kilo(p / 2.204)
}
func FetoM(fe Feet) Meters{
	return Meters(fe * 3.28)
}
func MtoFe(m Meters) Feet{
	return Feet(m / 3.28)
}


func (c Celsius) String() string{ 
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kilo) String() string{
	return fmt.Sprintf("%gkg", k)
}

func (p Pounds) String() string{ 
	return fmt.Sprintf("%glb", p)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gf", f)
}

func (m Meters) String() string{
	return fmt.Sprintf("%gm", m)
}