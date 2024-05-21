package gfwbanchecker

type BanedBy int

const (
	BanedChina BanedBy = iota
	BanedForeign
	NoBaned
	AllBaned
)

// 实现 String() 方法
func (b BanedBy) String() string {
	switch b {
	case BanedChina:
		return "BanedChina"
	case BanedForeign:
		return "BanedForeign"
	case NoBaned:
		return "NoBaned"
	case AllBaned:
		return "AllBaned"
	default:
		return "AllBaned"
	}
}

type CheckStatus int

const (
	Open CheckStatus = iota
	Close
)

func (c CheckStatus) String() string {
	switch c {
	case Open:
		return "Open"
	case Close:
		return "Close"
	default:
		return "Close"
	}
}

type ServerStatus int

const (
	Online ServerStatus = iota
	Offline
)

func (s ServerStatus) String() string {
	switch s {
	case Online:
		return "Online"
	case Offline:
		return "Offline"
	default:
		return "Offline"
	}
}

type CheckerResult struct {
	ChinaInfo   IPPortInfo `json:"china_info"`
	ForeignInfo IPPortInfo `json:"foreign_info"`
	BanedByWho  BanedBy    `json:"baned_by_who"`
}

type IPPortInfo struct {
	IP   string      `json:"ip"`
	Port string      `json:"port"`
	TCP  CheckStatus `json:"tcp"`
	ICMP CheckStatus `json:"icmp"`
}

type CheckerAPI interface {
	CheckIP(ip string, port int) CheckerResult
}
