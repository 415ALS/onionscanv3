package protocol

import (
	"fmt"
	"github.com/415ALS/onionscanv3/config"
	"github.com/415ALS/onionscanv3/report"
	"github.com/415ALS/onionscanv3/utils"
)

type VNCProtocolScanner struct {
}

func (vncps *VNCProtocolScanner) ScanProtocol(hiddenService string, osc *config.OnionScanConfig, report *report.OnionScanReport) {
	// MongoDB
	osc.LogInfo(fmt.Sprintf("Checking %s VNC(5900)\n", hiddenService))
	conn, err := utils.GetNetworkConnection(hiddenService, 5900, osc.TorProxyAddress, osc.Timeout)
	if err != nil {
		osc.LogInfo("Failed to connect to service on port 5900\n")
		report.VNCDetected = false
	} else {
		osc.LogInfo("Detected possible VNC instance\n")
		// TODO: Actual Analysis
		report.VNCDetected = true
	}
	if conn != nil {
		conn.Close()
	}
}
