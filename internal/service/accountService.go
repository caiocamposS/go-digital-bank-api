package service

import (
	"crypto/rand"
	"digital-bank-api/internal/dto/request"
	"digital-bank-api/internal/dto/response"
	"digital-bank-api/internal/models"
	"digital-bank-api/internal/repository"
	"math/big"
)

type AccountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: repo,
	}
}

func (s *AccountService) CreateAccount(req *request.AccountRequest, userID uint) (*response.AccountResponse, error) {
	number, err := generateAccountNumber()

	if err != nil {
		return nil, err
	}

	account := models.Account{
		UserID: userID,
		Number: number,
		Balance: 0,
		Type: req.Type,
	}

	err = s.accountRepository.Create(&account)

	if err != nil {
		return nil, err
	}

	resp := response.AccountResponse{
		ID: account.ID,
		UserID: account.UserID,
		Number: account.Number,
		Balance: account.Balance,
		Type: account.Type,
	}

	return &resp, nil
}


func generateAccountNumber() (string, error) {
	const digits = "0123456789"
	const length = 10

	accountNumber := make([]byte, length)

	for i := range accountNumber {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}

		accountNumber[i] = digits[n.Int64()]
	}

	return string(accountNumber), nil
}