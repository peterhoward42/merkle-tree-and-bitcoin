package bitcoin

import (
	"io/ioutil"
	"net/http"
)

func MakeDummyBlockFromSherlockHolmesText() (block Block) {
	res, _ := http.Get("https://sherlock-holm.es/stories/plain-text/bosc.txt")
	body, _ := ioutil.ReadAll(res.Body)
	sz := 512
	numRec := len(body) / sz // deliberate integer division
	for i := 0; i < numRec; i++ {
		record := body[i*sz : (i+1)*sz]
		block.Records = append(block.Records, record)
	}
	return
}
