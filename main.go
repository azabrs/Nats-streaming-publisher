package main

import (
	"log"
	publisher "nats-publisher/Publisher"

)



func main(){
	pub := publisher.New("publisher", "test-cluster", "orders", 10, 4222, 1)
	n, err := pub.Publish_data()
	if err != nil{
		log.Printf("was publish only %d message", n)
		log.Fatal(err)
	}
}