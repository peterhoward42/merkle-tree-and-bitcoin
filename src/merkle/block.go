package merkle

import ()

/*
	"net/http"
	"io/ioutil"
*/

type Block struct {
	records []Record
}

type Record []byte

/*
func (blk *Block) fillRecords() {
	res, _ := http.Get("https://sherlock-holm.es/stories/plain-text/bosc.txt")
	body, _ := ioutil.ReadAll(res.Body)
	sz := 512
	numRec := len(body) / sz // deliberate integer division
	for i := 0; i < numRec; i++ {
		record := body[i * sz: (i + 1) * sz]
		blk.records = append(blk.records, record)
	}
}
*/
