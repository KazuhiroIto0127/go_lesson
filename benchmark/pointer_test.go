package main

import (
	"testing"
)

func BenchmarkModifyStructByValue(b *testing.B) {
	s := MyStruct{A: 1, B: 2}
	for i := 0; i < b.N; i++ {
		s = modifyStructByValue(s)
	}
}

func BenchmarkModifyStructByPointer(b *testing.B) {
	s := MyStruct{A: 1, B: 2}
	for i := 0; i < b.N; i++ {
		modifyStructByPointer(&s)
	}
}
