package utils

import "github.com/Gokert/gnss-radar/internal/pkg/model"

var AuthorizedUsers = []model.Roles{
	model.RolesAdmin,
	model.RolesUser,
	model.RolesOperator,
}

var OperatorUsers = []model.Roles{
	model.RolesOperator,
	model.RolesAdmin,
}

var AllUsers = []model.Roles{
	model.RolesAdmin,
	model.RolesUser,
	model.RolesOperator,
	model.RolesUnknown,
}
