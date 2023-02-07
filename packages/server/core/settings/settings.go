package settings

import "github.com/PureML-Inc/PureML/server/tools/security"

// Settings defines common app configuration options.
type Settings struct {
	S3 S3Config `form:"s3" json:"s3"`

	AdminAuthToken TokenConfig `form:"adminAuthToken" json:"adminAuthToken"`
}

// New creates and returns a new default Settings instance.
func New() *Settings {
	return &Settings{
		AdminAuthToken: TokenConfig{
			Secret:   security.RandomString(50),
			Duration: 1209600, // 14 days,
		},
	}
}

type TokenConfig struct {
	Secret   string `form:"secret" json:"secret"`
	Duration int64  `form:"duration" json:"duration"`
}

type S3Config struct {
	Enabled        bool   `form:"enabled" json:"enabled"`
	Bucket         string `form:"bucket" json:"bucket"`
	Region         string `form:"region" json:"region"`
	Endpoint       string `form:"endpoint" json:"endpoint"`
	AccessKey      string `form:"accessKey" json:"accessKey"`
	Secret         string `form:"secret" json:"secret"`
	ForcePathStyle bool   `form:"forcePathStyle" json:"forcePathStyle"`
}
