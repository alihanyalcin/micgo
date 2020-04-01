package telemetry

import (
	"context"
	"project/internal/pkg/bootstrap/container"
	"project/internal/pkg/bootstrap/startup"
	"project/internal/pkg/di"
	"runtime"
	"sync"
	"time"
)

type SystemUsage struct {
	Memory     memoryUsage
	CpuBusyAvg float64
}

type memoryUsage struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects uint64
}

type CpuUsage struct {
	Busy, // time used by all processes. this ideally does not include system processes.
	Idle, // time used by the idle process
	Total uint64
}

var lastSample CpuUsage
var usageAvg float64

func NewSystemUsage() (s SystemUsage) {
	var rtm runtime.MemStats

	// read full memory stats
	runtime.ReadMemStats(&rtm)

	// Memory stats
	s.Memory.Alloc = rtm.Alloc
	s.Memory.TotalAlloc = rtm.TotalAlloc
	s.Memory.Sys = rtm.Sys
	s.Memory.Mallocs = rtm.Mallocs
	s.Memory.Frees = rtm.Frees

	// Live objects = Mallocs - Frees
	s.Memory.LiveObjects = s.Memory.Mallocs - s.Memory.Frees

	s.CpuBusyAvg = usageAvg

	return s
}

func cpuUsageAverage() {
	nextUsage := PollCpu()
	usageAvg = AvgCpuUsage(lastSample, nextUsage)
	lastSample = nextUsage
}

func BootstrapHandler(wg *sync.WaitGroup, ctx context.Context, _ startup.Timer, dic *di.Container) bool {
	loggingClient := container.LoggingClientFrom(dic.Get)
	loggingClient.Info("Telemetry starting.")

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			cpuUsageAverage()

			for seconds := 30; seconds > 0; seconds-- {
				select {
				case <-ctx.Done():
					loggingClient.Info("Telemetry stopped.")
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}
	}()

	return true
}
