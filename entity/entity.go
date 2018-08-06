package entity

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// ID entity id
type ID bson.ObjectId

// String convert an ID in a string
func (i ID) String() string {
	return bson.ObjectId(i).Hex()
}

// StringToID convert a string to an ID
func StringToID(s string) (id ID, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	id = ID(bson.ObjectIdHex(s))
	return
}

// NewID create a new ID
func NewID() (id ID) {
	id, _ = StringToID(bson.NewObjectId().Hex())
	return
}

// MarshalJSON will marshal ID to JSON
func (i ID) MarshalJSON() ([]byte, error) {
	return bson.ObjectId(i).MarshalJSON()
}

// UnmarshalJSON will convert a string to an ID
func (i *ID) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = s[1 : len(s)-1]
	if bson.IsObjectIdHex(s) {
		*i = ID(bson.ObjectIdHex(s))
	}

	return nil
}

// GetBSON implements bson.Getter.
func (i ID) GetBSON() (interface{}, error) {
	if i == "" {
		return "", nil
	}
	return bson.ObjectId(i), nil
}

// SetBSON implements bson.Setter.
func (i *ID) SetBSON(raw bson.Raw) error {
	decoded := new(string)
	bsonErr := raw.Unmarshal(decoded)
	if bsonErr == nil {
		*i = ID(bson.ObjectId(*decoded))
		return nil
	}
	return bsonErr
}
