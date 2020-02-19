package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func hitungKata(s string) map[string]int {
	kata := strings.Fields(s)
	hitung := make(map[string]int)
	for i := range kata {
		hitung[kata[i]]++
	}
	return hitung
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		fmt.Println(hitungKata(scanner.Text()))
		
	}

	

}