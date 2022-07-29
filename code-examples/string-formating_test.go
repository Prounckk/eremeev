package code_examples

import (
	"fmt"
	"strconv"
	"testing"
)

func Benchmark_Two_String_Sprintf_With_Integer(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = fmt.Sprintf("%s/%s/%d", "https://eremeev.ca", "resume/", i)
	}
}

func Benchmark_Two_String_Sprintf(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = fmt.Sprintf("%s/%s", "https://eremeev.ca", "resume/")
	}
}

func Benchmark_Two_String_Concatenating_With_Integer(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = "https://eremeev.ca" + "/" + "resume/" + strconv.Itoa(i)
	}
}

func Benchmark_Two_String_Concatenating(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = "https://eremeev.ca" + "/" + "resume/"
	}
}
