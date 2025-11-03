package main

import "testing"

func Example_main() {
	main()
	// Output: Hello world
}

func TestGreet(t *testing.T) {
	expected := "Hello world"

	greeting := greet()

	if greeting != expected {
		t.Errorf("Greeting = %q; want %q", greeting, expected)
	}
	//le %q permet d'afficher les guillemets dans le message d'erreur le %x affiche la valeur en hexadécimal
}
