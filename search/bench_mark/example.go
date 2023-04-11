package benchmark

func linearSearch(a, b []int) {
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				return
			}
		}
	}
}

func linearSearchWithHashMap(a, b []int) {
	hashmap := make(map[int]struct{}, len(b))
	for _, v := range b {
		hashmap[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := hashmap[v]; ok {
			return
		}
	}
}
