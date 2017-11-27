package cgobench

import "unsafe"

/*
#cgo CFLAGS: -O3

int test0() {
	return 1;
}

int test1(int a) {
	return a;
}

int test3(char *p1, char *p2, char *p3) {
	return 1;
}
*/
import "C"

var (
	BA = []byte{0, 1, 2, 3}
	BB = []byte{4, 5, 6, 7}
	BC = []byte{8, 9, 0, 1}
)

// *** CGo ***

func Test0() int {
	return int(C.test0())
}

func Test1(a int) int {
	return int(C.test1(C.int(a)))
}

func Test3(a, b, c []byte) int {
	cA := (*C.char)(unsafe.Pointer(&a[0]))
	cB := (*C.char)(unsafe.Pointer(&b[0]))
	cC := (*C.char)(unsafe.Pointer(&c[0]))
	return int(C.test3(cA, cB, cC))
}

// *** Go ***

//go:noinline
func TestGo0() int {
	return 1
}

//go:noinline
func TestGo1(a int) int {
	return a
}

//go:noinline
func TestGo3(a, b, c []byte) int {
	return 1
}
