package config

var DefaultConfig *Config

func GetConfigDefault() *Config {
	if DefaultConfig == nil {
		DefaultConfig = new(Config)

		DefaultConfig.Paths.Main, _ = GetConfigPathMain()
		DefaultConfig.Paths.Games, _ = GetConfigPathGame()

		DefaultConfig.Main.Version = VERSION
		DefaultConfig.Main.FileReadBuffer = 512
	}

	return DefaultConfig
}
