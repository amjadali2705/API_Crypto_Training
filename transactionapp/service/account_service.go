package service

import (
	"transactionapp/dao"
	"transactionapp/model"
)

func GetAccounts() ([]model.Account, error) {
	return dao.GetAllAccounts()
}

func GetAccount(id int) (model.Account, error) {
	return dao.GetAccountByID(id)
}

func Create(account model.Account) error {
	return dao.CreateAccount(account)
}

func Update(id int, user model.Account) error {
	return dao.UpdateAccount(id, user)
}

func Delete(id int) error {
	return dao.DeleteAccount(id)
}

func Transfer(fromID, toID int, amount float64) error {
	return dao.TransferAmount(fromID, toID, amount)
}

func Statement(id int) ([]model.Transaction, error) {
	return dao.MiniStatement(id)
}
