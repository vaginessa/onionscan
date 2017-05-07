package protocol

import (
	"github.com/Minyall/onionscan/config"
	"github.com/Minyall/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
