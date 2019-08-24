package contest

import (
	"log"
	"testing"
)

func TestFileSystem_Create(t *testing.T) {
	obj := Constructor()
	// param_1 := obj.Create("/a", 1);
	// log.Println(param_1)
	// param_2 := obj.Get("/a");
	// log.Println(param_2)

	param_1 := obj.Create("/leet", 1)
	log.Println(param_1)

	param_2 := obj.Create("/leet/code", 2)
	log.Println(param_2)

	param_3 := obj.Get("/leet/code")
	log.Println(param_3)

	create := obj.Create("/c/d", 1)
	log.Println(create)

	get := obj.Get("/c")
	log.Println(get)

}
