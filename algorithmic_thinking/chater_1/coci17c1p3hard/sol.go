package main

import "fmt"

const SIZE = 100000

var (
	passwordMap       = make(map[string]bool, SIZE)
	substringCountMap = make(map[string]int, SIZE*10)
)

func insert(value string) {
	fmt.Println("Inserting", value)
	for i := 0; i < len(value); i++ {
		substringCountMap[value[0:i]]++
	}
}

func getCount(password string) int {
	return substringCountMap[password]
}

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	if n > 0 {
		v := 0
		fmt.Scanf("%d\n", &v)
		if v == 1 {
			password := ""
			fmt.Scanf("%s\n", &password)
			insert(password)
		}
		if v == 2 {
			password := ""
			fmt.Scanf("%s\n", &password)
			if count := getCount(password); count > 0 {
				fmt.Println(count)
			}
		}
	}
}
