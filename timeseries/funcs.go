package timeseries

type F func(Time, float32, float32) float32

func Any(t Time, v1, v2 float32) float32 {
	if !IsNaN(v1) {
		return v1
	}
	return v2
}

func NanSum(t Time, sum, v float32) float32 {
	if IsNaN(sum) {
		sum = 0
	}
	if !IsNaN(v) {
		sum += v
	}
	return sum
}

func Max(t Time, max, v float32) float32 {
	if IsNaN(max) {
		return v
	}
	if IsNaN(v) {
		return max
	}
	if v > max {
		return v
	}
	return max
}

func Min(t Time, min, v float32) float32 {
	if IsNaN(min) {
		return v
	}
	if IsNaN(v) {
		return min
	}
	if v < min {
		return v
	}
	return min
}

func Defined(t Time, v float32) float32 {
	if IsNaN(v) {
		return 0
	}
	return 1
}

func NanToZero(t Time, v float32) float32 {
	if IsNaN(v) {
		return 0
	}
	return v
}
