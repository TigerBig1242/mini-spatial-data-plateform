package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type SpatialFeature struct {
	ID           uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	VallarisId   string          `gorm:"index:idx_vallaris_id,unique;not null" json:"vallaris_id"`
	CollectionId string          `gorm:"index" json:"collection_id"`
	Geometry     string          `gorm:"type:geometry(Geometry,4326);index:,type:gist" json:"geometry"`
	Properties   json.RawMessage `gorm:"type:jsonb" json:"properties"`
}
