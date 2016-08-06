package bitcoin

import (
	"io/ioutil"
	"net/http"
)

// MakeSetOfBlocksBasedOnContentsOfDownloadedBooks is a function that can
// download the text contents of small handfull of Sherlock Holmes books from
// the internet and use the contents of each to create a block. To create
// records inside each block to mimic the presence of transactions in a real
// Blockchain block, it splits the text into 512 byte chunks. There is no
// need for the chunks to be of the same size, but it is harmless if they are, 
// and makes the code simpler.
func MakeSetOfBlocksBasedOnContentsOfDownloadedBooks() (blocks []Block) {
	sherlockHolmesStories := []string{"bosc", "cree", "danc", "gold"}
	for _, urlFragment := range sherlockHolmesStories {
		block := makeBlockBasedOnBookText(urlFragment)
		blocks = append(blocks, block)
	}
	return
}

// makeBlockBasedOnBookText is capable of making a single Block from the text
// of a specified Sherlock Holmes book by downloading the text and splitting at
// 512 byte intervals.
func makeBlockBasedOnBookText(storyAbbreviation string) (block Block) {
	fullUrl := "https://sherlock-holm.es/stories/plain-text/" +
		storyAbbreviation + ".txt"
	httpResponse, _ := http.Get(fullUrl) // Error handling omitted for brevity
	body, _ := ioutil.ReadAll(httpResponse.Body)
	// Split the text into 512 byte records, spilling the remainder, for
	// simplicity.
	recordSize := 512
	numRecords := len(body) / recordSize
	for i := 0; i < numRecords; i++ {
		record := body[i*recordSize : (i+1)*recordSize]
		block.Records = append(block.Records, record)
	}
	return
}
