package bitcoin

import (
	"io/ioutil"
	"net/http"
)

func MakeBlockBasedOnBookText(storyAbbreviation string) (block Block) {
	fullUrl := "https://sherlock-holm.es/stories/plain-text/" +
		storyAbbreviation + ".txt"
	httpResponse, _ := http.Get(fullUrl)
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
