package elasticsearch

import (
	"github.com/sfomuseum/go-libraryofcongress-database"
)

/*

{"took":5,"timed_out":false,"_shards":{"total":5,"successful":5,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":14.263033,"hits":[{"_index":"libraryofcongress_20230921","_id":"sh2004006716","_score":14.263033,"_source":{"id":"sh2004006716","label":"Douglas DC-7 (Transport plane)","source":"lcsh"}}]}}
*/

type QueryResponse struct {
	Hits *QueryResponseHits `json:"hits"`
}

type QueryResponseHits struct {
	Total   *QueryResponseTotal    `json:"total"`
	Results []*QueryResponseResult `json:"hits"`
}

type QueryResponseTotal struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type QueryResponseResult struct {
	Index  string                `json:"_index"`
	Id     string                `json:"_id"`
	Score  float64               `json:"_score"`
	Result *database.QueryResult `json:"_source"`
}
