package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

type Small struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func value() Small {
	return Small{
		a: 1, b: 1, c: 1,
		d: "test", e: "hello", f: "world",
		g: 1.0, h: 2.222222, i: 66.666666,
	}
}

func ptr() *Small {
	return &Small{
		a: 1, b: 1, c: 1,
		d: "test", e: "hello", f: "world",
		g: 1.0, h: 2.222222, i: 66.666666,
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var s Small
	f, err := os.Create("stack.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = value()
	}
	trace.Stop()
	b.StopTimer()
	_ = fmt.Sprintf("%v", s.a)

}

func BenchmarkHeap(b *testing.B) {
	var s *Small
	f, err := os.Create("heap.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		s = ptr()
	}

	trace.Stop()

	b.StopTimer()
	_ = fmt.Sprintf("%v", s.a)

}
