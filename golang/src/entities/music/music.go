package music

import(
	"entities/cloudword"
)

type Artist struct {
	Id 		int 
	Name 	string	
	Albums 	[]*Album
}

type Album struct {
	Id 		int	
	Name 	string
	Artist  *Artist
	Tracks  []*Track
}

type Track struct {
	Id 		int	
	Name 	string
	Artist 	*Artist
	Album 	*Album
	Lyrics  []byte
}

func (t *Track) GetCloudWord(n int) cloudword.CloudWord {
	wc := cloudword.NewWordCounter(t.Lyrics)
	return wc.BuildCloud(n)
}

func (a *Album) GetTracksIds() []int {
	ids := make([]int, len(a.Tracks))
	for i, track := range a.Tracks {
		ids[i] = track.Id
	}
	return ids
}

func (a *Album) getWordCounter() cloudword.WordCounter {
	wc := cloudword.WordCounter{}
	for _, track := range a.Tracks {
		wc = cloudword.Merge(wc, cloudword.NewWordCounter(track.Lyrics))
	}
	return wc
}
 
func (a *Album) GetCloudWord(n int) cloudword.CloudWord {
	return a.getWordCounter().BuildCloud(n)
}

func (a *Artist) GetCloudWord(n int) cloudword.CloudWord {
	wc := cloudword.WordCounter{}
	for _, album := range a.Albums {
		wc = cloudword.Merge(wc, album.getWordCounter())
	}
	return wc.BuildCloud(n)
}

func (a *Artist) GetAlbumsIds() []int {
	ids := make([]int, len(a.Albums))
	for i, album := range a.Albums {
		ids[i] = album.Id
	}
	return ids
}