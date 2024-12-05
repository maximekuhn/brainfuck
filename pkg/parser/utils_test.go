package parser

import (
	"reflect"
	"testing"
)

func TestNodeStackEmpty(t *testing.T) {
	ns := newNodeStack()
	n, found := ns.pop()
	if found {
		t.Fatalf("pop(): expected found to be false")
	}
	if n != nil {
		t.Fatalf("pop(): expected to find nothing but found a non-nil node")
	}
}

func TestNodeStackPushAndPop(t *testing.T) {
	ns := newNodeStack()
	n1 := &Node{Type: NodeIncrement, Child: nil}
	n2 := &Node{Type: NodeDecrement, Child: nil}
	n3 := &Node{Type: NodeNext, Child: nil}
	n4 := &Node{Type: NodeOutput, Child: nil}

	ns.push(n1)
	ns.push(n2)
	ns.push(n3)
	ns.push(n4)

	if len(ns.stack) != 4 {
		t.Fatalf("expected ns.stack to have len 4 but got %d", len(ns.stack))
	}

	n4a, found := ns.pop()
	if !found {
		t.Fatalf("pop(): expected to find n4 got nothing")
	}
	if !reflect.DeepEqual(n4, n4a) {
		t.Fatalf("pop(): expected to find %v got %v", n4, n4a)
	}
	n3a, found := ns.pop()
	if !found {
		t.Fatalf("pop(): expected to find n3 got nothing")
	}
	if !reflect.DeepEqual(n3, n3a) {
		t.Fatalf("pop(): expected to find %v got %v", n3, n3a)
	}
	n2a, found := ns.pop()
	if !found {
		t.Fatalf("pop(): expected to find n2 got nothing")
	}
	if !reflect.DeepEqual(n2, n2a) {
		t.Fatalf("pop(): expected to find %v got %v", n2, n2a)
	}
	n1a, found := ns.pop()
	if !found {
		t.Fatalf("pop(): expected to find n1 got nothing")
	}
	if !reflect.DeepEqual(n1, n1a) {
		t.Fatalf("pop(): expected to find %v got %v", n1, n1a)
	}
}

func TestNodeStackPeek(t *testing.T) {
	ns := newNodeStack()
	n1 := &Node{Type: NodeIncrement, Child: nil}
	ns.push(n1)

	n1a, found := ns.peek()
	if !found {
		t.Fatalf("peek(): expected to find %v got nothing", n1)
	}
	if !reflect.DeepEqual(n1, n1a) {
		t.Fatalf("peek(): expected to find %v got %v", n1, n1a)
	}

	_, _ = ns.pop()
	_, found = ns.peek()
	if found {
		t.Fatal("peek(): expected to find nothing")
	}

}
