package server

import "github.com/nicolasparada/monorepo-demo/types"

func (svc *Service) User(id string) types.User {
	return types.User{ID: id}
}
