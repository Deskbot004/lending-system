package sql

import (
	"context"
	"fmt"
	"lending-system/ent"
	"log"
)

func AddLending(ctx context.Context, client *ent.Client) (*ent.Lending, error) {
	g, err := client.Lending.
		Create().
		SetDate("today").
		SetNotes("Noted").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating lending: %w", err)
	}
	log.Println("lending was created: ", g)
	return g, nil
}
