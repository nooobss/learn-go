package main

import (
	"fmt"
	"slices"
)

func potongan() {
	fmt.Println("jancok jokowi")
	var slicer []string
	fmt.Println("xyz", slicer, slicer == nil, len(slicer) == 0)

	slicer = make([]string, 3)
	fmt.Println("emp :", slicer, "len:", len(slicer), "cap:", cap(slicer))

	slicer[0] = "a"
	slicer[1] = "b"
	slicer[2] = "c"
	fmt.Println("set :", slicer)
	fmt.Println("get :", slicer[2])
	fmt.Println("len:", len(slicer))

	fmt.Println("len", len(slicer))

	slicer = append(slicer, "d")
	slicer = append(slicer, "e", "f")
	fmt.Println("appnd", slicer)

	c := make([]string, len(slicer))
	copy(c, slicer)
	fmt.Println("copy", "c")

	l := slicer[2:4]
	fmt.Println("slicess", l)

	ll := slicer[:3]
	fmt.Println("slicesss", ll)

	lll := slicer[2:]
	fmt.Println("acara belajar bhs inggris", lll)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
}
