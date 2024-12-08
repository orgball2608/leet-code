package medium

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for _, s := range word {
		if current.children[s-'a'] == nil {
			current.children[s-'a'] = &WordDictionary{}
		}
		current = current.children[s-'a']
	}
	current.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchHelper(word, 0)
}

func (this *WordDictionary) searchHelper(word string, index int) bool {
	if index == len(word) {
		return this.isEnd
	}

	char := rune(word[index])

	if char == '.' {
		for _, child := range this.children {
			if child != nil && child.searchHelper(word, index+1) {
				return true
			}
		}
		return false
	}

	if this.children[char-'a'] == nil {
		return false
	}

	return this.children[char-'a'].searchHelper(word, index+1)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
