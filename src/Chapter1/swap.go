package main

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
