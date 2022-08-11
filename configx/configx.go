package configx

type ConfigMaskLog struct {
	withMaskLogWithEncrypt bool
	withMaskLogWithSymbol  bool
	Env                    string
	Symbol                 string
}

func (c *ConfigMaskLog) WithMaskingLogWithEncrypted(env string) {
	c.withMaskLogWithEncrypt = true
	c.Env = env
}

func (c *ConfigMaskLog) WithMaskingLogWithSymbol(symbol string) {
	c.withMaskLogWithSymbol = true
	c.Symbol = symbol
}
