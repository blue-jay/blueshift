package boot

import (
	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/env"
	"github.com/blue-jay/blueshift/repository"
	"github.com/blue-jay/blueshift/usecase"
)

// Service represents all the services that the application uses.
type Service struct {
	RouterService domain.RouterCase
	UserService   domain.UserCase
	ViewService   domain.ViewCase
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices(config *env.Info) *Service {
	s := new(Service)

	RegisterServicesLegacy(config)

	s.RouterService = usecase.NewRouterCase(new(repository.RouterRepo))
	s.ViewService = usecase.NewViewCase()

	// Initialize the clients.
	/*db := repository.NewClient("db.json")

	// Store all the services for the application.
	s.UserService = usecase.NewUserCase(
		repository.NewUserRepo(db),
		new(passhash.Item))
	s.ViewService = view.New("../../view", "tmpl")*/

	return s
}
