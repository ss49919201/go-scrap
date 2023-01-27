package main

func merge(target []int) []int {
	half := len(target) / 2
	if half == 0 {
		return target
	}
	a := merge(target[:half])
	b := merge(target[half:])
	sorted := make([]int, 0, len(target))
	for {
		if len(a) == 0 && len(b) == 0 {
			break
		}

		if len(a) == 0 {
			sorted = append(sorted, b[0])
			b = b[1:]
		} else if len(b) == 0 {
			sorted = append(sorted, a[0])
			a = a[1:]
		} else {
			if a[0] > b[0] {
				sorted = append(sorted, b[0])
				b = b[1:]
			} else {
				sorted = append(sorted, a[0])
				a = a[1:]
			}
		}
	}
	return sorted
}
