package main

import (
	"context"
	"fmt"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/grpcServer/profile"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4002", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	c := profile.NewProfileServiceClient(conn)
	consumer := profile.Profile{Id: "2001", FirstName: "Archi", LastName: "mukh", Age: "30", Email: "a@gmail.com", Password: "archi@1988", ConfirmPassword: "archi@1988"}
	result, err := c.CreateProfile(context.Background(), &consumer)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
