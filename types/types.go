package types

type Site struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
	Config   string `bson:"config" json:"config"`
}
