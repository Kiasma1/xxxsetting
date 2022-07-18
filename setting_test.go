package xxxsetting

import (
	"fmt"
	xlog "github.com/Kiasma1/xxxlogger"
	"go.uber.org/zap"
	"os"
	"testing"
)

type IDSettingS struct {
	ID string
}

var IDSetting IDSettingS

func TestMain(m *testing.M) {
	xlog.InitDevLogger("./tmp/dev.log")

	err := os.Setenv("SETTINGS_ID", "10")
	if err != nil {
		xlog.Error("setenv error", zap.Error(err))
		return
	}

	s, err := NewSettingFromEnv()
	if err != nil {
		xlog.Error("new setting from env error", zap.Error(err))
		return
	}

	fmt.Println(s.Get("ID"))

	s2, err := NewSettingFromFile(".")
	if err != nil {
		xlog.Error("new setting from file error", zap.Error(err))
		return
	}

	err = s2.ReadSection("IDSetting", &IDSetting)
	if err != nil {
		return
	}

	fmt.Println(IDSetting)
	return
}
