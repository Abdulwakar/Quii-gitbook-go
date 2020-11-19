package main

import "testing"

func TestHello(t *testing.T)  {
	assertCorrectMessage := func(t *testing.T,got,want string){
		t.Helper()
		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying Hello to People", func(t *testing.T) {
		got := Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t,got,want)
	})

	t.Run("Say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Abdul","Spanish")
		want := "Halo, Abdul"
		assertCorrectMessage(t,got,want)
	})

}
