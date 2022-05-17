package code_examples

import (
	"testing"
)

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

func mergeTwoSliceOfFloatsWithCopy(floatsToIterate []float64, floatsToAdd []float64) (int, error) {
	sliceLength := len(floatsToIterate) * len(floatsToAdd)
	newSlice := make([]float64, sliceLength)
	if len(floatsToIterate) == 0 || len(floatsToAdd) == 0 {
		return len(newSlice), nil
	}
	var i int
	for _, _ = range floatsToIterate {
		i = +copy(newSlice[i:], floatsToAdd)
	}

	return len(newSlice), nil
}

func Benchmark_WithAppend(b *testing.B) {
	type args struct {
		sliceOfFloats []float64
		floatsToAdd   []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "slice of 100 floats",
			args: args{
				sliceOfFloats: []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11},
				floatsToAdd:   []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11}},
		},

		{
			name: "slice of 10000 floats",
			args: args{
				sliceOfFloats: []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11},
				floatsToAdd:   []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11}},
		},
	}
	for _, bb := range tests {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := mergeTwoSliceOfFloatsWithAppend(bb.args.sliceOfFloats, bb.args.floatsToAdd)
				if err != nil {
					return
				}
			}
		})
	}
}

func Benchmark_withAppendToDefinedSlice(b *testing.B) {
	type args struct {
		sliceOfFloats []float64
		floatsToAdd   []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "slice of 100 floats",
			args: args{
				sliceOfFloats: []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11},
				floatsToAdd:   []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11}},
		},

		{
			name: "slice of 10000 floats",
			args: args{
				sliceOfFloats: []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11},
				floatsToAdd:   []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11}},
		},
	}
	for _, bb := range tests {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := mergeTwoSliceOfFloatsWithAppendToSliceWithSize(bb.args.sliceOfFloats, bb.args.floatsToAdd)
				if err != nil {
					return
				}
			}
		})
	}
}

func Benchmark_withCopy(b *testing.B) {
	type args struct {
		sliceOfFloats []float64
		floatsToAdd   []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "slice of 100 floats",
			args: args{
				sliceOfFloats: []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11},
				floatsToAdd:   []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11}},
		},

		{
			name: "slice of 10000 floats",
			args: args{
				sliceOfFloats: []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11},
				floatsToAdd:   []float64{1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11, 1.2, 1.3, 1.4, 5, 6, 7, 8, 9, 10, 11}},
		},
	}
	for _, bb := range tests {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := mergeTwoSliceOfFloatsWithCopy(bb.args.sliceOfFloats, bb.args.floatsToAdd)
				if err != nil {
					return
				}

			}
		})
	}
}
