package config

type (
	Config struct {
		//	application configurations
		AppPort    string `mapstructure:"app_port"`
		AppMode    string `mapstructure:"app_mode"`
		BaseUrlApp string `mapstructure:"base_url_app"`

		// mongodb
		Mongo struct {
			URI      string `mapastructure:"uri"`
			DBName   string `mapstructure:"dbname"`
			UserColl string `mapstructure:"user_coll"`
			LinkColl string `mapstructure:"link_coll"`
		} `mapstructure:"mongodb"`
	}
)
