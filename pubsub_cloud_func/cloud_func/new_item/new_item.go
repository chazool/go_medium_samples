package new_item

import (
	"context"
	"fmt"
	"log"

	"github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc/app/dto"
	"github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc/app/service"
	"github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc/pkg"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("newItem", NewItem)
}

type MessagePublishedData struct {
	Message PubSubMessage
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func NewItem(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData

	// Convert event data to MessagePublishedData struct
	if err := e.DataAs(&msg); err != nil {
		log.Printf("error converting event data: %v", err)
		return fmt.Errorf("event.DataAs: %w", err)
	}

	// Build an Item struct from the message data
	i, err := pkg.StructBuilder[dto.Item](msg.Message.Data)
	if err != nil {
		log.Printf("error building item structure: %v", err)
		return fmt.Errorf("struct builder: %w", err)
	}

	// Attempt to add the new item using the service
	srv := service.NewItemService()
	if err := srv.New(i); err != nil {
		log.Printf("error occurred while adding new item: %v", err)
		return fmt.Errorf("adding item: %w", err)
	}

	return nil
}
