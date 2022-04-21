package ipaddress

import (
	"challenge/ent"
	"challenge/ent/ipaddress"
	"context"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// stores values used to create the IPAddressRecord
type createIPAddressRecord struct {
	uuid string
	ip   string
	resp string
}

// stores values used to update the IPAddressRecord
type updateIPAddressRecord struct {
	resp      string
	updatedAt time.Time
}

// parses errors from querying IPDetails
func getIPDetailsError(resp []*ent.IPAddress, err error) error {
	if err != nil {
		return err
	} else if len(resp) > 1 {
		return fmt.Errorf("found more than one IP Address record")
	}
	return nil
}

// checks if the given IP exists in the store
func queryIPRecord(ctx context.Context, client *ent.Client, ip string) ([]*ent.IPAddress, error) {
	resp, err := client.IPAddress.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.In(ipaddress.FieldIPAddress, ip))
		}).
		All(ctx)
	err = getIPDetailsError(resp, err)
	return resp, err
}

// creates an IP record in the store
func createIPRecord(ctx context.Context, client *ent.Client, ipar createIPAddressRecord) error {
	_, err := client.IPAddress.Create().
		SetUUID(ipar.uuid).
		SetIPAddress(ipar.ip).
		SetResponseCode(ipar.resp).
		Save(ctx)

	return err
}

// updates an IP record in the store
func updateIPRecord(ctx context.Context, client *ent.Client, ipar updateIPAddressRecord) error {
	_, err := client.IPAddress.Create().
		SetResponseCode(ipar.resp).
		SetUpdatedAt(ipar.updatedAt).
		Save(ctx)

	return err
}

// checks if the string provided is a valid ipv4 value
func validIPv4(ip string) bool {
	return strings.Count(ip, ":") < 2
}
