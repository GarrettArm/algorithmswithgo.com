package module01

func PrimeCount(n int) map[int]int {
	output := make(map[int]int)
	for i := n; i > 1; i-- {
		localN := int(n)
		for localN%i == 0 {
			if output[i] == 0 {
				output[i] = 1
			} else {
				output[i] += 1
			}
			localN = localN / i
		}
	}
	return output
}

func GCD(a, b int) int {
	var factorsA, factorsB map[int]int
	commons := make(map[int]int)
	factorsA, factorsB = PrimeCount(a), PrimeCount(b)
	for k, _ := range factorsA {
		if factorsB[k] != 0 {
			if factorsA[k] > factorsB[k] {
				commons[k] = factorsB[k]
			} else {
				commons[k] = factorsA[k]
			}
		}
	}
	sum := 1
	for k, v := range commons {
		if k*v > sum {
			sum = k * v
		}
	}
	return sum
}
