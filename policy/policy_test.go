package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 62, Capacity: 86, Latency: 13, Risk: 21, Weight: 9}
	if got := Score(signal); got != 51 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 89, Capacity: 89, Latency: 12, Risk: 14, Weight: 12}
	if got := Score(signal); got != 169 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 68, Capacity: 98, Latency: 8, Risk: 13, Weight: 8}
	if got := Score(signal); got != 143 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
