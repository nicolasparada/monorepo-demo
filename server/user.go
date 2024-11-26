package server

import "github.com/nicolasparada/monorepo-demo/types"

func (svc *Service) User(id, name string) types.User {
	return types.User{ID: id, Name: name}
}
