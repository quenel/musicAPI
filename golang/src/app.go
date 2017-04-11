package main

import(
	"usecases"
	"controller/musixmatch"
	"controller/rpcserver"
	"log"
)

func main() {
	// Defining controllers
	musixmatchHandler := musixmatch.MusixmatchHandler{}

	// Defining usecases
	musicRepo := usecases.NewMusicRepo(musixmatchHandler)
	rpcserver.Server.MusicRepo = musicRepo
	
	log.Println("Music API Microservice enabled")
	select{}
}