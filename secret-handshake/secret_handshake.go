package secret

import "math"

var actions = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(code uint) []string {
	var shake []string
	for i := 0.0; i <= 3; i++ {
		u := uint(math.Pow(2, i))
		if code&u == u {
			shake = append(shake, actions[u])
		}
	}
	if code&16 == 16 {
		for i, j := 0, len(shake)-1; i < j; i, j = i+1, j-1 {
			shake[i], shake[j] = shake[j], shake[i]
		}
	}

	return shake
}
