package storage

import "github.com/google/uuid"

type ObjectID uuid.UUID

// Creates a new unique ObjectID.
func NewObjectID() ObjectID {
	guid := uuid.New()

	return ObjectID(guid)
}

// ParseObjectID decodes s into a ObjectID or returns an error.
// It accepts strings in the following formats:
// - xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// - urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// - Microsoft encoding {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
// - raw hex encoding: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func ParseObjectID(value string) (id ObjectID, err error) {
	if guid, err := uuid.Parse(value); err == nil {
		id = ObjectID(guid)
	}

	return
}

// String returns the string form of ObjectID, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if ObjectID is invalid.
func (id ObjectID) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}
