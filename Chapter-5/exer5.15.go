package main

func max(vals...int)int{
	if len(vals) == 0 {
		return 0
	}
	n := vals[0]
	for _, m := range vals {
		if m > n {
			n = m
		}
	}
	return n
}

func min(vals...int)int{
	if len(vals) == 0 {
		return 0
	}
	n := vals[0]
	for _, m := range vals {
		if m < n {
			n = m
		}
	}
	return n
}

func maxOne(one int, vals...int)int{
	if len(vals) == 0 {
		return one
	}
	vals = append(vals, one)
	n := vals[0]
	for _, m := range vals {
		if m > n {
			n = m
		}
	}
	return n
}

func minOne(one int, vals...int)int{
	if len(vals) == 0 {
		return one
	}
	vals = append(vals, one)
	n := vals[0]
	for _, m := range vals {
		if m < n {
			n = m
		}
	}
	return n
}