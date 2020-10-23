package domain

// EffectType ...
type EffectType string

const (
	// ALLOW ...
	ALLOW EffectType = "Allow"

	// DENY ...
	DENY = "Deny"
)

// Resource ...
type Resource struct {
	Name   string     `json:"name,omitempty" bson:"name,omitempty"`
	Effect EffectType `json:"effect,omitempty" bson:"effect,omitempty"`
}
