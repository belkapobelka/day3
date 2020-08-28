package main

import (
	"context"
	pb "day3/proto/consignment"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
)

const (
	address         = "localhost:50051"
	defaultFilename = "command.json"
)

func parseJSON(file string) (*pb.Command, error) {
	var command *pb.Command
	fileBody, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(fileBody, &command)
	return command, nil
}

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot cpnnact to port: %v", err)
	}
	defer connection.Close()

	client := pb.NewShippingServiceClient(connection)

	command, err := parseJSON(defaultFilename)
	if err != nil {
		log.Fatalf("cannot parse .json file %v", err)
	}
	resp, err := client.CreateCommand(context.Background(), command)
	if err != nil {
		log.Fatalf("cannot create command %v", err)
	}
	log.Printf("Created: %t", resp.Created)
	log.Printf("Body: %v", resp.Command)

	getAll, err := client.GetAllCommands(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("cannot create command %v", err)
	}
	log.Printf("All: %v", getAll)
}
