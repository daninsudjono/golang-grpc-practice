package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	config "chapter-c30/common/config"
	model "chapter-c30/common/model"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
)

// RPC client garage
func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

// RPC client user
func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n001",
		Name:     "Agadyan",
		Password: "putrake7",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	user2 := model.User{
		Id:       "n002",
		Name:     "Skuyranger",
		Password: "mantap",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	garage1 := model.Garage{
		Id:   "q001",
		Name: "Gad's Garage",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	garage2 := model.Garage{
		Id:   "q002",
		Name: "Gad's Second Garage",
		Coordinate: &model.GarageCoordinate{
			Latitude:  46.123123123,
			Longitude: 56.1231313123,
		},
	}

	// Testing the actual user rpc client
	user := serviceUser()
	fmt.Println("\n", "====> user test")

	// register user1
	user.Register(context.Background(), &user1)
	// register user2
	user.Register(context.Background(), &user2)

	// show all users

	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	// Testing the actual garage rpc client

	garage := serviceGarage()
	fmt.Println("\n", "====> garage test")

	//add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	//add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})

	// show all garages of user 1

	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

}
