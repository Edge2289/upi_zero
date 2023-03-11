package unionpay

import "upi/expand/unionpay/core"

type Service struct {
	Client *core.Client
	Config Config
}

type Config struct {
	Host string
	Account string
	SysCode string
	PublicKey string
	PrivateKey string
}

func (s *Service) Sign()  {

}
