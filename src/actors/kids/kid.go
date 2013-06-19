package kids

type Message string

const Poke Message = "Poke"
const Feed Message = "Feed"

type Kid struct {
	name     string
	messages chan Message
	stop     chan bool
}

func (kid Kid) Name() string { return kid.name }

func (kid Kid) Start() (chan<-Message, chan bool) {
	go func() {
		for {
			var output string

			select {
			case msg := <-kid.messages:
				switch msg {
				case Poke:
					output = "Ow, quit it!"
				case Feed:
					output = "Gurgle...Burp..."
				}

				println("OO " + kid.Name() + ": " + output)
			case wait:= <-kid.stop:
				println("OO " + kid.Name() + ": Bye!")
				if wait {
					kid.stop <- true //ackowledge the stop
				}
				return
			}
		}
	}()

	return kid.messages, kid.stop
}

func (kid Kid) Send(message Message) {
	kid.messages <- message
}

func (kid Kid) Stop(wait bool) {
	kid.stop <- wait

	if wait {
		<-kid.stop
	}
}

func CreateKid(name string) Kid {
	messages := make(chan Message)
	stop := make(chan bool)

	kid := Kid{name: name, messages: messages, stop: stop}

	return kid
}
