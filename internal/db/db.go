package db

import (
	"github.com/jmoiron/sqlx"
)

type Db struct {
	client *sqlx.DB
}

// TODO move to calling code
type Storage interface {
	GetBalance(int) (int, error)
	CreateInvoice(int, int)
	CreateWithdraw(int, int)
	Clear()
}

func NewDb(client *sqlx.DB) Storage {
	return &Db{
		client: client,
	}
}

func (d *Db) GetBalance(clientId int) (res int, err error) {
	err = d.client.Get(&res, `
		select (select coalesce(sum(amount), 0) from public.invoice where client_id=$1) - (select coalesce(sum(amount), 0)
		from public.withdraw
		where client_id=$1)
	`, clientId)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (d *Db) CreateInvoice(amount int, clientID int) {
	_, err := d.client.Exec(`insert into public.invoice (amount, client_id) values ($1, $2);`, amount, clientID)
	if err != nil {
		panic(err)
	}
}

// TODO add rollback logging
func (d *Db) CreateWithdraw(amount int, clientID int) {
	tx, err := d.client.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(`LOCK TABLE public.withdraw;`)
	if err != nil {
		panic(err)
	}

	sqlInsertQuery := `
		insert into public.withdraw (amount, client_id)
		select $1, $2
		where 0 <= (
			select (
				(
					select coalesce(sum(i.amount), 0)
					from public.invoice i
					where i.client_id = $2
				) - (
					select coalesce(sum(w.amount), 0)
					from public.withdraw w
					where w.client_id = $2
				) - $1
			)
		);
	`
	_, err = tx.Exec(sqlInsertQuery, amount, clientID)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func (d *Db) Clear() {
	_, err := d.client.Exec(`delete from public.withdraw;`)
	if err != nil {
		panic(err)
	}
	_, err = d.client.Exec(`delete from public.invoice;`)
	if err != nil {
		panic(err)
	}
}
