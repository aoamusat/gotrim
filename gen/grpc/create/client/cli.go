// Code generated by goa v3.16.1, DO NOT EDIT.
//
// create gRPC client CLI support package
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

package client

import (
	"encoding/json"
	"fmt"

	create "olayml.xyz/gotrim/gen/create"
	createpb "olayml.xyz/gotrim/gen/grpc/create/pb"
)

// BuildCreatePayload builds the payload for the create create endpoint from
// CLI flags.
func BuildCreatePayload(createCreateMessage string) (*create.CreatePayload, error) {
	var err error
	var message createpb.CreateRequest
	{
		if createCreateMessage != "" {
			err = json.Unmarshal([]byte(createCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"long_url\": \"Quia architecto odio odio qui.\"\n   }'")
			}
		}
	}
	v := &create.CreatePayload{
		LongURL: message.LongUrl,
	}

	return v, nil
}