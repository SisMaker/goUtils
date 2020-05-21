package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++

		}

	}
	return strings[:i]

}

func nonempty2(strings []string) []string {

	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)

		}

	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]

}

func main() {
	s := []int{5, 6, 7, 8, 9}

	fmt.Println(s[:len(s)])

	fmt.Println(remove2(s, 2))

	m := map[string][2]int{
		"a": {1, 2},
		"b": {3, 4},
	}
	s1 := m["a"]
	s2 := s1[:]
	fmt.Printf("sssss %v %v \n", s1, s2)

	ssss := [][][][][][]int{}
	fmt.Printf("%T \n", ssss)
}
