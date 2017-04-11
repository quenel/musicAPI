package usecases

import (
	"entities/music"
	"testing"
)

type MockMusicSource struct{}

func (source MockMusicSource) GetTrackForAlbum(album *music.Album) ([]*music.Track, error) {
	return []*music.Track{
		&music.Track{album.Artist.Id + 1, "Sultan of Swings", album.Artist, album, []byte{}},
		&music.Track{album.Artist.Id + 2, "Money for noting", album.Artist, album, []byte{}},
		&music.Track{album.Artist.Id + 3, "Romeo & Juliet", album.Artist, album, []byte{}},
		&music.Track{album.Artist.Id + 4, "Lady writer", album.Artist, album, []byte{}},
	}, nil
}

func (source MockMusicSource) GetTopAlbumsForArtist(artist *music.Artist, top int) ([]*music.Album, error) {
	return []*music.Album{
		&music.Album{artist.Id + 1, "Brothers in Arms", artist, []*music.Track{}},
	}, nil
}

func (source MockMusicSource) GetLyricsForTrack(track *music.Track) ([]byte, error) {
	return []byte("Lyrics never matter!"), nil
}

func TestNewMusicRepo(t *testing.T) {
	repo := NewMusicRepo(MockMusicSource{})

	if v, e := repo.Artists.Values[134]; !e || v.Name != "Dire Straits" {
		t.Error("repo shoud contain Dire Straits, having", v)
	}

	if v, e := repo.Albums.Values[135]; !e || v.Name != "Brothers in Arms" {
		t.Error("repo shoud contain Brothers in Arms, having", v)
	}

	if v, e := repo.Tracks.Values[135]; !e || v.Name != "Sultan of Swings" || string(v.Lyrics) != "Lyrics never matter!" {
		t.Error("repo shoud contain Sultan of Swings, having", v)
	}
}