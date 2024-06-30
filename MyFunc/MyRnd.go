package MyFunc

import (
	"math/rand/v2"
	"time"
)

func GetNumbRnd() int {
	app1 := time.Now()
	app := app1.Nanosecond()
	//fmt.Printf("%d %T", app, app)
	s3 := rand.NewPCG(0, (uint64)(app))
	r3 := rand.New(s3)
	return r3.IntN(100)
}
