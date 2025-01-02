package datastructures

import (
	"fmt"
	"strings"
)

const indentation = 2

type Dag[T comparable] struct {
	Roots map[T]*DagNode[T]
}

type DagNode[T comparable] struct {
	Value    T
	Children map[T]*DagNode[T]
}

func NewDag[T comparable]() *Dag[T] {
	return &Dag[T]{
		Roots: make(map[T]*DagNode[T]),
	}
}

func (t *Dag[T]) Insert(prev T, next T) {
	prevNode := t.search(prev)
	if prevNode == nil {
		prevNode = &DagNode[T]{
			Value:    prev,
			Children: make(map[T]*DagNode[T]),
		}
		t.Roots[prev] = prevNode
	}

	nextNode, ok := t.Roots[next]
	if ok {
		delete(t.Roots, next)
	} else {
		nextNode = t.search(next)

		if nextNode == nil {
			nextNode = &DagNode[T]{
				Value:    next,
				Children: make(map[T]*DagNode[T]),
			}
		}
	}

	prevNode.Children[next] = nextNode
}

func (t *Dag[T]) search(value T) *DagNode[T] {
	for _, root := range t.Roots {
		if res := root.search(value); res != nil {
			return res
		}
	}

	return nil
}

func (t *DagNode[T]) search(value T) *DagNode[T] {
	if t.Value == value {
		return t
	}

	for _, child := range t.Children {
		if res := child.search(value); res != nil {
			return res
		}
	}

	return nil
}

func (t *Dag[T]) ToString() string {
	nodes := make([]string, len(t.Roots))
	idx := 0
	for _, val := range t.Roots {
		nodes[idx] = val.toString(0)
		idx++
	}

	return strings.Join(nodes, "\n")
}

func (t *DagNode[T]) toString(depth int) string {
	indent := strings.Repeat(" ", depth*indentation)

	if len(t.Children) == 0 {
		return fmt.Sprintf("%v%v", indent, t.Value)
	}

	nodes := make([]string, len(t.Children))
	idx := 0
	for _, child := range t.Children {
		nodes[idx] = child.toString(depth + 1)
		idx++
	}

	return fmt.Sprintf("%v%v: {\n%v\n%v}", indent, t.Value, strings.Join(nodes, "\n"), indent)
}
