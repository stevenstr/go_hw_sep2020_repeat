/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Implement sort.Interface interface for People type.
	Sort people in order of increasing age.
	If two people have the same age - sort them by name.
	Add test cases.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"module/domain"
)

func main() {
	ivanIvanovDate1, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ivanIvanovDate2, err := time.Parse("2006-Jan-02", "2003-Aug-05")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	artiomIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//create array of test data
	p := domain.People{
		domain.Person{FirstName: "Ivan", LastName: "Ivanov", BirthDay: ivanIvanovDate1},
		domain.Person{FirstName: "Ivan", LastName: "Ivanov", BirthDay: ivanIvanovDate2},
		domain.Person{FirstName: "Artiom", LastName: "Ivanov", BirthDay: artiomIvanovDate},
	}

	p.Show("Unsorted:")

	var interInterface domain.Interface = p
	sort.Sort(interInterface)

	p.Show("Sorted:")
}
