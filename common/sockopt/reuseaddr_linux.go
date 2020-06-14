package sockopt

import (
	"net"
	"syscall"

	"github.com/Dreamacro/clash/log"
)

func UDPReuseaddr(c *net.UDPConn) {
	rc, err := c.SyscallConn()
	if err != nil {
		log.Warnln("Failed to get Raw UDP Connection: %s", err)
		return
	}

	rc.Control(func(fd uintptr) {
		err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
		if err != nil {
			log.Warnln("Failed to Reuse UDP Address: %s", err)
			return
		}
	})
}
