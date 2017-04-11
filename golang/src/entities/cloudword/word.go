package cloudword

import (
	"sort"
	"bufio"
	"bytes"
	"regexp"
	"strings"
)


type WordOccurence struct {
	Word 		string 	`json:"word"`
	Occurence 	int 	`json:"occurence"`		
}

type WordCounter map[string]int

type CloudWord []WordOccurence

var	isAlpha = regexp.MustCompile(`(^[A-Za-z]+$)`).MatchString

func NewWordCounter(b []byte) WordCounter {
	wc := WordCounter{}

	b = bytes.Replace(b, []byte(`\n`), []byte(" "), -1)
	buff := bytes.NewBuffer(b)
	scanner := bufio.NewScanner(buff)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// only words of more that 5 runes
		if len(scanner.Text()) < 5 || !isAlpha(scanner.Text()) {
			continue
		}
		lowerText := strings.ToLower(scanner.Text())
		if count, exist := wc[lowerText] ; exist {
			wc[lowerText] = count + 1
			continue
		}
		wc[lowerText] = 1
	}

	return wc
}

func Merge(wc1, wc2 WordCounter) WordCounter {
	res := make(map[string]int)
	for word, counter1 := range wc1 {
		if counter2, exist := wc2[word] ; exist {
			res[word] = counter1 + counter2 
			continue
		}
		res[word] = counter1
	}
	for word, counter2 := range wc2 {
		if _, exist := res[word] ; !exist {
			res[word] = counter2
		}
	}

	return res
}

func (wc WordCounter) BuildCloud(n int) CloudWord {
	if n < 0 {
		return CloudWord{}
	}
	cloud := CloudWord(make([]WordOccurence, 0))
	for word, occurence := range wc {
		cloud = append(cloud, WordOccurence{word, occurence})
	} 
	sort.Sort(cloud)
	if n > len(wc) {
		return cloud
	}
	return cloud[0 : n]
}

func (cloud CloudWord) Len() int {
	return len(cloud)
}

func (cloud CloudWord) Swap(i, j int) {
	cloud[i], cloud[j] = cloud[j], cloud[i]
} 

func (cloud CloudWord) Less(i, j int) bool {
	return cloud[i].Occurence > cloud[j].Occurence
}