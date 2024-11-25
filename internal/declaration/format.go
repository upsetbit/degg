package declaration

type (
	Format int
)

const (
	JSON Format = iota
	YAML
	TOML
	UNKNOWN
)
