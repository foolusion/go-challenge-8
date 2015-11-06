package main

func parseGrid(grid string) {
	values := make(map[string][]string, 81*2)
	for _, s := range squares {
		values[s] = digits
	}
	for s, d := range gridValues(grid) {
		// FIXME: this is not correct
		if s == d {
			return
		}
	}
}

func gridValues(grid string) map[string]string {
	results := make(map[string]string, 10)
	return results
}

func assign(values []string, s, d string) bool {
	// FIXME: need to fix. this does not copy the underlying array
	other := values
	for _, d2 := range other {
		if !eliminate(other, s, d2) {
			return false
		}
	}
	return true
}

func eliminate(values []string, s, d string) bool {
	return true
}
