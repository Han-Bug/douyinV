package rpc

func Init() error {
	err := initUserRpc()
	if err != nil {
		return err
	}
	return nil
}
