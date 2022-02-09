package main

type words map[string]struct{}

func sliceToMap(w []string) words {
	set := make(map[string]struct{}, len(w))
	for _, s := range w {
		set[s] = struct{}{}
	}

	return set
}
