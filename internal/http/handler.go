package http

import (
	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
	repointerfaces "github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

type Handler struct {
	Router     httpinterfaces.Router
	Repository repointerfaces.Repository
}
