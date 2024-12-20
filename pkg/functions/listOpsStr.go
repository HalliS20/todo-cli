package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func SeperateListsN(lis []string, seperator string, n int) [][]string {
	lists := [][]string{}
	for i := 0; i < n; i++ {
		newlist := []string{}
		lists = append(lists, newlist)
	}

	for i := 0; i < len(lis); i++ {
		splitted := strings.Split(lis[i], seperator)
		if len(splitted) != n {
			fmt.Printf("ERROR: Seperation Count Dont match N\n")
			panic("list location: " + strconv.Itoa(i))
		}
		for j := 0; j < n; j++ {
			lists[j] = append(lists[j], splitted[j])
		}
	}
	return lists
}

func ListDStoDI(lists [][]string) [][]int {
	converted := [][]int{}
	for _, innerList := range lists {
		newList := Map(innerList, PanicErrors(strconv.Atoi))
		converted = append(converted, newList)
	}
	return converted
}

func SplitList(lis []string, seperator string) [][]string {
	newList := [][]string{}
	for i := 0; i < len(lis); i++ {
		splitted := strings.Split(lis[i], seperator)
		newList = append(newList, splitted)
	}
	return newList
}
