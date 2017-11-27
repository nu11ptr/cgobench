package cgobench_test

import (
	"testing"

	"github.com/nu11ptr/cgobench"
)

// *** CGo ***

func BenchmarkCGo0Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ret := cgobench.Test0(); ret != 1 {
			panic("Wrong return value")
		}
	}
}

func BenchmarkCGo1IntArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ret := cgobench.Test1(i); ret != i {
			panic("Wrong return value")
		}
	}
}

func BenchmarkCGo3PtrArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ret := cgobench.Test3(cgobench.BA, cgobench.BB, cgobench.BC); ret != 1 {
			panic("Wrong return value")
		}
	}
}

// *** Go ***

func Benchmark0Args(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ret := cgobench.TestGo0(); ret != 1 {
			panic("Wrong return value")
		}
	}
}

func Benchmark1IntArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ret := cgobench.TestGo1(i); ret != i {
			panic("Wrong return value")
		}
	}
}

func Benchmark3PtrArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ret := cgobench.TestGo3(cgobench.BA, cgobench.BB, cgobench.BC); ret != 1 {
			panic("Wrong return value")
		}
	}
}
