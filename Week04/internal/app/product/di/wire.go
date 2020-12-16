//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kwstars/Go-000/Week04/internal/app/product/controllers"
	"github.com/kwstars/Go-000/Week04/internal/app/product/dao"
	"github.com/kwstars/Go-000/Week04/internal/app/product/services"
	"github.com/kwstars/Go-000/Week04/internal/pkg/database"
	"github.com/kwstars/Go-000/Week04/internal/pkg/transports/http"
)

var providerSet = wire.NewSet(
	database.ProviderSet,
	dao.ProvierSets,
	services.ProviderSet,
	controllers.ProviderSet,
	http.ProviderSet,
	NewApp,
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(providerSet))
}
