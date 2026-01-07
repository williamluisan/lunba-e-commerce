package gin

import (
	userHandler "github.com/williamluisan/lunba-e-commerce/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
}