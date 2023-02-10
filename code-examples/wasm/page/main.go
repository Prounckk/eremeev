package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"

	"github.com/prounckk/eremeev/code-examples/wasm/dom"
)

func calculateWasm(this js.Value, args []js.Value) any {
	input := dom.GetStringFromElement("calculate")
	if input == "" {
		input = "0"
	}
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return err
	}

	start := time.Now().UnixNano() / int64(time.Millisecond)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res = res + j + i
		}
	}
	end := time.Now().UnixNano() / int64(time.Millisecond)
	executionTime := end - start
	dom.SetValue("total_hours_go", "innerHTML", strconv.Itoa(res))
	dom.SetValue("execution_time_go", "innerHTML", strconv.FormatInt(executionTime, 10))

	return res
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("calculateWasm", js.FuncOf(calculateWasm))
	<-make(chan bool)
}
