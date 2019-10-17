package config

import (
	"fmt"
	"graphql-rbac/pkg/config"

	"github.com/spf13/viper"
)

const (
	authPostgresURI      = "AUTH_POSTGRES_URI"
	authMigrationVersion = "AUTH_MIGRATION_VERSION"

	gatewayAddr                = "GATEWAY_ADDR"
	gatewayEndpoint            = "GATEWAY_ENDPOINT"
	gatewayServePlayground     = "GATEWAY_SERVE_PLAYGROUND"
	gatewayPlaygroundEndpoint  = "GATEWAY_PLAYGROUND_ENDPOINT"
	gatewayEnableIntrospection = "GATEWAY_ENABLE_INTROSPECTION"

	seedUserLogin    = "SUPER_USER_LOGIN"
	seedUserPassword = "SUPER_USER_LOGIN"
	seedRoleTitle    = "SEED_ROLE_TITLE"
	seedRoleSuper    = "SEED_ROLE_SUPER"

	sessionAccessTokenTTL  = "ACCESS_TOKEN_TTL"
	sessionRefreshTokenTTl = "REFRESH_TOKEN_TTL"
)

// NewDefaults returns map with default service params
func NewDefaults() map[string]interface{} {
	defaults := make(map[string]interface{})

	defaults[authPostgresURI] = "postgresql://postgres:postgres@localhost:5432/test?sslmode=disable"
	defaults[authMigrationVersion] = 0

	defaults[gatewayAddr] = ":10000"
	defaults[gatewayEndpoint] = "/graphql"
	defaults[gatewayServePlayground] = true
	defaults[gatewayPlaygroundEndpoint] = "/playground"
	defaults[gatewayEnableIntrospection] = true

	defaults[seedUserLogin] = "root"
	defaults[seedUserPassword] = "root"
	defaults[seedRoleTitle] = "ROOT"
	defaults[seedRoleSuper] = true

	defaults[sessionAccessTokenTTL] = 1000000
	defaults[sessionRefreshTokenTTl] = 5000000

	return defaults
}

// Config contains all parameters for starting the service
type Config struct {
	DB      DBConfig
	Gateway GatewayConfig
	Session SessionConfig
	Seed    SeedConfig
}

// NewConfig returns initialized service config
func NewConfig(defaults map[string]interface{}) (Config, error) {
	v, err := config.BindEnv(viper.New(), defaults, true, nil)
	if err != nil {
		return Config{}, fmt.Errorf("reading env error: %w", err)
	}

	cfg := Config{
		DB: DBConfig{
			PostgresURI:      v.GetString(authPostgresURI),
			MigrationVersion: v.GetUint(authMigrationVersion),
		},
		Gateway: GatewayConfig{
			Addr:                v.GetString(gatewayAddr),
			Endpoint:            v.GetString(gatewayEndpoint),
			ServePlayground:     v.GetBool(gatewayServePlayground),
			PlaygroundEndpoint:  v.GetString(gatewayPlaygroundEndpoint),
			EnableIntrospection: v.GetBool(gatewayEnableIntrospection),
		},
		Session: SessionConfig{
			AccessTokenTTL:  v.GetInt(sessionAccessTokenTTL),
			RefreshTokenTTL: v.GetInt(sessionRefreshTokenTTl),
		},
		Seed: SeedConfig{
			SeedUserLogin:    v.GetString(seedUserLogin),
			SeedUserPassword: v.GetString(seedUserPassword),
			SeedRoleTitle:    v.GetString(seedRoleTitle),
			SeedRoleSuper:    v.GetBool(seedRoleSuper),
		},
	}

	return cfg, nil
}

// DBConfig
type DBConfig struct {
	PostgresURI      string
	MigrationVersion uint
}

// GatewayConfig
type GatewayConfig struct {
	Addr                string
	Endpoint            string
	ServePlayground     bool
	PlaygroundEndpoint  string
	EnableIntrospection bool
}

// SessionConfig
type SessionConfig struct {
	AccessTokenTTL  int
	RefreshTokenTTL int
}

// SeedConfig
type SeedConfig struct {
	SeedUserLogin    string
	SeedUserPassword string
	SeedRoleTitle    string
	SeedRoleSuper    bool
}
