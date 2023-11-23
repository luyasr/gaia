package zerolog

type Config struct {
	Mode       string
	Filepath   string `default:"."`
	Filename   string `default:"app.log"`
	MaxSize    int    `default:"10"`
	MaxBackups int    `default:"5"`
	MaxAge     int    `default:"30"`
	Compress   bool   `default:"false"`
}
