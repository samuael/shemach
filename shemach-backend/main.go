// package main

// import (
// 	"log"
// 	"time"

// 	// "github.com/samuael/Project/CarInspection/platforms/helper"
// 	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
// 	"github.com/samuael/shemach/shemach-backend/platforms/helper"
// )

// func main() {
// 	// This Main funciton is a place where we are going to test some functions
// 	// the main entry point for the application is the main.go file in the cmd/server folder.
// 	println(time.Now().Unix())
// 	// for(int i=0;i<=100;++i) printf("\r[%3d%%]",i);

// 	log.Println(string(helper.MarshalThis(&model.Infoadmin{})))
// 	val, _ := helper.HashPassword("admin")
// 	println(string(val))
// }

package main

import (
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
)

func main() {
	// fmt.Println(fqdn.)
	// val := helper.MarshalThis(&model.Address{UniqueAddressName: "Ambab Ber", Zone: "Assosa", Woreda: "01", City: "Assosa", Kebele: "04", Latitude: 45898955.44343, Longitude: 432423423423.99, Region: "Benishangul"})
	// log.Println(string(val))

	// println(time.Now().Add(time.Hour * 1566).Unix())

	message := &model.Message{
		Targets: []int{-1},
		Lang:    "amh",
		Data:    "-- ",
	}
	// translation.TranslateIt("something")
	println(string(helper.MarshalThis(message)))
}
