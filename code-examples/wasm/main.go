package main

import (
	"fmt"
	"math"
	"syscall/js"
	"time"

	"github.com/prounckk/eremeev/code-examples/wasm/dom"
)

func checkEmail() {
	_ = dom.GetString("email", "value")
	
}

func RoundTime(input float64) int {
	var result float64
	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}
	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)
	return int(i)
}

func YearsOfExperience(this js.Value, args []js.Value) interface{} {
	firstDay := time.Date(2019, time.August, 01, 01, 0, 0, 0, time.UTC)
	diff := time.Since(firstDay)
	yearsOfExperience := RoundTime(diff.Seconds()/31207680)
	return yearsOfExperience
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("YearsOfExperience", js.FuncOf(YearsOfExperience))
	<-make(chan bool)
}
