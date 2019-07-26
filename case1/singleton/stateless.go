package singleton

import (
	"fmt"
	"time"
)

// This file is in another package.

// StatelessReport class manages report for the destination.
type StatelessReport struct {
}

// SendTo send the report to the input destination.
func (r *StatelessReport) SendTo(dest string) {
	// Time Consuming work here.
	time.Sleep(time.Second)
	fmt.Printf("The report has been sent to: %s\n", dest)
}

// ReportSingleton contains StatelessReport to send the report.
type ReportSingleton struct {
	singleton *StatelessReport
}

// GetInstanceStatelessReport get the singleton instance from this variable.
func (rs *ReportSingleton) GetInstanceStatelessReport() *StatelessReport {
	if rs.singleton == nil {
		rs.singleton = new(StatelessReport)
	}
	return rs.singleton
}
