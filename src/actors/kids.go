package main

import (
	"actors/kids"
)

func ooStyle(wait bool) {
	bart := kids.CreateKid("Bart")
	lisa := kids.CreateKid("Lisa")
	bart.Start()
	lisa.Start()

	bart.Send(kids.Poke)
	lisa.Send(kids.Poke)

	bart.Send(kids.Feed)
	lisa.Send(kids.Feed)

	bart.Stop(wait)
	lisa.Stop(wait)
}

const Poke = "Poke"
const Feed = "Feed"

func actor(name string) (messages chan string, stop chan bool) {
	messages = make(chan string)
	stop = make(chan bool)

	go func() {
		for {
			output := ""

			select {
			case msg := <-messages:
				switch msg {
				case Poke:
					output = "Ow, quit it!"
				case Feed:
					output = "Burp...Gurgle..."
				}

				println(name + ": " + output)
			case wait := <-stop:
				println(name + ": Bye!")
				if wait {
					println("About to acknowledge stop for " + name)
					stop <- true
					println("Finished sending stop to " + name)
				}
				return
			}
		}
	}()

	return
}

func actorStyle(wait bool) {
	bart, bartStop := actor("Bart")
	lisa, lisaStop := actor("Lisa")

	bart <- Poke
	lisa <- Poke

	bart <- Feed
	lisa <- Feed

	lisaStop <- wait
	bartStop <- wait

	if wait {
		// Wait for the actors to respond to the stop messages
		for i := 0; i < 2; i++ {
			println("i = ", i)
			select {
			case <-bartStop:
				println("Bart is stopped")
			case <-lisaStop:
				println("Lisa is stopped")
			}
		}
	}

}

func main() {
	wait := true
	println("\nOO STYLE:")
	ooStyle(wait)
	println("\nACTORS AND MESSAGES WITH RAW CHANNELS")
	actorStyle(wait)

	panic("This is the end--what is still running?")
}
