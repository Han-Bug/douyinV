package rpc

import "time"

var (
	EtcdAddress    = ""
	MuxConnection  = 1
	Timeout        = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond
)

func Init() error {
	if err := initUserRpc(); err != nil {
		return err
	}

	return nil
}
