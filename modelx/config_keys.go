package modelx

import (
	"strings"

	"gitdev.devops.krungthai.com/core-services/core-service-lib/logx.git"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	ConfigExpectedConfidenceLevel = "EXPECTED_CONFIDENCE_LEVEL"
	ConfigRequiredTrustedSource   = "REQUIRED_TRUSTED_SOURCE"
	ConfigIsSaveToMainVault       = "IS_SAVE_TO_MAIN_VAULT"
	ConfigMainVaultType           = "MAIN_VAULT_TYPE"
	ConfigTrustedSourceList       = "TRUSTED_SOURCE_LIST"
)

type Config struct {
	Key   string `gorm:"column:key" json:"key"`
	Value string `gorm:"column:value" json:"value"`
}

func (Config) TableName() string {
	return "config"
}

func SetConfigDBToViper(db *gorm.DB) {
	var configs []Config
	if err := db.Find(&configs).Error; err != nil {
		logx.Panic(err)
	}

	for _, config := range configs {
		viper.Set(config.Key, config.Value)
	}
}

func BuildConfigKey(configKeys ...string) string {
	return strings.ToUpper(strings.Join(configKeys, "."))
}
