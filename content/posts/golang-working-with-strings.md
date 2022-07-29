+++
title = "Golang: Strings Concatenating "
date = "2022-07-29"
author = "prounckk"
tags = ["golang"]
keywords = ["golang", "benchmark","string", "concatenating"]
description = ""
showFullContent = false
+++

Working with strings is part of our daily routine, and most of the time, developers do not care what is happening under the hood. But let's see how our daily code changes might impact performance. (Yes, it's a copy past from the previous post, thanks for noticing!)

While working on the article, I wanted to get and then share with you my findings. It's something that most devs know,
but who am I ;) a guy with no CS degree... so don't judge me


## What is a string?

String in Golang is a sequence of bytes, where each byte of string can be accessed by index. Same as array or slice. You can also check the length of the string by the len() function. The important thing is that strings are immutable, so you can't change the value of a string. You can only create a new one or convert it to a rune slice, modify it and convert it back to a string. Magic! (I hope so)

```go
package main

import (
	"fmt"
)

func main() {
	hello := "Hello World"
	fmt.Println(hello[0])   // prints 72. Why? bcs it's uint8 representation of the letter "H"
	fmt.Printf("%c\n", hello[0])  // prints H
	fmt.Println(hello[1])   // prints 101
	fmt.Printf("%c\n", hello[1])  // prints e. You know why, right? 
	fmt.Println(len(hello)) // prints 11

	// hello[0] = "h"      // throws error "cannot assign to hello[0] (value of type byte)" AAAA what to do!?
	runes := []rune(hello)    // convert to rune slice
	runes[0] = 'h'        // now we can change the first element of the string
	hello = string(runes) // convert to string again
	fmt.Println(hello)    // prints "hello World"
}

```

## What is a rune?
Stop! Wait. What is a `[]rune(hello)` in the code? What is the uint8 representation of the letter "H"? Let's see the code:

```go
package main

import (
	"fmt"
)

func main() {
	str := "Français québécois" // i'm from Montreal, Canada. We speak French here, special French.
	fmt.Println(str) // prints "Français québécois". As expected.
	fmt.Println(len(str)) // prints 21. WHAT? 21? why?
	runes := []rune(str) // convert to rune slice
	fmt.Println(len(runes)) // prints 18. Do you see why?
}
```

So `ç` and `é` are special characters. Each one represented by two bytes in memory. The function `len()` counts the number of bytes, not the number of characters. So the string is 21 bytes long, but it's 18 runes long. You can see it as a `char` data type(but not really) if you wish. In "normal" life, it's not a problem because most of the time, we check the length of the string to see if it's empty or not. But please, use the rune slice to access the characters. Want to learn more about runes? [Check out the docs](https://golang.org/pkg/unicode/utf8/#Rune) or [this article](https://www.educative.io/answers/what-is-the-rune-type-in-golang).


## Concatenating strings

As i mentioned before, strings are immutable, so you can't change the value of a string. But most of the time, you don't really need to modify it; you just want to concatenate some strings or format them. For example, working on API requests, you want to concatenate the base URL and the endpoint, and you probably want to pass some params with the URL


```go
package main

import (
	"fmt"
	"strconv"
)
func main() {
    baseURL := "http://example.com"
    endpoint := "/api/v1/users"
    url := baseURL + endpoint // this is a concatenation
    url2 := fmt.Sprintf("%s%s", baseURL, endpoint) // this is also a concatenation
	urlWithDigits := baseURL + endpoint + "?=" + strconv.Itoa(100) // this is a concatenation with a digit
    urlWithDigits2 := fmt.Sprintf("%s%s?=%d", baseURL, endpoint, 100) // this is also a concatenation with a digit

    fmt.Println(url) //http://example.com/api/v1/users
	fmt.Println(url2) //http://example.com/api/v1/users
	fmt.Println(urlWithDigits) // http://example.com/api/v1/users?=100
	fmt.Println(urlWithDigits2) // http://example.com/api/v1/users?=100
}
```
So you can see there are at least two ways to concatenate strings in Golang. One uses the + operator, and the other uses the fmt.Sprintf() function. You can find other ways here: [Different ways to concatenate two strings in Golang](https://www.geeksforgeeks.org/different-ways-to-concatenate-two-strings-in-golang/) 

But which one is the best? Let's run some benchmarks to see which one is the best.

```aidl
go test -bench=String -benchmem
goos: darwin
goarch: arm64
pkg: github.com/prounckk/eremeev/code-examples
Benchmark_Two_String_Sprintf_With_Integer-8             11333443	102.7 ns/op		71 B/op		1 allocs/op
Benchmark_Two_String_Sprintf-8                          20571031	59.68 ns/op		48 B/op		1 allocs/op
Benchmark_Two_String_Concatenating_With_Integer-8       27059809	44.75 ns/op		71 B/op		1 allocs/op
Benchmark_Two_String_Concatenating-8                    418224721	2.970 ns/op		16 B/op		0 allocs/op
PASS
ok      github.com/prounckk/eremeev/code-examples       6.695s
```

Ok! now are talking! The first benchmark is the one that uses the fmt.Sprintf() function. The second set of benchmarks uses the + operator. And as you can see, the + operator is faster than the fmt.Sprintf() function and does not require any memory allocation if we don't make any conversions from integers to strings. But even with string conversion, the + operator is faster than the fmt.Sprintf() function. 

So, what is the best way to concatenate strings in Golang? The best way is how your team agreed on it! I might be mistaken, but fmt.Sprintf() looks sexy and has better readability than the + operator. But again. talk to your lead dev and your team and have a conclusion about which one is the best fit for your codebase.

Cheers! Happy codding! 


## Links 

[Benchmarks code is here](https://github.com/Prounckk/eremeev/blob/main/code-examples/string-formating_test.go), please feel free to modify and run your personal benchmarks

