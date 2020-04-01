package telemetry

import (
	"bitbucket.org/bertimus9/systemstat"
	"math"
)

func PollCpu() CpuUsage {
	linuxSample := systemstat.GetCPUSample()
	return CpuUsage{
		Busy:  linuxSample.Nice + linuxSample.User,
		Idle:  linuxSample.Idle,
		Total: linuxSample.Total,
	}
}

func AvgCpuUsage(init, final CpuUsage) (avg float64) {
	linuxInit := systemstat.CPUSample{
		Idle:  init.Idle,
		Total: init.Total,
	}

	linuxFinal := systemstat.CPUSample{
		Idle:  final.Idle,
		Total: final.Total,
	}

	avg = systemstat.GetSimpleCPUAverage(linuxInit, linuxFinal).BusyPct

	if avg < .000001 || math.IsNaN(avg) {
		return 0.0
	}

	return avg
}
