package gfwbanchecker

type GFWBanChecker struct{}

func NewGFWBanChecker() *GFWBanChecker {
	return &GFWBanChecker{}
}

// CheckIP
//
//	@Description: 检测IP是否被ban
//	@receiver c
//	@param ip
//	@param port
//	@param inChina
//	@return CheckerResult
func (c *GFWBanChecker) CheckIP(ip string, port int, inChina bool) CheckerResult {
	if inChina {
		return c.checkIPInChina(ip, port)
	}
	return c.checkIPInForeign(ip, port)
}

// CheckIPInChina
//
//	@Description: 在国内检测IP是否被ban
//	@receiver c
//	@param ip
//	@param port
//	@return CheckerResult
func (c *GFWBanChecker) checkIPInChina(ip string, port int) CheckerResult {
	//TODO 直接ping + tcpconnect 检测
	return CheckerResult{}

}

//
// CheckIPInForeign
//  @Description: 在国外检测IP是否被ban
//  @receiver c
//  @param ip
//  @param port
//  @return CheckerResult
//

func (c *GFWBanChecker) checkIPInForeign(ip string, port int) CheckerResult {
	//TODO 使用网络接口检测
	return CheckerResult{}
}
