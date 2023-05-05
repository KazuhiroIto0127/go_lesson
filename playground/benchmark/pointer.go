package main

type MyStruct struct {
	A int
	B int
}

func modifyStructByValue(s MyStruct) MyStruct {
	s.A = 10
	s.B = 20
	return s
}

func modifyStructByPointer(s *MyStruct) {
	s.A = 10
	s.B = 20
}
