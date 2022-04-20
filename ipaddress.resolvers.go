package ipaddress

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"challenge/ent"
	"challenge/ent/ipaddress"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

func (r *iPAddressResolver) UUID(ctx context.Context, obj *ent.IPAddress) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Enqueue(ctx context.Context, input []string) ([]*ent.IPAddress, error) {
	var export []*ent.IPAddress

	for _, ia := range input {
		_, err := r.client.IPAddress.Create().
			SetUUID(uuid.New().String()).
			SetIPAddress(ia).
			SetResponseCode("400").
			Save(ctx)

		if err != nil {
			return nil, err
		}
		export = append(export, &ent.IPAddress{
			IPAddress: ia,
		})
	}
	return export, nil
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*ent.IPAddress, error) {
	resp, err := r.client.IPAddress.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.In(ipaddress.FieldIPAddress, ip))
		}).
		All(ctx)

	err = getIPDetailsError(resp, err)
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
