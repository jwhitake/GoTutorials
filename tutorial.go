package main

// import "fmt"

// const uSixteenBitMax float64 = 65535
// const kphMultiple float64 = 1.60934

// type car struct {
// 	gasPedal      uint16
// 	brakePedal    uint16
// 	steeringWhele int16
// 	topSpeedKph   float64
// }

// func (c car) kmh() float64 {
// 	return float64(c.gasPedal) * (c.topSpeedKph / uSixteenBitMax)
// }

// func (c car) mph() float64 {
// 	return float64(c.gasPedal) * (c.topSpeedKph / uSixteenBitMax / kphMultiple)
// }

// func (c *car) newTopSpeed(newSpeed float64) {
// 	c.topSpeedKph = newSpeed
// }

// func newerTopSpeed(c car, speed float64) car {
// 	c.topSpeedKph = speed
// 	return c
// }

// func main() {
// 	aCar := car{
// 		gasPedal:      65000,
// 		brakePedal:    0,
// 		steeringWhele: 12531,
// 		topSpeedKph:   225.0}

// 	fmt.Println(aCar.topSpeedKph)
// 	fmt.Println(aCar.kmh())
// 	fmt.Println(aCar.mph())

// 	//aCar.newTopSpeed(500)
// 	aCar = newerTopSpeed(aCar, 500)

// 	fmt.Println(aCar.kmh())
// 	fmt.Println(aCar.mph())

// }

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/about", aboutHandler)
// 	http.ListenAndServe(":8000", nil)

// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Whoa Go web server")qw
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "John wrote this")
// }

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	fmt.Print("Enter a grade: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(input)

// 	// broken := "This i$ an intere$ting problem"
// 	// fmt.Println(broken)
// 	// replacer := strings.NewReplacer("$", "s")
// 	// fixed := replacer.Replace(broken)
// 	// fmt.Println(fixed)

// }
