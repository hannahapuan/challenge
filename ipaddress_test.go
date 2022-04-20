package ipaddress

import (
	"challenge/ent"
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func IPAddressCreate_Test() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	rec1, err := client.IPAddress.Create().
		SetUUID(uuid.New()).
		SetIPAddress("123.345.346").
		SetResponseCode("200").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a rec: %v", err)
	}
	fmt.Printf("%d: %q\n", rec1.ID, rec1.IPAddress)
	rec2, err := client.IPAddress.Create().
		SetUUID(uuid.New()).
		SetIPAddress("222.222.222").
		SetResponseCode("404").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a rec: %v", err)
	}
	fmt.Printf("%d: %q\n", rec2.ID, rec2.IPAddress)
	// Output:
	// 1: "123.345.346"
	// 2: "222.222.222"
}
