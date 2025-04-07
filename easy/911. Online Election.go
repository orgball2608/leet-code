package easy

type TopVotedCandidate struct {
	leaders []int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	votes := make(map[int]int)
	leaders := make([]int, len(times))
	maxVote, curLeader := 0, -1
	for i, p := range persons {
		votes[p]++
		if votes[p] >= maxVote {
			maxVote = votes[p]
			curLeader = p
		}
		leaders[i] = curLeader
	}
	return TopVotedCandidate{
		leaders: leaders,
		times:   times,
	}
}

// Time: O(N)
// Space: O(N)

func (this *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(this.times), func(i int) bool { return this.times[i] > t }) - 1
	return this.leaders[i]
}

// Time: O(log N)
// Space: O(1)

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
