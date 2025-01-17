package MySQRT

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
	return solution
}
