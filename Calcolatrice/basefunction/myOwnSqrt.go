package basefunction

func MyMax[V int | float32 | float64](a V, b V) V {
	if a > b {
		return a
	} else {
		return b
	}
}

func MyMin[V int | float32 | float64](a V, b V) V {
	if a < b {
		return a
	} else {
		return b
	}
}

type calcConstants struct {
	accuracy float64
}

var values = calcConstants{accuracy: 0.001}

func SetConstants(accuracy float64) {
	values.accuracy = accuracy

}

// function to calculate power of base "base" with power n
func MyOwnPwr[V int | float32 | float64, Y int | float32 | float64](base V, n Y) float64 {

	// i would like to have calculation also for rational power but takes too much computation time
	//notRatPart := takeRatPart((float64(base)))

	if base > 0 || base < 0 {
		app := 1.0
		for i := 0; i < int(n); i++ {
			app = app * (float64)(base)
		}
		//if notRatPart != 0 {
		//	app = MyOwnRoot(app, int(notRatPart))
		//}
		return app
	}

	return 0.0
}

// function to calculate root of number "root" with power "power"
func MyOwnRoot[V int | float32 | float64](root V, power int) float64 {
	begin := 0.0
	end := float64(root)
	rad := 0.0
	var currVal float64
	for currVal < float64(root)-values.accuracy || currVal > float64(root)+values.accuracy {
		rad = begin + (end-begin)/2
		currVal = MyOwnPwr(rad, power)
		if currVal > float64(root)+values.accuracy {
			end = rad
		} else {
			begin = rad
		}
	}
	return rad
}

// function to calculate logarithm of number y with power base "base"
func MyOwnLog[V int | float32 | float64](y V, base V) float64 {
	begin := 0.0
	end := float64(y)
	power := 0.0
	var currVal float64
	for currVal < float64(y)-1 || currVal > float64(y)+1 {
		power = begin + (end-begin)/2
		currVal = MyOwnPwr(base, power)
		if currVal > float64(y)+values.accuracy {
			end = power
		} else {
			begin = power
		}
	}
	return power
}
