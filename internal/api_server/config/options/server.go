package options

import "github.com/gin-gonic/gin"

type RunOptions struct {
	Mode        string   `json:"mode"`
	Healthz     bool     `json:"healthz"`
	Middlewares []string `json:"middlewares"`
}

func NewRunOptions() *RunOptions {
	return &RunOptions{
		Mode:        gin.ReleaseMode,
		Healthz:     true,
		Middlewares: []string{},
	}
}
