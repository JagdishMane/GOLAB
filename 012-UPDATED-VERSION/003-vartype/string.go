package main

import "fmt"

func main() {
	s := "Hello playground"
	s2 := `Hello 
              playground`
	fmt.Println(s)
	fmt.Println(s2)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i :=0; i < len(s); i++ {
		fmt.Printf("%#U\n",s[i]) // Print UTF-8 code point.
	}

	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("At index position %d we have hex %#X\n", i, v)
	}
}