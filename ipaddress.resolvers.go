package ipaddress

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"challenge/ent"
	"challenge/spamhaus"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *iPAddressResolver) UUID(ctx context.Context, obj *ent.IPAddress) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Enqueue(ctx context.Context, input []string) ([]*ent.IPAddress, error) {
	// Enqueue kicks off a background job to do the DNS lookup and store it in
	// the database for each IP passed in for future lookups
	// If the lookup has already happened, enqueue queues it up again and
	// updates the `response` and `updated_at` fields
	var export []*ent.IPAddress
	var errs error
	for _, ia := range input {
		if !validIPv4(ia) {
			errs = fmt.Errorf("%v: %v", errs, fmt.Errorf("invalid IP provided. Only IPv4 is supported."))
			continue
		}
		// check if ip address exists in table
		resp, err := queryIPRecord(ctx, r.client, ia)
		if err != nil {
			errs = fmt.Errorf("%v: %v", errs, err)
			continue
		}

		// if ip address does not exist, create a new record
		if len(resp) == 0 {
			// query spamhaus for response code of blacklist status
			resp, err := spamhaus.Query(ia)
			if err != nil {
				errs = fmt.Errorf("%w: %w", errs, err)
				continue
			}

			// create record to store ipaddress record information in
			ipar := createIPAddressRecord{
				uuid: uuid.New().String(),
				ip:   ia,
				resp: fmt.Sprintf("%+v\n%+v", resp, err),
			}

			// create IP record
			err = createIPRecord(ctx, r.client, ipar)
			if err != nil {
				errs = fmt.Errorf("%v: %v", errs, err)
				continue
			}

			export = append(export, &ent.IPAddress{
				IPAddress: ipar.ip,
			})
			continue
		}

		// ip address exists in table, update

		// create record to store ipaddress record information in
		ipar := updateIPAddressRecord{
			resp:      fmt.Sprintf("%+v\n%+v", resp, err),
			updatedAt: time.Now(),
		}

		err = updateIPRecord(ctx, r.client, ipar)
		if err != nil {
			errs = fmt.Errorf("%v: %v", errs, err)
			continue
		}

		export = append(export, &ent.IPAddress{
			IPAddress: ia,
		})
	}

	return export, errs
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*ent.IPAddress, error) {
	// GetIPDetails looks up IP Address in the database
	resp, err := queryIPRecord(ctx, r.client, ip)
	if err != nil {
		return nil, err
	} else if len(resp) == 0 {
		return nil, nil
	}
	return resp[0], nil
}

// IPAddress returns IPAddressResolver implementation.
func (r *Resolver) IPAddress() IPAddressResolver { return &iPAddressResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type iPAddressResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
