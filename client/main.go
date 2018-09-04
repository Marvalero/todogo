package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Marvalero/todogo/protobuf"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Todogo Client")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("127.0.0.1:18066", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTodoAppClient(conn)

	todo, err := client.CreateTodo(context.Background(), &pb.Todo{Content: "Go on Holidays"})
	if err != nil {
		log.Fatalf("fail to create todo: %v", err)
	}
	fmt.Println("Created: ", todo)

	todo, err = client.CreateTodo(context.Background(), &pb.Todo{Content: "Buy a present"})
	if err != nil {
		log.Fatalf("fail to create todo: %v", err)
	}
	fmt.Println("Created: ", todo)
	stream, err := client.ListTodos(context.Background(), &pb.ListTodosParams{})
	if err != nil {
		log.Fatalf("fail to create list todo: %v", err)
	}
	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListTodos(_) = _, %v", client, err)
		}
		log.Println("List:", todo)
	}

}
