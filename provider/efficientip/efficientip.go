package efficientip

import (
	"context"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type EfficientIPConfig struct {
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	DryRun       bool
	Host         string
	Port         int
	Username     string
	Password     string
	SSlVerify    bool
}

type EfficientIPProvider struct {
	provider.BaseProvider
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
	dryRun       bool
}

func NewEfficientIPProvider(config EfficientIPConfig) (*EfficientIPProvider, error) {
	eipProvider := &EfficientIPProvider{
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
		dryRun:       config.DryRun,
	}
	return eipProvider, nil
}

func (EfficientIPProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	panic("implement me")
	return endpoints, nil
}

func (EfficientIPProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	panic("implement me")

}

func (EfficientIPProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	panic("implement me")
}

func (EfficientIPProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	panic("implement me")
}
