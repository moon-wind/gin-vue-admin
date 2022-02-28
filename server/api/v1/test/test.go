package test

import (
	"github.com/gin-gonic/gin"
	"net"
)

type TestApi struct{}

// @Tags Test
// @Summary 测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test/t1 [get]
func (e *TestApi) Test1(c *gin.Context) {
	ip, err := getLocalIP()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    0,
			"message": ip,
		})
	}
}

func getLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
