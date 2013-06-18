package main

import "actors/kids"

func ooStyle() {
	bart := kids.CreateKid("Bart")
	lisa := kids.CreateKid("Lisa")
	bart.Start()
	lisa.Start()

	bart.Send(kids.Poke)
	lisa.Send(kids.Poke)

	bart.Send(kids.Feed)
	lisa.Send(kids.Feed)

	bart.Stop(true)
	lisa.Stop(true)
}

func msgStyle() {
	bart, _ := kids.CreateKid("Bart").Start()
	lisa, _ := kids.CreateKid("Lisa").Start()

	bart <- kids.Poke
	lisa <- kids.Poke

	bart <- kids.Feed
	lisa <- kids.Feed
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
			case <-stop:
				println(name + ": Bye!")
				stop <- true
				return
			}
		}
	}()

	return
}

func actorStyle() {
	bart, _ := actor("Bart")
	lisa, _ := actor("Lisa")

	bart <- Poke
	lisa <- Poke

	bart <- Feed
	lisa <- Feed
}

func main() {
	println("\nOO STYLE:")
	ooStyle()
	println("\nOO AND MESSAGE STYLE")
	msgStyle()
	println("\nACTORS AND MESSAGES WITH RAW CHANNELS")
	actorStyle()
}
