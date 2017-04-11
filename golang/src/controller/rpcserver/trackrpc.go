package rpcserver

import(
	"net/http"
	"log"
	"fmt"
	"github.com/labstack/echo"
)

const(
	GetTracksIdsRPCName = "getTracksIds"
	GetTrackRPCName = "getTrack"
	GetTrackCloudWordRPCName = "getTrackCloudWord"
)

func GetTracksIdsRPC(c echo.Context) error {
	// Limit and offset to be done on API Part 
	return c.JSON(http.StatusOK, Server.MusicRepo.GetTracksIds())
}

func GetTrackRPC(c echo.Context) error {
	query := struct{
		Id 				int 	`json:"id"`
	}{}
	if err := c.Bind(&query) ; err != nil {
		log.Println("error binding request, error:", err)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}
	
	track, exist := Server.MusicRepo.GetTrack(query.Id)
	if !exist {
		log.Println("error :", fmt.Errorf("track id %d does not exists", query.Id))
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	res := map[string]interface{}{
		"id" : track.Id,
		"name" : track.Name,
		"album" : map[string]interface{}{
			"id" : track.Album.Id,
		},
		"artist" : map[string]interface{}{
			"id" : track.Artist.Id,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func GetTrackCloudWordRPC(c echo.Context) error {
	query := struct {
		Id 		int `json:"id"`
		Size 	int `json:"size"`
	}{}
	if err := c.Bind(&query) ; err != nil {
		log.Println("error binding request, error:", err)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}
	
	track, exist := Server.MusicRepo.GetTrack(query.Id)
	if !exist {
		log.Println("error :", fmt.Errorf("track id %d does not exists", query.Id))
		return c.JSON(http.StatusNotFound, struct{}{})
	}
	if query.Size < 0 {
		log.Println("cannot build cloud of size: %d", query.Size)
		return c.JSON(http.StatusBadRequest, struct{}{})
	}
	
	return c.JSON(http.StatusOK, track.GetCloudWord(query.Size))
}