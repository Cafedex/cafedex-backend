package models

import (
	"time"
)

type Guide struct {
	Title     	string    	`json:"brewtitle,omitempty" bson:"brewtitle,omitempty"`
	Author			string			`json:"author,omitempty" bson:"author,omitempty"`
	CreatedAt 	time.Time 	`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Components	[]Component	
}

type Component struct {
	Bloom 		bool					`bson:"bloom"`
	Volume		int64					`bson:"volume"`
	Timing		time.Duration	`bson:"timing"`	
	Note			string 				`bson:"note"`
}