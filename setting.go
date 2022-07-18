package xxxsetting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSettingFromEnv() (*Setting, error) {
	vp := viper.New()

	// viper config
	vp.SetEnvPrefix("SETTINGS")
	vp.AutomaticEnv()

	s := &Setting{vp}
	return s, nil
}

func NewSettingFromFile(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) Get(key string) interface{} {
	return s.vp.Get(key)
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}