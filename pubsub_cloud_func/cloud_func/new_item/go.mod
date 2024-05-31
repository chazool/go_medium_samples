module github.com/chazool/go_medium_samples/pubsub_cloud_func/cloud_func/new_item

go 1.22.2

replace github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc => ../../

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.8.1
	github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc v0.0.0-00010101000000-000000000000
	github.com/cloudevents/sdk-go/v2 v2.14.0

)

require (
	cloud.google.com/go/functions v1.15.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
)
