package ast

import "github.com/byte-wright/lush/internal/common"

type EnvVar struct {
	Name     string
	Location *common.Location
}
