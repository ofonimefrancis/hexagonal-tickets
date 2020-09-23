package ticket

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	OPEN = "OPEN"
)

type TicketService interface {
	CreateTicket(ticket *Ticket) error
	FindTicketById(id string) (*Ticket, error)
	FindAllTickets() ([]*Ticket, error)
}

type ticketService struct {
	repo TicketRepository
}

func NewTicketService(repo TicketRepository) TicketService {
	return &ticketService{repo}
}

func (ts *ticketService) CreateTicket(ticket *Ticket) error {
	ticket.ID = uuid.New().String()
	ticket.Created = time.Now()
	ticket.Updated = time.Now()
	ticket.Status = OPEN

	if err := ts.repo.CreateTicket(ticket); err != nil {
		logrus.WithField("error", err).Error("Error creating ticket")
		return err
	}
	logrus.WithField("id", ticket.ID).Info("Created new ticket")
	return nil
}

func (ts *ticketService) FindTicketById(id string) (*Ticket, error) {
	ticket, err := ts.repo.FindTicketById(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error finding ticket")
		return ticket, err
	}
	logrus.WithField("id", id).Info("Found ticket")
	return ticket, nil
}

func (ts *ticketService) FindAllTickets() ([]*Ticket, error) {
	tickets, err := ts.repo.FindAllTickets()

	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error finding all tickets")
		return tickets, err
	}
	logrus.Info("Found all tickets")
	return tickets, nil
}
