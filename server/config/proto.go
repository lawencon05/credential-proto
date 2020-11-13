package config

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"lawencon.com/credential/model"
	"lawencon.com/credential/service"
)

var userService service.UserService = service.UserServiceImpl{}

// SetProto for init proto server
func SetProto() {
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatal("Failed to listen with err =>", err)
	}

	s := grpc.NewServer()
	model.RegisterUserServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve with err =>", err)
	}
}

type server struct {
	model.UnimplementedUserServiceServer
}

func (*server) Register(ctx context.Context, u *model.Users) (*model.Resp, error) {
	log.Println("Registered from client =>", u)
	var user = model.UsersDb{
		Username: u.Username,
		Password: u.Password,
	}
	err := userService.Register(&user)
	if err != nil {
		log.Println("Failed registered err =>", err)
		return &model.Resp{}, errors.New("Failed register")
	}

	log.Println("Success registered from client =>", u)
	return &model.Resp{Code: "200", Msg: "success"}, nil
}

func (*server) Login(ctx context.Context, u *model.Users) (*model.Resp, error) {
	log.Println("Login from client =>", u)
	var user = model.UsersDb{
		Username: u.Username,
		Password: u.Password,
	}
	err := userService.Login(&user)
	if err != nil {
		log.Println("Failed login err =>", err)
		return &model.Resp{}, errors.New("Invalid Username/Password")
	}

	log.Println("Success login from client =>", u)
	return &model.Resp{Code: "200", Msg: user.Token}, nil
}

func (*server) ValidateToken(ctx context.Context, t *model.Token) (*model.Resp, error) {
	log.Println("Validate Token from client =>", t.Data)
	err := userService.ValidateToken(t.Data)
	if err != nil {
		log.Println("Failed validate token err =>", err)
		return &model.Resp{}, errors.New("Invalid Token")
	}

	log.Println("Success validate token from client =>", t)
	return &model.Resp{Code: "200", Msg: "success"}, nil
}
