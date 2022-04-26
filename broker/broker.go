package broker

import (
	"context"
	"fmt"

	"code.cloudfoundry.org/lager"
	brokerapi "github.com/pivotal-cf/brokerapi/domain"
)

type SchedulerBroker struct {
	Logger    lager.Logger
	Config    Config
	Instances map[string]brokerapi.GetInstanceDetailsSpec
	Bindings  map[string]brokerapi.GetBindingSpec
}

type Config struct {
	ServiceName    string
	ServicePlan    string
	BaseGUID       string
	Credentials    interface{}
	Tags           string
	ImageURL       string
	SysLogDrainURL string
	Free           bool
}

func NewSchedulerImpl(logger lager.Logger) (broker *SchedulerBroker) {
	var credentials interface{}

	fmt.Printf("Credentials: %v\n", credentials)

	return &SchedulerBroker{
		Logger:    logger,
		Instances: map[string]brokerapi.GetInstanceDetailsSpec{},
		Bindings:  map[string]brokerapi.GetBindingSpec{},
		Config: Config{
			BaseGUID:    "29140B3F-0E69-4C7E-8A35",
			ServiceName: "OCF Scheduler",
			ServicePlan: "basic",
			Credentials: credentials,
			Tags:        "scheduler",
			ImageURL:    "https://example.com/image",
			Free:        true,
		},
	}
}

// Services gets the catalog of services offered by the service broker
//   GET /v2/catalog
func (broker *SchedulerBroker) Services(ctx context.Context) ([]brokerapi.Service, error) {
	planList := []brokerapi.ServicePlan{
		{
			ID:          broker.Config.BaseGUID,
			Name:        broker.Config.ServicePlan,
			Description: broker.Config.ServiceName,
			Metadata: &brokerapi.ServicePlanMetadata{
				DisplayName: "basic",
			},
		}}
	return []brokerapi.Service{
		{
			ID:          "ocf-scheduler-ID",
			Name:        "ocf-scheduler",
			Description: "OCF Scheduler for CRON Jobs",
			Bindable:    true,
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName: broker.Config.ServiceName,
				ImageUrl:    broker.Config.ImageURL,
			},
			Plans:         planList,
			PlanUpdatable: true,
		}}, nil
}

// Provision creates a new service instance
//   PUT /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, fmt.Errorf("unimplemented")
}

// Deprovision deletes an existing service instance
//  DELETE /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, fmt.Errorf("unimplemented")
}

// GetInstance fetches information about a service instance
//   GET /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) GetInstance(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {
	return brokerapi.GetInstanceDetailsSpec{}, fmt.Errorf("unimplemented")
}

// Update modifies an existing service instance
//  PATCH /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, fmt.Errorf("unimplemented")
}

// LastOperation fetches last operation state for a service instance
//   GET /v2/service_instances/{instance_id}/last_operation
func (broker *SchedulerBroker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, fmt.Errorf("unimplemented")
}

// Bind creates a new service binding
//   PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}
func (broker *SchedulerBroker) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding,
	error) {
	return brokerapi.Binding{}, fmt.Errorf("unimplemented")
}

// Unbind deletes an existing service binding
//   DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}
func (broker *SchedulerBroker) Unbind(ctx context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {
	return brokerapi.UnbindSpec{}, fmt.Errorf("unimplemented")
}

// GetBinding fetches an existing service binding
//   GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}
func (broker *SchedulerBroker) GetBinding(ctx context.Context, instanceID string, bindingID string) (brokerapi.GetBindingSpec, error) {
	return brokerapi.GetBindingSpec{}, fmt.Errorf("unimplemented")
}

// LastBindingOperation fetches last operation state for a service binding
//   GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation
func (broker *SchedulerBroker) LastBindingOperation(ctx context.Context, instanceID string, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, fmt.Errorf("unimplemented")
}
