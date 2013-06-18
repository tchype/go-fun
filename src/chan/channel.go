package main

import "time"


func waitAndReply(channel chan<-string, delay time.Duration) {
	print("Start waiting for " + delay.String() + " at " + time.Now().String() + "...\n")
	time.Sleep(delay)
	channel <- "End wait of " + delay.String() + " at " + time.Now().String() + "\n"
}

func main() {
	c := make(chan string, 2)

	go waitAndReply(c, 5 * time.Second)
	go waitAndReply(c, 1 * time.Second)

	for i:= 0; i < cap(c); i++ {
		output := <-c
		print(output)
	}
}
