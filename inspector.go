package main

type Inspector struct {
	Base           `bson:",inline" json:",inline"`
	Username       string `bson:"username,omitempty" json:"username,omitempty"`
	Password       string `bson:"password,omitempty" json:"password,omitempty"`
	Name           string `bson:"name,omitempty" json:"name,omitempty"`
	ProfilePicture string `bson:"profile_picture,omitempty" json:"profile_picture,omitempty"`
	CreatedBy      string `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy      string `bson:"updated_by,omitempty" json:"updated_by"`

	LastLogin     *int64 `bson:"last_login,omitempty"`
	PartnershipId string `bson:"partnership_id,omitempty" json:"partnership_id,omitempty"`
}
