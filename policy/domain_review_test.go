package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 76, Slack: 25, Drag: 8, Confidence: 48}
	if got := DomainReviewScore(item); got != 201 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "ship" {
		t.Fatalf("lane = %s", got)
	}
}
