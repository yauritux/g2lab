package main

import "fmt"

func main() {
	fmt.Println("various slicing way")
	fmt.Println("-------------------")
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	printBytes(b) // Yields: len=6 cap=6 [g o l a n g]
	b = b[1:4]
	printBytes(b) // Yields: len=3 cap=5 [o l a]
	b = b[:2]
	printBytes(b) // Yields: len=2 cap=5 [o l]
	b = b[:4]
	printBytes(b) // Yields: len=4 cap=5 [o l a n]

	fmt.Println()
	fmt.Println("using default offset (0) and length (len(x))")
	fmt.Println("--------------------------------------------")
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:]
	printStrings(s) // Yields: len=3 cap=3 [Лайка Белка Стрелка]
	
	d := []byte{'r', 'o', 'a', 'd'}
	printBytes(d) // Yields: len=4 cap=4 [r o a d]
	e := d[2:] 
	printBytes(e) // Yields: len=2 cap=2 [a d]
	e[1] = 'm'
	printBytes(e) // Yields: len=2 cap=2 [a m]
	printBytes(d) // Yields: len=4 cap=4 [r o a m]

	fmt.Println()
	fmt.Println("extend the len based on its' capacity")
	fmt.Println("-------------------------------------")
	b1 := make([]byte, 5)
	printBytes(b1) // Yields: len=5 cap=5 [     ]
	b1 = b1[2:4]
	printBytes(b1) // Yields: len=2 cap=3 [  ]
	b1 = b1[:cap(s)] // here we extend the len based on its' capacity
	printBytes(b1) // Yields: len=3 cap=3 [   ]

	// slice copy implementation
	fmt.Println()
	fmt.Println("slice copy implementation")
	fmt.Println("-------------------------")
	newb1 := make([]byte, len(b1), (cap(b1)+1) * 2) // +1 in the case of cap(b1) = 0
	for i := range b1 {
		newb1[i] = b1[i]
	}
	b1 = newb1
	printBytes(b1) // Yields: len=3 cap=8 [   ]
	printBytes(newb1) // Yields: len=3 cap=8 [   ]

	// copy implementation (using built-in function)
	fmt.Println()
	fmt.Println("using built-in copy function")
	fmt.Println("----------------------------")
	newb2 := make([]byte, len(b1), (cap(b1)+1) * 2) 
	copy(newb2, b1)
	b1 = newb2
	printBytes(b1) // Yields: len=3 cap=18 [   ]
	printBytes(newb2) // Yields: len=3 cap=18 [   ]

	// using custom appendByte function for complete control
	fmt.Println()
	fmt.Println("using custom appendByte function for complete control")
	fmt.Println("-----------------------------------------------------")
	p := []byte{2, 3, 5}
	p = appendByte(p, 7, 11, 13)
	printBytesAsNumber(p) // Yields: len=6 cap=8 [2 3 5 7 11 13]

	// using built-in append function
	fmt.Println()
	fmt.Println("using built-in append function")
	fmt.Println("------------------------------")
	a := make([]int, 1)
	printInts(a) // Yields: len=1 cap=1 [0]
	a = append(a, 1, 2, 3)
	printInts(a) // Yields: len=4 cap=4 [0 1 2 3]
	str1 := []string{"Yauri", "Tanaka"}
	str2 := []string{"Harmonie", "Noburo"}
	str1 = append(str1, str2...) // same as append(str1, str2[0], str2[1])
	printStrings(str1) // Yields: len=4 cap=4 [Yauri Tanaka Harmonie Noburo]

	// using custom appendString function
	fmt.Println()
	fmt.Println("using custom appendString function")
	fmt.Println("----------------------------------")
	q := []string{"yauri", "jackie", "liu han"}
	q = appendString(q, "tang wei", "Shibata")
	printStrings(q) // Yields: len=5 cap=12 [yauri jackie liu han tang wei Shibata]

}

func appendString(slice[] string, data ...string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]string, (n + 1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:], data)
	return slice
}

func appendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n + 1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:], data)
	return slice
}

func printBytes(b []byte) {
	fmt.Printf("len=%d cap=%d %c\n", len(b), cap(b), b)
}

func printBytesAsNumber(b []byte) {
	fmt.Printf("len=%d cap=%d %d\n", len(b), cap(b), b)
}

func printInts(i[] int) {
	fmt.Printf("len=%d cap=%d %d\n", len(i), cap(i), i)
}

func printStrings(s []string) {
	fmt.Printf("len=%d cap=%d %s\n", len(s), cap(s), s)
}

// custom implementation of AppendByte function
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice) 
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func appendByteWithTrailingDefaultGrowthValue(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n + 1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	//slice = slice[0:n]
	copy(slice[m:n], data)	
	return slice
}
