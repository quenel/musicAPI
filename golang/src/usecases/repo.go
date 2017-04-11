package usecases

import (
	"entities/music"
	"log"
	"sync"
)

const (
	MaxWorker = 5
)

var TargetArtists = []music.Artist{
	{134, "Dire Straits", []*music.Album{}},
	{49, "Radiohead", []*music.Album{}},
	{426, "Eminem", []*music.Album{}},
}

type MusicRepo struct {
	Artists ArtistsRepo
	Albums  AlbumsRepo
	Tracks  TracksRepo
}

type ArtistsRepo struct {
	mu *sync.RWMutex
	Values map[int]*music.Artist
	Ids []int
}

type AlbumsRepo struct {
	mu *sync.RWMutex
	Values map[int]*music.Album
	Ids []int
}

type TracksRepo struct {
	mu *sync.RWMutex
	Values map[int]*music.Track
	Ids []int
}

type MusicSource interface {
	GetTrackForAlbum(album *music.Album) ([]*music.Track, error)
	GetTopAlbumsForArtist(artist *music.Artist, top int) ([]*music.Album, error)
	GetLyricsForTrack(track *music.Track) ([]byte, error)
}

func NewMusicRepo(source MusicSource) *MusicRepo {
	repo := &MusicRepo{
		Artists: ArtistsRepo{
			new (sync.RWMutex),
			make(map[int]*music.Artist),
			make([]int, 0),
		},
		Albums: AlbumsRepo{
			new (sync.RWMutex),
			make(map[int]*music.Album),
			make([]int, 0),
		},
		Tracks: TracksRepo{
			new (sync.RWMutex),
			make(map[int]*music.Track),
			make([]int, 0),
		},
	}

	digestArtist := func(r *MusicRepo, source MusicSource, artist *music.Artist, output chan interface{}, wg *sync.WaitGroup) {
		albums, err := source.GetTopAlbumsForArtist(artist, 3)
		if err != nil {
			log.Fatal("Could not get albums for artist", artist, ",error:", err)
			return
		}
		artist.Albums = albums
		r.Artists.mu.Lock()
		r.Artists.Values[artist.Id] = artist
		r.Artists.Ids = append(r.Artists.Ids, artist.Id)
		r.Artists.mu.Unlock()

		for _, album := range albums {
			output <- album
		}
		wg.Add(len(albums) - 1)
	}

	digestAlbums := func(r *MusicRepo, source MusicSource, album *music.Album, output chan interface{}, wg *sync.WaitGroup) {
		tracks, err := source.GetTrackForAlbum(album)
		if err != nil {
			log.Fatal("Could not get albums for artist", album, ",error:", err)
			return
		}
		album.Tracks = tracks
		r.Artists.mu.Lock()
		r.Albums.Values[album.Id] = album
		r.Albums.Ids = append(r.Albums.Ids, album.Id)
		r.Artists.mu.Unlock()

		for _, track := range tracks {
			output <- track
		}
		wg.Add(len(tracks) - 1)
	}

	digestTrack := func(r *MusicRepo, source MusicSource, track *music.Track, wg *sync.WaitGroup) {
		lyrics, err := source.GetLyricsForTrack(track)
		if err != nil {
			log.Fatal("could not get lyrics for track", track, "error:", err)
		}
		track.Lyrics = lyrics
		r.Tracks.mu.Lock()
		r.Tracks.Values[track.Id] = track
		r.Tracks.Ids = append(r.Tracks.Ids, track.Id)
		r.Tracks.mu.Unlock()

		wg.Done()
	}

	worker := func(jobs chan interface{}, wg *sync.WaitGroup) {
		for job := range jobs {
			switch input := job.(type) {
			case music.Artist:
				digestArtist(repo, source, &input, jobs, wg)
			case *music.Album:
				digestAlbums(repo, source, input, jobs, wg)
			case *music.Track:
				digestTrack(repo, source, input, wg)
			}
		}
	}

	jobs, wg := make(chan interface{}, 1000), &sync.WaitGroup{}
	for i := 0; i < MaxWorker; i++ {
		go worker(jobs, wg)
	}
	for _, artist := range TargetArtists {
		wg.Add(1)
		jobs <- artist
	}

	wg.Wait()
	close(jobs)

	return repo
}

func (r *MusicRepo) GetArtist(id int) (*music.Artist, bool) {
	artist, exist := r.Artists.Values[id]
	if !exist {
		return nil, false
	}
	return artist, true
}

func (r *MusicRepo) GetArtistsIds() []int {
	return r.Artists.Ids
} 

func (r *MusicRepo) GetAlbum(id int) (*music.Album, bool) {
	album, exist := r.Albums.Values[id]
	if !exist {
		return nil, false
	}
	return album, true
}

func (r *MusicRepo) GetAlbumsIds() []int {
	return r.Albums.Ids
}

func (r *MusicRepo) GetTrack(id int) (*music.Track, bool) {
	track, exist := r.Tracks.Values[id]
	if !exist {
		return nil, false
	}
	return track, true
}

func (r *MusicRepo) GetTracksIds() []int {
	return r.Tracks.Ids
} 
