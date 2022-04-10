package config

var DefaultConfig *Config

// The default config settings.
//
// Version as given in config/config.go
// FileReadBuffer = 512
func GetConfigDefault() *Config {
	if DefaultConfig == nil {
		DefaultConfig = new(Config)

		DefaultConfig.Paths.Main, _ = GetConfigPathMain()
		DefaultConfig.Paths.Games, _ = GetConfigPathGame()

		DefaultConfig.Main = main{
			Version:        VERSION,
			FileReadBuffer: 512,
		}

		DefaultConfig.Games = make(map[string]string)
	}

	return DefaultConfig
}
