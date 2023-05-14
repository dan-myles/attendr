package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ASU_Watchlist holds the schema definition for the ASU_Watchlist entity.
type ASU_Watched_Class struct {
	ent.Schema
}

// Fields of the ASU_Watchlist.
func (ASU_Watched_Class) Fields() []ent.Field {
	return []ent.Field{
		// Regular Class Data
		field.Int("id"),
		field.String("title"),
		field.String("instructor"),
		field.String("subject"),
		field.String("subject_number"),
		field.Bool("has_open_seats").Default(false),
		field.Time("tracked_at").Default(time.Now),

		// Required Data to Query ASU's API
		field.String("class_number").Unique(),
		field.String("term"),
	}
}

// Edges of the ASU_Watchlist.
func (ASU_Watched_Class) Edges() []ent.Edge {
	return nil
}
