//ff:func feature=validate type=test-helper control=sequence topic=ssac-sqlc
//ff:what queueInterface — 테스트용 ssacmeta.PackageInterface fixture (queue)

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/ssacmeta"

// queueInterface returns a minimal fixture for the ssac queue package
// mirroring the real canonical_queries declared in ssac/pkg/queue/interface.yaml.
func queueInterface() *ssacmeta.PackageInterface {
	return &ssacmeta.PackageInterface{
		Package: "queue",
		Ports: []ssacmeta.Port{
			{Name: "QueuePublish", UsedBy: []string{"Publish", "PublishTx"}},
			{Name: "QueuePoll", UsedBy: []string{"Subscribe"}},
			{Name: "QueueAck", UsedBy: []string{"Subscribe"}},
		},
	}
}
