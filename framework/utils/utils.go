package utils

import(
	"github.com/dilmorja/backero"
)

func ToTarget[Tar *backero.Target|any](extended Tar) *backero.Target {
	response := backero.Target(extended)
	return &response
}