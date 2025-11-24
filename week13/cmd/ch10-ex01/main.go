package main

import (
	"fmt"
	"log"
	calendar "week13/pkg/calendar"
)

func main() {
	today := calendar.Event{} // Date가 임베딩 되어 있음

	err := today.SetTitle("Final Exam D-14")
	if err != nil {
		log.Fatal(err)
	}

	err = today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}

	err = today.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n%d년 %d월 %d일\n", today.GetTitle(), today.GetYear(), today.GetMonth(), today.GetDay())
}
