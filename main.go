package main

import (
	"github.com/alaraiabdiallah/media-service/data_sources/mongods"
	httpProtocol "github.com/alaraiabdiallah/media-service/protocols/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mongods.CheckConnection()
	httpProtocol.Run()
}
