package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Marvalero/todogo/protobuf"
	"google.golang.org/grpc"
)

type todoServer struct {
	todolist []pb.Todo
}

func (ts *todoServer) ListTodos(params *pb.ListTodosParams, stream pb.TodoApp_ListTodosServer) error {
	for _, todo := range ts.todolist {
		if err := stream.Send(&todo); err != nil {
			return err
		}
	}
	return nil
}

func (ts *todoServer) CreateTodo(ctx context.Context, newTodo *pb.Todo) (*pb.Todo, error) {
	ts.todolist = append(ts.todolist, *newTodo)
	return newTodo, nil
}

func main() {
	fmt.Println("Starting Todogo")
	lis, err := net.Listen("tcp", ":18066")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTodoAppServer(grpcServer, &todoServer{})
	grpcServer.Serve(lis)
}
