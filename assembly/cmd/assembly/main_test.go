package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1,2)
	if result != 3 {
		t.Error(`add error`)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Add(1,2)
	}
}
