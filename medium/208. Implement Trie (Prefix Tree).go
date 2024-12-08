package medium

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: map[rune]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, s := range word {
		if _, exist := current.children[s]; !exist {
			current.children[s] = &Trie{
				children: map[rune]*Trie{},
			}
		}

		current = current.children[s]
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, s := range word {
		if _, exist := current.children[s]; !exist {
			return false
		}
		current = current.children[s]
	}

	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, s := range prefix {
		if _, exist := current.children[s]; !exist {
			return false
		}
		current = current.children[s]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
