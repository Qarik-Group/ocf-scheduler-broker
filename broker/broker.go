package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

func envOr(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func NewSchedulerImpl(logger lager.Logger) (broker *SchedulerBroker) {
	credentials := map[string]interface{}{}
	credentials["api_endpoint"] = envOr("SCHEDULER_URL", "http://localhost:8000")

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
	var parameters interface{}
	json.Unmarshal(details.RawParameters, &parameters)
	broker.Instances[instanceID] = brokerapi.GetInstanceDetailsSpec{
		ServiceID:  details.ServiceID,
		PlanID:     details.PlanID,
		Parameters: parameters,
	}
	fmt.Println(broker.Instances[instanceID])
	spec := brokerapi.ProvisionedServiceSpec{}
	return spec, nil
}

// Deprovision deletes an existing service instance
//  DELETE /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{
		IsAsync:       true,
		OperationData: "I have no idea what operation data is",
	}, nil
}

// GetInstance fetches information about a service instance
//   GET /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) GetInstance(ctx context.Context, instanceID string) (spec brokerapi.GetInstanceDetailsSpec, err error) {
	if val, ok := broker.Instances[instanceID]; ok {
		fmt.Println("Found and returned!")
		return val, nil
	}
	return brokerapi.GetInstanceDetailsSpec{}, fmt.Errorf("InstanceID %s not found", instanceID)
}

// Update modifies an existing service instance
//  PATCH /v2/service_instances/{instance_id}
func (broker *SchedulerBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {

	return brokerapi.UpdateServiceSpec{
		IsAsync: true,
	}, nil
}

// LastOperation fetches last operation state for a service instance
//   GET /v2/service_instances/{instance_id}/last_operation
func (broker *SchedulerBroker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{
		State: "success",
	}, nil
}

// Bind creates a new service binding
//   PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}
func (broker *SchedulerBroker) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding,
	error) {
	var parameters interface{}
	broker.Bindings[bindingID] = brokerapi.GetBindingSpec{
		Credentials: broker.Config.Credentials,
		Parameters:  parameters,
	}
	return brokerapi.Binding{
		Credentials: broker.Config.Credentials,
	}, nil
}

// Unbind deletes an existing service binding
//   DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}
func (broker *SchedulerBroker) Unbind(ctx context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {
	return brokerapi.UnbindSpec{}, nil
}

// GetBinding fetches an existing service binding
//   GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}
func (broker *SchedulerBroker) GetBinding(ctx context.Context, instanceID string, bindingID string) (brokerapi.GetBindingSpec, error) {
	if val, ok := broker.Bindings[bindingID]; ok {
		return val, nil
	}
	return brokerapi.GetBindingSpec{}, fmt.Errorf("BindingID %s not found", bindingID)
}

// LastBindingOperation fetches last operation state for a service binding
//   GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation
func (broker *SchedulerBroker) LastBindingOperation(ctx context.Context, instanceID string, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{
		State: "success",
	}, nil
}
