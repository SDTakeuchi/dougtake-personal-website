package model

type Tag interface {
	ID() uint64
	Name() string
}
