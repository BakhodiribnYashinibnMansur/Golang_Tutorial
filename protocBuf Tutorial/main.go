package main

import (
	fmt "fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	bond := &Person{
		Name: "James Bond",
		Age:  27,
		SocialFollower: &SocialFollower{
			Youtube: 457,
			Twitter: 784,
		},
	}
	data, err := proto.Marshal(bond)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	newBond := &Person{}
	err = proto.Unmarshal(data, newBond)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newBond)
	fmt.Println(newBond.GetName())
	fmt.Println(newBond.GetAge())
	fmt.Println(newBond.GetSocialFollower())
	fmt.Println(newBond.SocialFollower.GetYoutube())
	fmt.Println(newBond.SocialFollower.GetTwitter())
}
