package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// InstallmentEntity holds the schema definition for the InstallmentEntity entity.
type InstallmentEntity struct {
	ent.Schema
}

// Fields of the InstallmentEntity.
func (InstallmentEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("customers"),
		field.String("employees"),
		field.String("modelcar"),
		field.String("carregistration"),
		field.Int("numberofperiods").Positive(),
		field.Time("dateofpurchase"),
	}
}

// Edges of the InstallmentEntity.
func (InstallmentEntity) Edges() []ent.Edge {
	return nil
}
