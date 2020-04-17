package config

var (
	providerServiceModels map[string]*ProviderServiceModel
	consumerServiceModels map[string]*ConsumerServiceModel
)

type ProviderServiceModel struct {
	serviceKey string
	*ServiceConfig
}

type ConsumerServiceModel struct {
	serviceKey string
	*ReferenceConfig
}

func init(){
	providerServiceModels = make(map[string]*ProviderServiceModel)
	consumerServiceModels = make(map[string]*ConsumerServiceModel)
}

func registryProvider(key string, value *ProviderServiceModel){
	providerServiceModels[key] = value
}

func registryConsumer(key string, value *ConsumerServiceModel){
	consumerServiceModels[key] = value
}