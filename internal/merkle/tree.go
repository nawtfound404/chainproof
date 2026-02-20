package merkel

import "crypto/sha256"

type Tree struct {
	Leaves [][]byte
	Levels [][][]byte
}

func NewTree(hashes [][]byte) *Tree {
	tree := &Tree{
		Leaves: hashes,
	}
	tree.build()
	return tree
}

func (t *Tree) build() {
	current := t.Leaves
	t.Levels = append(t.Levels, current)

	for len(current) > 1 {
		var next [][]byte
		for i := 0; i < len(current); i += 2 {
			if i+1 == len(current) {
				next = append(next, current[i])
				continue
			}

			combined := append(current[i], current[i+1]...)
			hash := sha256.Sum256(combined)
			next = append(next, hash[:])
		}
		current = next
		t.Levels = append(t.Levels, current)

	}
}

func (t *Tree) Root() []byte {
	lastlevel := t.Levels[len(t.Levels)-1]
	return lastlevel[0]
}
