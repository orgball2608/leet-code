package medium

import "container/heap"

type Tweet struct {
	ID        int
	Timestamp int
}

type TweetHeap []Tweet

func (h TweetHeap) Len() int {
	return len(h)
}

func (h TweetHeap) Less(i, j int) bool {
	return h[i].Timestamp < h[j].Timestamp
}

func (h TweetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Twitter struct {
	userTweets  map[int][]Tweet
	userFollows map[int]map[int]bool
	timestamp   int
}

func Constructor() Twitter {
	return Twitter{
		userTweets:  make(map[int][]Tweet),
		userFollows: make(map[int]map[int]bool),
		timestamp:   0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, exists := this.userTweets[userId]; !exists {
		this.userTweets[userId] = []Tweet{}
	}

	this.userTweets[userId] = append(this.userTweets[userId], Tweet{ID: tweetId, Timestamp: this.timestamp})
	this.timestamp++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	heap.Init(h)

	if tweets, exists := this.userTweets[userId]; exists {
		for _, tweet := range tweets {
			heap.Push(h, tweet)
			if h.Len() > 10 {
				heap.Pop(h)
			}
		}
	}

	if follows, exists := this.userFollows[userId]; exists {
		for followedId := range follows {
			if tweets, exists := this.userTweets[followedId]; exists {
				for _, tweet := range tweets {
					heap.Push(h, tweet)
					if h.Len() > 10 {
						heap.Pop(h)
					}
				}
			}
		}
	}

	result := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Tweet).ID
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := this.userFollows[followerId]; !exists {
		this.userFollows[followerId] = make(map[int]bool)
	}

	this.userFollows[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if follows, exists := this.userFollows[followerId]; exists {
		delete(follows, followeeId)
	}
}
