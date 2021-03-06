package ipfs

import (
	"errors"
	"github.com/ipfs/go-ipfs/commands"
	coreCmds "github.com/ipfs/go-ipfs/core/commands"
)

var pubErr = errors.New(`Name publish failed`)

// Publish a signed IPNS record to our Peer ID
func Publish(ctx commands.Context, hash string) (string, error) {
	args := []string{"name", "publish", hash}
	req, cmd, err := NewRequest(ctx, args)
	if err != nil {
		return "", err
	}
	res := commands.NewResponse(req)
	cmd.Run(req, res)
	resp := res.Output()
	returnedVal := resp.(*coreCmds.IpnsEntry).Value
	if res.Error() != nil {
		return "", res.Error()
	}
	if returnedVal != hash {
		return "", pubErr
	}
	log.Infof("Published %s to IPNS", hash)
	return returnedVal, nil
}
