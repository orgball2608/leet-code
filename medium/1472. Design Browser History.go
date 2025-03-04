package medium

type Tab struct {
	Url  string
	Next *Tab
	Prev *Tab
}

type BrowserHistory struct {
	Curr *Tab
}

func Constructor(homepage string) BrowserHistory {
	home := &Tab{
		Url: homepage,
	}
	return BrowserHistory{
		Curr: home,
	}
}

func (this *BrowserHistory) Visit(url string) {
	newTab := &Tab{
		Url:  url,
		Prev: this.Curr,
		Next: nil,
	}

	this.Curr.Next = newTab
	this.Curr = newTab
}

// Time: O(1)
// Space: O(1)

func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && this.Curr.Prev != nil {
		this.Curr = this.Curr.Prev
		steps--
	}
	return this.Curr.Url
}

// Time: O(min(steps, N))
// Space: O(1)

func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && this.Curr.Next != nil {
		this.Curr = this.Curr.Next
		steps--
	}
	return this.Curr.Url
}

// Time: O(min(steps, N))
// Space: O(1)

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

// Time: O(N)
// Space: O(
