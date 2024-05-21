package gfwbanchecker

type GFWBanChecker struct {
	IPCheckerApiPool []CheckerAPI `json:"ip_checker_api"`
}

func NewGFWBanChecker(customCheckerApiPool []CheckerAPI) *GFWBanChecker {
	return &GFWBanChecker{
		IPCheckerApiPool: customCheckerApiPool,
	}
}

func NewDefaultGFWBanChecker() *GFWBanChecker {
	return &GFWBanChecker{}
}

// CheckIP
//
//	@Description: 检测IP是否被ban
//	@receiver c
//	@param ipOrDomain
//	@param port
//	@param inChina
//	@return CheckerResult
func (c *GFWBanChecker) CheckIP(ipOrDomain string, port int, inChina bool) CheckerResult {
	if inChina {
		return c.checkIPInChina(ipOrDomain, port)
	}
	return c.checkIPInForeign(ipOrDomain, port)
}

// CheckIPInChina
//
//	@Description: 在国内检测IP是否被ban
//	@receiver c
//	@param ipOrDomain
//	@param port
//	@return CheckerResult
func (c *GFWBanChecker) checkIPInChina(ipOrDomain string, port int) CheckerResult {
	//TODO 直接ping + tcpconnect 检测
	return CheckerResult{}

}

//
// CheckIPInForeign
//  @Description: 在国外检测IP是否被ban
//  @receiver c
//  @param ipOrDomain
//  @param port
//  @return CheckerResult
//

func (c *GFWBanChecker) checkIPInForeign(ipOrDomain string, port int) CheckerResult {
	//TODO 使用网络接口检测
	return CheckerResult{}
}
