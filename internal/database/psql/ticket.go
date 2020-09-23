package psql

import (
	"database/sql"

	"github.com/ofonimefrancis/hexagonal-tickets/internal/ticket"
)

type ticketRepository struct {
	db *sql.DB
}

func NewPostgresTicketRepository(db *sql.DB) ticket.TicketRepository {
	return &ticketRepository{db}
}

func (tr *ticketRepository) CreateTicket(ticket *ticket.Ticket) error {
	tr.db.QueryRow("INSERT INTO tickets(creator, assigned, title, description, status, points, created, updated) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		ticket.Creator, ticket.Assigned, ticket.Title, ticket.Description, ticket.Status, ticket.Points, ticket.Created, ticket.Updated
	).Scan(&ticket.ID)
	return nil 
}
