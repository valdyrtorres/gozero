package main

import (
	"fmt"
)

func main() {
	fmt.Println("Entrevista")

	for k, v := range "test" {

		fmt.Println(k)
		fmt.Println(v)
	}

	/*
		s1 := []int{1, 2, 3}
		s2 := s1[:]
		s1[0] = 4
		fmt.Print(s2)
	*/

	/*
		s1 := []int{1, 2, 3}
		s2 := s1[:]
		s2 = append(s2, 5)
		s1[0] = 4
		fmt.Print(s2)
	*/

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := s1[1:]
	s3 := s2[1:2:5]
	fmt.Println(s3)

}
