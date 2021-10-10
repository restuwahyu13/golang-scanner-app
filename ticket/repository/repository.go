package repository

import (
	"pborebuild/ticket/entity"
	"pborebuild/ticket/schema"
)

type Ticket interface {
	BiayaTicket(input *schema.SchemaTicket) *entity.EntityTicket
}

type BiayaTicket struct {
	ticket Ticket
}

func (r *BiayaTicket) BiayaTicket(input *schema.SchemaTicket) *entity.EntityTicket {
	var ticket entity.EntityTicket

	if ticket.TicketCount >= 50 {
		ticket.Pesawat = input.Pesawat
		ticket.Harga = input.Harga
		ticket.Discount = input.Discount
		ticket.Total = (ticket.Harga - ticket.Discount) * ticket.TicketCount
	} else {
		ticket.Pesawat = input.Pesawat
		ticket.Harga = input.Harga
		ticket.Discount = input.Discount
		ticket.Total = input.Harga * input.TicketCount
	}

	return &ticket
}

func NewTicket(input *schema.SchemaTicket) *entity.EntityTicket {
	var ticket BiayaTicket
	return ticket.BiayaTicket(input)
}
