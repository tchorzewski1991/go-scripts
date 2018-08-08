// https://www.hackerrank.com/challenges/compare-the-triplets

package main

import "fmt"

// Create 'particitant' struct to make our intention more clear
type participant struct {
	scores []int32
	points int32
}

func main() {
	// Setup participants with random scores and initial value for points
	p1 := &participant{scores: []int32{97, 1, 4},  points: 0}
	p2 := &participant{scores: []int32{92, 23, 2}, points: 0}

	results := compareParticipants(p1, p2)
	fmt.Println(results)
}

// Compare participants and return slice with points for achieved
// by each participant personally.
func compareParticipants(p1 *participant, p2 *participant) []int32 {
	for i := 0; i < 3; i++ {
		if p1.scores[i] > p2.scores[i] { p1.points += 1 }
		if p1.scores[i] < p2.scores[i] { p2.points += 1 }
	}

	return []int32{p1.points, p2.points}
}
