package main

import (
	"flag"
	"fmt"
	
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/friendsofgo/gopherapi/pkg/server"
	
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}