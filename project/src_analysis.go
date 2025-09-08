package project

// Using for calculate the src
// key is IP/Host
// value is the IP/Host hitted times
var srcount map[string]int

func CalculateSrc(srcs []string) map[string]int {
	// initialize Src
	srcount = make(map[string]int)

	// read src
	for _, data := range srcs {
		if _, exists := srcount[data]; exists {
			srcount[data]++
		} else {
			srcount[data] = 1
		}
	}
	return srcount
}
