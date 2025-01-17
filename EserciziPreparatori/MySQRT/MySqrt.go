package MySQRT

func roundFloat(val float32, precision float32) float32 {
	ratio := float32(1.0)
	for ratio < 1/(10*precision) {
		ratio = ratio * 10
	}
	return float32(int(val*ratio)) / ratio
}

func MyRad2(input int, tolerance float32) float32 {
	start := float32(0)
	end := float32(input)
	solution := end / 2
	app := solution * solution
	for app > float32(input) || app < float32(input)-tolerance {
		app = solution * solution
		if app > float32(input) {
			end = solution
		} else {
			start = solution
		}
		solution = (start + end) / 2
		app = solution * solution
	}
	return roundFloat(solution, tolerance)
}
