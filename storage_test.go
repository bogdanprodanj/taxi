package main

import "testing"

func BenchmarkGetAllRequests(b *testing.B) {
	s := newStorage()
	for i := 0; i < b.N; i++ {
		s.getAllRequests()
	}
}

func BenchmarkMakeRequest(b *testing.B) {
	s := newStorage()
	for i := 0; i < b.N; i++ {
		s.makeRequest()
	}
}
