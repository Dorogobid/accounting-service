package usecases

import (
	"context"
	"database/sql"
	"log"

	"github.com/Dorogobid/accounting-service/internal/database"
)

type UseCases struct {
	ctx context.Context
	queries *database.Queries
}

func New(ctx context.Context, db *sql.DB) *UseCases{
	queries := database.New(db)
	uc := UseCases{ctx: ctx, queries: queries}
	return &uc
}

func (u *UseCases) GetClientTypes() ([]database.ClientType, error) {
	clientTypes, err := u.queries.GetClientTypes(u.ctx)
	if err != nil {
		log.Printf("Error getting client_types: %s\n", err)
		return nil, err
	}
	return clientTypes, nil
}

func (u *UseCases) GetClientsWithType() ([]database.GetClientsWithTypeRow, error) {
	clientTypes, err := u.queries.GetClientsWithType(u.ctx)
	if err != nil {
		log.Printf("Error getting clients with type_name: %s\n", err)
		return nil, err
	}
	return clientTypes, nil
}