// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/external"
)

func resolveClientConfig(svc *Client) func(configs []interface{}) error {
	return func(configs []interface{}) error {
		if value, ok, err := external.ResolveUseARNRegion(configs); err != nil {
			return err
		} else if ok {
			svc.UseARNRegion = value
		}

		return nil
	}
}
