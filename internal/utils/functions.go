package utils

func SumInt(arg1, arg2 int) int {
	return arg1 + arg2
}

func SumFloat32(args ...float32) float32 {
	var result float32
	for _, arg := range args {
		result += arg
	}

	return result
}

func SumFloat64(args ...float64) float64 {
	var result float64
	for _, arg := range args {
		result += arg
	}

	return result
}

func FindMax(numbers []int) int {
	var result int

	if len(numbers) == 0 {
		return result
	}

	for _, number := range numbers {
		if number > result {
			result = number
		}
	}

	return result
}

func FilterEvenNumbers(numbers []int) []int {
	var result []int
	for _, v := range numbers {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
}
