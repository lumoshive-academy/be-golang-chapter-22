package repository

import (
	"be-golang-chapter-22/template-html-implem/model"
	"database/sql"
	"errors"
	"fmt"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return CustomerRepository{DB: db}
}

func (cr *CustomerRepository) Login(customer *model.Customer) error {
	query := `SELECT username, password, email FROM customers WHERE username=$1 AND password=$2`
	err := cr.DB.QueryRow(query, customer.Username, customer.Password).Scan(&customer.Username, &customer.Password, &customer.Email)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CustomerRepository) CustomerByID(id int) (*model.Customer, error) {
	customer := model.Customer{}
	query := `SELECT name, username, password, email, status, token FROM customers WHERE id=$1`
	err := cr.DB.QueryRow(query, id).Scan(&customer.Name, &customer.Username, &customer.Password, &customer.Email, &customer.Status, &customer.Token)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cr *CustomerRepository) AllCustomer() (*[]model.Customer, error) {
	customers := []model.Customer{}
	query := `SELECT id, name, username, password, email, status FROM customers`
	rows, err := cr.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Username, &customer.Password, &customer.Email, &customer.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &customers, nil
}

func (cr *CustomerRepository) Register(customer *model.Customer) error {
	query := `INSERT INTO customers (name, username, password, email, status, token) VALUES ($1, $2, $3, $4, $5, $6)`
	result, err := cr.DB.Exec(query, customer.Name, customer.Username, customer.Password, customer.Email, customer.Status, customer.Token)
	if err != nil {
		return err
	}

	row, _ := result.RowsAffected()
	if row <= 0 {
		return errors.New("insert failed")
	}

	return nil
}

func (cr *CustomerRepository) CheckToken(token string) string {
	var tokenResult string
	query := `SELECT token FROM customers WHERE token=$1`
	err := cr.DB.QueryRow(query, token).Scan(&tokenResult)
	// if row.{
	// 	return row.Err()
	// }

	fmt.Println(err)
	return tokenResult
}
