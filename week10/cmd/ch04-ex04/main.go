package main

import (
	"week10/pkg/greeting"
	"week10/pkg/greeting/deutsch"
	korean "week10/pkg/greeting/korean"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()
	korean.Anyung()
}
