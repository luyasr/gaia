package zerolog

type Mode int

const (
	ModeSize Mode = iota
	ModeTime
)

type Config struct {
	Mode         Mode   `default:"1"` // log cutting mode 1: size 2: time
	Filepath     string `default:"."`
	Filename     string `default:"app.log"`
	MaxSize      int    `default:"10"`
	MaxBackups   int    `default:"5"`
	MaxAge       int    `default:"30"`
	Compress     bool   `default:"false"`
	MaxAgeDay    int    `default:"7"`
	RotationTime int    `default:"1"`
}
