package elasticsearch

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestParseQueryResponse(t *testing.T) {

	r := strings.NewReader(`{"took":5,"timed_out":false,"_shards":{"total":5,"successful":5,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":14.263033,"hits":[{"_index":"libraryofcongress_20230921","_id":"sh2004006716","_score":14.263033,"_source":{"id":"sh2004006716","label":"Douglas DC-7 (Transport plane)","source":"lcsh"}}]}}`)

	var query_rsp *QueryResponse

	dec := json.NewDecoder(r)
	err := dec.Decode(&query_rsp)

	if err != nil {
		t.Fatalf("Failed to parse query, %v", err)
	}

	if query_rsp.Hits.Total.Value != 1 {
		t.Fatalf("Unexpected count: %d", query_rsp.Hits.Total.Value)
	}

	if query_rsp.Hits.Results[0].Result.Label != "Douglas DC-7 (Transport plane)" {
		t.Fatalf("Unexpected label: '%s'", query_rsp.Hits.Results[0].Result.Label)
	}
}
