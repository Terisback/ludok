package ludok

import (
	"time"
)

const (
	dt_max_samples = 60
)

var (
	dt        float64
	dts       [dt_max_samples]float64
	averageDT float64
	lastdt    time.Time = time.Now()

	fps       int
	fpsResult int
	lastFPS   time.Time = time.Now()
)

func calcFrameTime() {
	tmp := time.Since(lastdt)
	lastdt = time.Now()
	dt = float64(tmp.Nanoseconds()) / 1000000

	// Calc average delta time
	averageDT = dt
	for i := 1; i < dt_max_samples; i++ {
		dts[i-1] = dts[i]
		averageDT += dts[i]
	}
	dts[dt_max_samples-1] = dt
	averageDT /= dt_max_samples

	if time.Since(lastFPS).Seconds() >= 1 {
		lastFPS = time.Now()
		fps += 1
		fpsResult = fps
		fps = 0
	} else {
		fps += 1
	}
}

func GetDeltaTime() float64 {
	return dt
}

func GetAvgDeltaTime() float64 {
	return averageDT
}

func GetFPS() int {
	return fpsResult
}
