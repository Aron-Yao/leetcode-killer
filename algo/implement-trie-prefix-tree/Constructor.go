const SIZE = 26

type Trie struct {
    end bool
    children [SIZE]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    var now *Trie
    now = this

    for i, w := range word {
        trie := now.children[w-'a']

        if trie == nil {
            t := Constructor()
            now.children[w-'a'] = &t
            now = &t
        } else {
            now = trie
        }
        
        if i == len(word) - 1 {
            now.end = true
        }
    }
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    var now *Trie
    now = this
    
    for _, w := range word {
        if trie := now.children[w-'a']; trie == nil {
            return false
        } else {
            now = trie
        }
    }

    if now.end {
        return true
    } else {
        return false
    }
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    var now *Trie
    now = this
    
    for _, w := range prefix {
        if trie := now.children[w-'a']; trie == nil {
            return false
        } else {
            now = trie
        }
    }

    return true
}
