package db

import "github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"

var (
	DbStore      Store
	Repositories map[string]interfaces.Repository = make(map[string]interfaces.Repository)
)
