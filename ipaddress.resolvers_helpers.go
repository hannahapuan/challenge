package ipaddress

import (
	"challenge/ent"
	"fmt"
)

func getIPDetailsError(resp []*ent.IPAddress, err error) error {
	if err != nil {
		return err
	} else if len(resp) > 1 {
		return fmt.Errorf("found more than one IP Address record")
	}
	return nil
}
