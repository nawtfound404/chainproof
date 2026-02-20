package merkel

import (
	"bytes"
	"crypto/sha256"
)

type Tree struct {
	Leaves [][]byte
	Levels [][][]byte
}

type ProofNode struct {
	Hash   []byte
	IsLeft bool
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

func (t *Tree) GenerateProof(index int) []ProofNode {
	var proof []ProofNode
	currentIndex := index

	for level := 0; level < len(t.Levels)-1; level++ {

		nodes := t.Levels[level]

		if currentIndex%2 == 0 {
			//right sibling exists
			if currentIndex+1 < len(nodes) {
				proof = append(proof, ProofNode{
					Hash:   nodes[currentIndex+1],
					IsLeft: false,
				})
			}
		} else {
			proof = append(proof, ProofNode{
				Hash:   nodes[currentIndex-1],
				IsLeft: true,
			})
		}

		currentIndex = currentIndex / 2
	}

	return proof
}

func VerifyProof(leaf []byte, proof []ProofNode, root []byte) bool {

	computed := leaf

	for _, node := range proof {
		if node.IsLeft {
			combined := append(node.Hash, computed...)
			h := sha256.Sum256(combined)
			computed = h[:]
		} else {
			combined := append(computed, node.Hash...)
			h := sha256.Sum256(combined)
			computed = h[:]
		}
	}

	return bytes.Equal(computed, root)
}
