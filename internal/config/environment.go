package config

type EnvironmentConfig string

var (
	EnvDev  = EnvironmentConfig("dev")
	EnvProd = EnvironmentConfig("prod")
)

var environmentMap = map[string]EnvironmentConfig{
	"dev":  EnvDev,
	"prod": EnvProd,
}

func checkEnvironment(environment string) (env EnvironmentConfig, ok bool) {
	env, ok = environmentMap[environment]
	return
}
