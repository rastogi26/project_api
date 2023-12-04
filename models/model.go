package models


// Netflix = public
type Netflix struct {
     ID      string `json:"_id,omitempty" bson:"_id,omitempty"`
	 Title   string `json:"title" bson:"title"`
	 Genre   string `json:"genre" bson:"genre"`

}