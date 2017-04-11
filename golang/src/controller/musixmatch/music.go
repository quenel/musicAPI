package musixmatch

import (
	"encoding/json"
	"entities/music"
	"errors"
	"fmt"
	"strconv"
)

const (
	artistAlbumsPath = "artist.albums.get"
	albumTrackPath   = "album.tracks.get"
	trackLyricsPath  = "track.lyrics.get"
)

const (
	artistIdParam         = "artist_id"
	artistRatingParam     = "s_artist_rating"
	albumIdParam          = "album_id"
	groupByAlbumNameParam = "g_album_name"
	resultSizeParam       = "page_size"
	resultPageParam       = "page"
	lyricsFilterParam     = "f_has_lyrics"
	trackIdParam          = "track_id"
)

const (
	acsendant  = "asc"
	descendant = "desc"
)

type MusixmatchHandler struct {}

func (m MusixmatchHandler) GetLyricsForTrack(track *music.Track) ([]byte, error) {
	parameters := make(map[string]string)
	parameters[trackIdParam] = strconv.Itoa(track.Id)

	req, err := buildRequest(trackLyricsPath, parameters)
	if err != nil {
		return []byte{}, fmt.Errorf("could not get track %s, lyrics, error: %s", track, err)
	}

	trackLyrics := struct {
		Lyrics struct {
			Body json.RawMessage `json:"lyrics_body"`
		} `json:"lyrics"`
	}{}

	err = get(req, &trackLyrics)
	if err != nil {
		return []byte{}, fmt.Errorf("could not get tracks %s, lyrics, error: %s", track, err)
	}

	return trackLyrics.Lyrics.Body, nil
}

func (m MusixmatchHandler) GetTrackForAlbum(album *music.Album) ([]*music.Track, error) {
	paramters := make(map[string]string)
	paramters[albumIdParam] = strconv.Itoa(album.Id)
	paramters[resultSizeParam] = "100"
	paramters[resultPageParam] = "1"
	paramters[lyricsFilterParam] = "true"

	req, err := buildRequest(albumTrackPath, paramters)
	if err != nil {
		return []*music.Track{}, fmt.Errorf("could not get album %s tracks, error: %s", album, err)
	}

	trackList := struct {
		Tracks []struct {
			Track struct {
				Id   int    `json:"track_id"`
				Name string `json:"track_name"`
			} `json:"track"`
		} `json:"track_list"`
	}{}

	err = get(req, &trackList)
	if err != nil {
		return []*music.Track{}, fmt.Errorf("could not get album: %s, track list, error: %s", album, err)
	}

	tracks := make([]*music.Track, len(trackList.Tracks))
	for i, wrapper := range trackList.Tracks {
		tracks[i] = &music.Track{wrapper.Track.Id, wrapper.Track.Name, album.Artist, album, []byte{}}
	}

	return tracks, nil
}

func (m MusixmatchHandler) GetTopAlbumsForArtist(artist *music.Artist, top int) ([]*music.Album, error) {
	if top > 10 {
		return []*music.Album{}, errors.New("cannot get more that top 10")
	}

	paramters := make(map[string]string)
	paramters[artistIdParam] = strconv.Itoa(artist.Id)
	paramters[artistRatingParam] = descendant
	paramters[groupByAlbumNameParam] = "1"
	paramters[resultSizeParam] = strconv.Itoa(top)

	req, err := buildRequest(artistAlbumsPath, paramters)
	if err != nil {
		return []*music.Album{}, fmt.Errorf("could not get artist: %d, error: %s", artist.Id, err)
	}

	albumList := struct {
		Albums []struct {
			Album struct {
				Id   int    `json:"album_id"`
				Name string `json:"album_name"`
			} `json:"album"`
		} `json:"album_list"`
	}{}
	err = get(req, &albumList)
	if err != nil {
		return []*music.Album{}, fmt.Errorf("could not get artist: %d, album list, error: %s", artist.Id, err)
	}

	albums := make([]*music.Album, top)
	for i, wrapper := range albumList.Albums {
		albums[i] = &music.Album{wrapper.Album.Id, wrapper.Album.Name, artist, []*music.Track{}}
	}

	return albums, nil
}
