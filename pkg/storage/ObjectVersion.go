package storage

type ObjectVersion int64

func NewObjectVersion(value int64) ObjectVersion {
	return ObjectVersion(value)
}

func (version ObjectVersion) Next() ObjectVersion {
	value := int64(version)
	return NewObjectVersion(value + 1)
}
