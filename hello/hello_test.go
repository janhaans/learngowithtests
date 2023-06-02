package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Jan", "")
		want := "Hello, Jan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hola to people in Spanish", func(t *testing.T) {
		got := Hello("Jan", "Spanish")
		want := "Hola, Jan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hola to people in French", func(t *testing.T) {
		got := Hello("Jan", "French")
		want := "Bonjour, Jan"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
