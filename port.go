package ngrpc

import "fmt"

func GetListenPort(v uint16) string {
	if v > 0 {
		return fmt.Sprintf("0.0.0.0:%d", v)
	}
	return ":0"
}
