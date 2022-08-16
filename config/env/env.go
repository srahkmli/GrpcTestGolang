package env

type EnvMode byte

const (
	UNDEFINED EnvMode = iota
	PRODUCTION
	STAGE
	DEVELOPMENT
)
