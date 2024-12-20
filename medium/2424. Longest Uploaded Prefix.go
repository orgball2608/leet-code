package medium

type LUPrefix struct {
	longestPrefix int
	uploaded      []bool
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		longestPrefix: 0,
		uploaded:      make([]bool, n+1),
	}
}

func (this *LUPrefix) Upload(video int) {
	this.uploaded[video] = true

	for this.longestPrefix+1 < len(this.uploaded) && this.uploaded[this.longestPrefix+1] {
		this.longestPrefix++
	}
}

func (this *LUPrefix) Longest() int {
	return this.longestPrefix
}
