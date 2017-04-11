package musixmatch

import (
	"entities/music"
	"log"
	"testing"
)

func TestGetAlbum(t *testing.T) {
	source := MusixmatchHandler{}
	albums, err := source.GetTopAlbumsForArtist(&music.Artist{134, "Dire Straits", []*music.Album{}}, 3)
	if err != nil {
		t.Error("could not get top %d albums for artist, error: %s", 3, err)
	}
	log.Println(albums)

	allTracks := make([]*music.Track, 0)
	for _, album := range albums {
		tracks, err := source.GetTrackForAlbum(album)
		if err != nil {
			t.Error("could not get track for album %s, error: %s", album, err)
		}
		allTracks = append(allTracks, tracks ...)
	}
	for _, track := range allTracks {
		log.Println(track.Album.Name, track.Name)
	}
}

func TestDigestLyrics(t *testing.T) {
	source := MusixmatchHandler{}
	track := &music.Track{}
	track.Id = 16697746

	lyrics, err := source.GetLyricsForTrack(track)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(lyrics))
}