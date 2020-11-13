package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "lawencon.com/cred/client/model"
)

var host = "localhost:1111"
var ctx = context.Background()
var client pb.UserServiceClient

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}
	defer conn.Close()

	client = pb.NewUserServiceClient(conn)

	// register()
	login()
}

func register() {
	defer catchError()
	res, err := client.Register(ctx,
		&pb.Users{Username: "admin", Password: "admin"})

	if err != nil {
		log.Fatal("Error register =>", err)
	}

	log.Println("Success register =>", res)
}

func login() {
	defer catchError()
	res, err := client.Login(ctx,
		&pb.Users{Username: "admin", Password: "admin"})

	if err != nil {
		log.Fatal("Error login =>", err)
	}

	log.Println("Success login =>", res)
}

func validateToken() {}

func catchError() {
	if err := recover(); err != nil {
		log.Println("Error :", err)
	}
}
