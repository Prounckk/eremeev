+++
title = "Golang: Append vs Copy "
date = "2022-05-17"
author = "prounckk"
tags = ["golang","data structure"]
keywords = ["golang", "benchmark", "data structure", "append vs copy"]
description = "Working with data structures is part of our daily routine, and most of the time, developers do not care what is happening
under the hood. We add or delete elements, sort them or iterate through them to find what we need. But let's see how our
daily code changes might impact performance."
showFullContent = false
+++

Working with data structures is part of our daily routine, and most of the time, developers do not care what is happening
under the hood. We add or delete elements, sort them or iterate through them to find what we need. But let's see how our
daily code changes might impact performance.

While working on the article, I wanted to get and then share with you my findings. It's something that most devs know,
but who am I ;) a guy with no CS degree... so don't judge me


## Array vs Slice

An array has a predefined length, and it can NOT grow.

Array
`var x [5]int
`

A slice is a pointer to an array it can change the size.

Slice
`var x []int
`

### Fun Fact:

You might expect that working with Arrays is faster because it sits in the known and dedicated memory segment, while
slice is a pointer and when it's growing it forces memory allocation. It is right! But the array is slightly slower if you
work with local scope!

```aidl
goos: darwin
goarch: arm64
BenchmarkSliceGlobal-8        926196          1257.0 ns/op         0 B/op          0 allocs/op
BenchmarkArrayGlobal-8       2110324           567.0 ns/op         0 B/op          0 allocs/op
BenchmarkSliceLocal-8        2275382           535.0 ns/op         0 B/op          0 allocs/op
BenchmarkArrayLocal-8        1802491           647.4 ns/op         0 B/op          0 allocs/op
```

Check the code on your local and compare the results. The code is available at [StackOverflow](https://stackoverflow.com/questions/30525184/array-vs-slice-accessing-speed/72204705#72204705)

## Append vs Copy

Now let's compare two main functions that we can use to make a slice grow

`slice = append(slice, anotherSlice...)`

AND

`copy(slice, anotherSlice)`

### Append

Most of the time, devs are using append(). It's simple, elegant and provides us with good readability...

```go
func mergeTwoSliceOfFloatsWithAppend(floatsToIterate []float64, floatsToAdd []float64) (int, error) {
    var newSlice []float64
    if len(floatsToIterate) == 0 || len(floatsToAdd) == 0 {
        return len(newSlice), nil
    }
    for _, _ = range floatsToIterate {
        newSlice = append(newSlice, floatsToAdd...)
    }
    return len(newSlice), nil
}
```

### Copy

On the other hand, copy() needs some work to be done before it performs its action. First of all, the new slice's length needs to be specified and then track the index where the new slice has to be added. Kind of complicated, right?

```go
func mergeTwoSliceOfFloatsWithCopy(floatsToIterate []float64, floatsToAdd []float64) (int, error) {
    sliceLength := len(floatsToIterate) * len(floatsToAdd)
    newSlice := make([]float64, sliceLength)
    if len(floatsToIterate) == 0 || len(floatsToAdd) == 0 {
        return len(newSlice), nil
    }
    var i int
    for _, _ = range floatsToIterate {
        i =+ copy(newSlice[i:], floatsToAdd)
    }
    return len(newSlice), nil
}

```
So, let's compare the results!
Copy works not just faster but also requires only one memory allocation!

```aidl
Benchmark_withCopy/slice_of_100_floats      88.00 ns/op	     896 B/op    1 allocs/op
Benchmark_withCopy/slice_of_10000_floats     4456 ns/op	   81920 B/op    1 allocs/op
Benchmark_WithAppend/slice_of_100_floats    252.5 ns/op	    2480 B/op    5 allocs/op
Benchmark_WithAppend/slice_of_10000_floats  18985 ns/op	  356224 B/op   12 allocs/op
```

Looks good, right?
But what if the magic is in the `newSlice := make([]float64, sliceLength)`? Let's see

```go
func mergeTwoSliceOfFloatsWithAppendToSliceWithSize(floatsToIterate []float64, floatsToAdd []float64) (int, error) {
	sliceLength := len(floatsToIterate) * len(floatsToAdd)
	newSlice := make([]float64, sliceLength)
	if len(floatsToIterate) == 0 || len(floatsToAdd) == 0 {
		return len(newSlice), nil
	}
	for _, _ = range floatsToIterate {
		newSlice = append(newSlice, floatsToAdd...)
	}
	return len(newSlice), nil
}
```

Running our benchmark on the same set of floats

```aidl
Benchmark_withCopy/slice_of_100_floats              88.00 ns/op	     896 B/op    1 allocs/op
Benchmark_withCopy/slice_of_10000_floats             4456 ns/op	   81920 B/op    1 allocs/op
Benchmark_WithAppend/slice_of_100_floats            252.5 ns/op	    2480 B/op    5 allocs/op
Benchmark_WithAppend/slice_of_10000_floats          18985 ns/op	  356224 B/op   12 allocs/op
Benchmark_withAppendToDefinedSlice/100_floats       193.8 ns/op	    2688 B/op    2 allocs/op
Benchmark_withAppendToDefinedSlice/10000_floats     23807 ns/op	  507904 B/op    4 allocs/op
```
Append with a predefined slice seems to be working better. It does less memory allocation. Did it help? Not really! It's slower than regular append, and it takes 2x times more spaces in memory now than it needs. Append() do not use the slice with predefined length but add a new slice at its end, so the first part of our new slice is empty but takes space in memory. Just try to run the code with a debugger or print the function's output.

[Code is here](https://github.com/Prounckk/eremeev/blob/main/code-examples/append-vs-copy_test.go), please feel free to modify and run your personal benchmarks

What can I say? I was hoping to get similar results, but tests never lie. copy() works faster and needs fewer memory allocations, but the readability of the code is not the selling point of copy() function. I will use append() in my code if I foresee that nobody will pass a massive chunk of data. Most developers know this function, and speaking of trade-offs, I prefer not to spend expensive dev time understanding the code, but add a few bucks on memory. But if the code expects to be performance-focused, then copy(), and some additional lines of comments would be my to-go option. 