package main

import (
	"fmt"
	calendar "week13/pkg/calendar"
)

func main() {
	today := calendar.Date{}
	// today.year = 2025 이거 쓰지 마셈
	today.SetYear(2025)
	today.SetMonth(11)
	today.SetDay(24)
	// fmt.Println(today.GetYear(), "년", today.GetMonth(), "월", today.GetDay(), "일")
	fmt.Printf("%d년 %d월 %d일\n", today.GetYear(), today.GetMonth(), today.GetDay())
}
