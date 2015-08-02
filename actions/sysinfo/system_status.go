package sysinfo

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func SystemStatus() (string, error) {
	idle0, total0 := getCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks
	ramUsage, err := getRamSample()
	if err != nil {
		return "", fmt.Errorf("An error occurred getting ram status: %s", err.Error())
	}
	output := fmt.Sprintf("CPU usage:\n%f%% [busy: %f, total: %f]\n\nRam usage:\n%s", cpuUsage, totalTicks-idleTicks, totalTicks, ramUsage)
	return output, nil
}

func getRamSample() (string, error) {
	cmd, err := exec.Command("free", "-m").Output()
	return string(cmd), err
}

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}
