package tree

type TrieTree struct {
	Root *TrieNode
}

type TrieNode struct {
	Val      string
	IsEnd    bool
	Children map[string]*TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		Root: &TrieNode{
			Children: make(map[string]*TrieNode),
		},
	}
}

func (tree *TrieTree) Insert(word string) {
	cur := tree.Root

	for i := 0; i < len(word); i++ {
		if node, exist := cur.Children[word[i:i+1]]; exist {
			cur = node
		} else {
			cur.Children[word[i:i+1]] = &TrieNode{
				Val:      word[i : i+1],
				IsEnd:    false,
				Children: make(map[string]*TrieNode),
			}
			cur = cur.Children[word[i:i+1]]
		}
	}

	cur.IsEnd = true
}

func (tree *TrieTree) Search(word string) bool {
	cur := tree.Root

	for i := 0; i < len(word); i++ {
		node, exist := cur.Children[word[i:i+1]]
		if !exist {
			return false
		}

		cur = node
	}

	return cur.IsEnd
}
