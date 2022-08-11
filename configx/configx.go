package configx

type ConfigMaskLog struct {
	IsMaskLogWithEncrypt bool
	IsMaskLogWithSymbol  bool
	Env                  string
	Symbol               string
}

func (c *ConfigMaskLog) WithMaskingLogWithEncrypt(env string) {
	c.IsMaskLogWithEncrypt = true
	c.Env = env
}

func (c *ConfigMaskLog) WithMaskingLogWithSymbol(symbol string) {
	c.IsMaskLogWithSymbol = true
	c.Symbol = symbol
}
