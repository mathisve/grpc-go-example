package main

import (
	"context"
	"fmt"
	"github.com/Mathisco-01/grpc-go-example/hash"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()

	c := hash.NewHashServiceClient(conn)

	for {
		var text string
		fmt.Print("Enter text or type 'exit' to leave: ")
		fmt.Scanln(&text)

		if text == "exit" {
			break
		}

		response, err := c.Sha1Message(context.Background(), &hash.Message{Body: text})
		if err != nil {
			log.Fatalf("Something went wrong: %s", err)
		}
		log.Printf("SHA1: %s", response.Body)

		response, err = c.Sha265Message(context.Background(), &hash.Message{Body: text})
		if err != nil {
			log.Fatalf("Something went wrong: %s", err)
		}
		log.Printf("SHA265: %s", response.Body)

		response, err = c.Sha512Message(context.Background(), &hash.Message{Body: text})
		if err != nil {
			log.Fatalf("Something went wrong: %s", err)
		}
		log.Printf("SHA512: %s", response.Body)
	}
}
