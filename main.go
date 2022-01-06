package main

import (
	"github.com/sebastiankennedy/head-first-go/ch1"
	"github.com/sebastiankennedy/head-first-go/ch2"
	"github.com/sebastiankennedy/head-first-go/ch3"
	ch4_greeting "github.com/sebastiankennedy/head-first-go/ch4/greeting"
	ch4_greet_dansk "github.com/sebastiankennedy/head-first-go/ch4/greeting/dansk"
	ch4_greeting_deutsch "github.com/sebastiankennedy/head-first-go/ch4/greeting/deutsch"
	ch5_average "github.com/sebastiankennedy/head-first-go/ch5/average"
	ch6_average "github.com/sebastiankennedy/head-first-go/ch6/average"
)

func main() {
	ch1.HelloWorld()
	ch1.Conversisons()
	ch1.Functions()
	ch1.ShortVariableDeclarations()
	ch1.Types()
	ch1.Values()
	ch1.VariableDeclarations()

	ch2.Guess()
	ch2.CallingMethods()
	ch2.PassFail()

	ch3.Paint()

	ch4_greeting.Hi()
	ch4_greeting.Hello()
	ch4_greet_dansk.Hej()
	ch4_greet_dansk.GodMorgen()
	ch4_greeting_deutsch.Hallo()
	ch4_greeting_deutsch.GutenTag()

	ch5_average.Average()

	ch6_average.Average()
}
