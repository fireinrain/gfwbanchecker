package gfwbanchecker

type IP112Api struct {
	ChinaCheckUrl   string `json:"china_check_url"`
	ForeignCheckUrl string `json:"foreign_check_url"`
}

func (api *IP112Api) CheckIP(ip string, port int) CheckerResult {
	return CheckerResult{}
}

func NewIP112Api() *IP112Api {
	return &IP112Api{}
}
