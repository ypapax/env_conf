package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type ku struct {
	Mu mu
}
type mu struct {
	F1 string `envconfig:"F1"`
	F2 string `envconfig:"F2" required:"true"`
}

func main() {
	var s ku
	if err := envconfig.Process("myapp", &s); err != nil {
		log.Println("err: ", err)
		return
	}
	log.Printf("s: %+v", s)
}
