package gin

import "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler"

type Dependencies struct {
	UserHandler *handler.UserHandler
}