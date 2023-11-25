package main

import (
	"fmt"
	"log"

	_ "github.com/cockroachdb/errors/errorspb"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/kentro-lab/cloudy-neigh/proto/api"
)

func main() {
	db, err := pebble.Open("", &pebble.Options{FS: vfs.NewMem()})
	if err != nil {
		log.Fatal(err)
	}
	key := []byte("hello")
	if err := db.Set(key, []byte("world"), pebble.Sync); err != nil {
		log.Fatal(err)
	}
	value, closer, err := db.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", key, value)
	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	msg := api.Vector{Values: []float32{1, 2, 3}}
	fmt.Println("mag: ", msg.String())
}
