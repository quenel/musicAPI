package rpcserver

import(
	"net/http"
	"log"
	"fmt"
	"github.com/labstack/echo"
)

const (
	GetArtistsIdsRPCName = "getArtistsIds"
	GetArtistRPCName = "getArtist"
	GetArtistCloudWordRPCName = "getArtistCloudWord"
)

func GetArtistsIdsRPC(c echo.Context) error {
	return c.JSON(http.StatusOK, Server.MusicRepo.GetArtistsIds())
}

func GetArtistRPC(c echo.Context) error {
	query := struct{
		Id 	int 	`json:"id"`
	}{}
	if err := c.Bind(&query) ; err != nil {
		log.Println("error binding request, error:", err)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}

	artist, exist := Server.MusicRepo.GetArtist(query.Id)
	if !exist {
		log.Println("error :", fmt.Errorf("track id %d does not exists", query.Id))
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	res := map[string]interface{}{
		"id" : artist.Id,
		"name" : artist.Name,
		"albums" : artist.GetAlbumsIds(), 
	}

	return c.JSON(http.StatusOK, res)
}

func GetArtistCloudWordRPC(c echo.Context) error {
	query := struct {
		Id 		int `json:"id"`
		Size 	int `json:"size"`
	}{}
	if err := c.Bind(&query) ; err != nil {
		log.Println("error binding request, error:", err)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}
	artist, exist := Server.MusicRepo.GetArtist(query.Id)
	if !exist {
		log.Println("error :", fmt.Errorf("track id %d does not exists", query.Id))
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	return c.JSON(http.StatusOK, artist.GetCloudWord(query.Size))
}