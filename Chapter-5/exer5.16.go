package main

func Join(seperator string, strings...string)string {
	var summedString string
	for _, s := range strings {
		summedString += s + seperator
	}
	return summedString
}