package main

import (
	"fmt"
	"github.com/wbjxxzx/stackutil"
	"github.com/wbjxxzx/stringutil"
	"strings"
	"time"
)

func main() {
	fmt.Println(stringutil.Reverse("hello, world"))
	var str string = " This  is an example 	of a  string"
	fmt.Printf("T/F? Does the string %#v have prefix %#v?\n", str, "Th")
	fmt.Print(strings.HasPrefix(str, "Th"))

	s1 := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", s1)

	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str2 := "go1|the ABC of GO|25"
	s2 := strings.Split(str2, "|")
	fmt.Printf("splitted in slice: %v\n", s2)
	for _, val := range s2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	fmt.Printf("%s\n", strings.Join(s1, ":"))

	t := time.Now()
	fmt.Println(t.Format("02 Nov 2006 15:04"))
	s := stackutil.New()
	s.Push(1)
	s.Push("hello world")
	s.Push(struct{ int }{1})
	fmt.Println(s)
}
