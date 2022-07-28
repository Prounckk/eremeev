package code_examples

import (
	"fmt"
	"testing"
)


func BenchmarkStringSprintf(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = fmt.Sprintf("%s/%s", "https://eremeev.ca", "resume/")
	}
}

func BenchmarkStringConcatenating(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = "https://eremeev.ca" + "/" + "resume/"
	}
}
