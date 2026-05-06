package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 62, Capacity: 86, Latency: 13, Risk: 21, Weight: 9}, wantScore: 51, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 89, Capacity: 89, Latency: 12, Risk: 14, Weight: 12}, wantScore: 169, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 68, Capacity: 98, Latency: 8, Risk: 13, Weight: 8}, wantScore: 143, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
