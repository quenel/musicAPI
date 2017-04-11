package rpcserver

import(
	"usecases"
	"github.com/labstack/echo"
)

const (
	// GlobalPath is the path we use for calling RPCs
	RPCPath = "/rpc/"
)

var (
	Server *RPCServer
)

type RPCServer struct{
	MusicRepo *usecases.MusicRepo
}

func init() {
	Server = &RPCServer{}
	e := echo.New()
	
	e.POST(RPCPath + GetArtistsIdsRPCName, GetArtistsIdsRPC)
	e.POST(RPCPath + GetArtistRPCName, GetArtistRPC)
	e.POST(RPCPath + GetArtistCloudWordRPCName, GetArtistCloudWordRPC)
	e.POST(RPCPath + GetAlbumsIdsRPCName, GetAlbumsIdsRPC)
	e.POST(RPCPath + GetAlbumRPCName, GetAlbumRPC)
	e.POST(RPCPath + GetAlbumCloudWordRPCName, GetAlbumCloudWordRPC)
	e.POST(RPCPath + GetTracksIdsRPCName, GetTracksIdsRPC)
	e.POST(RPCPath + GetTrackRPCName, GetTrackRPC)
	e.POST(RPCPath + GetTrackCloudWordRPCName, GetTrackCloudWordRPC)

	go e.Start(":8080")
}

