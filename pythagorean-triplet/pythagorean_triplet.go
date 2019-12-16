package pythagorean

import "math"

// Triplet represents a pythagorean triplet
type Triplet [3]int

//Range returns a list of all Pythagorean triplets with sides in the
//range min to max inclusive.
func Range(min, max int) []Triplet {
	var trips []Triplet
	for i := min; i <= max-2; i++ {
		for j := i + 1; j <= max-1; j++ {
			r := math.Sqrt(float64(i*i + j*j))
			_, f := math.Modf(r)
			if f == 0 {
				if int(r) <= max {
					trips = append(trips, Triplet{i, j, int(r)})
				}
			}
		}
	}
	return trips
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	var out []Triplet
	trips := Range(1, p)
	for _, t := range trips {
		if t[0]+t[1]+t[2] == p {
			out = append(out, t)
		}
	}
	return out
}
