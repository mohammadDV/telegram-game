package service

import (
	"testing"

	"context"
	"github.com/mohammaddv/telegram-game/internal/entity"
	"github.com/mohammaddv/telegram-game/internal/repository/mocks"

	"github.com/mohammaddv/telegram-game/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_UpdateOrCreate(t *testing.T) {
	accRep := mocks.NewAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(entity.Account{
		ID:        12,
		FirstName: "Reza",
	}, nil).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(ent entity.Account) bool {
		return ent.FirstName == "Ali"
	})).Return(nil).Once()

	newAcc, savedAcc, err := s.UpdateOrCreate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "Ali",
	})

	assert.NoError(t, err)
	assert.Equal(t, savedAcc, false)
	assert.Equal(t, newAcc.FirstName, "Ali")

	accRep.AssertExpectations(t)
}

func TestAccountService_UpdateOrCreateWithUserNoExists(t *testing.T) {
	accRep := mocks.NewAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{}, repository.ErrorNotFound,
	).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(ent entity.Account) bool {
		return ent.FirstName == "Ali"
	})).Return(nil).Once()

	newAcc, savedAcc, err := s.UpdateOrCreate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "Ali",
	})

	assert.NoError(t, err)
	assert.Equal(t, savedAcc, true)
	assert.Equal(t, newAcc.FirstName, "Ali")

	accRep.AssertExpectations(t)
}

func TestAccountService_UpdateOrCreateWithUserHasNotChanged(t *testing.T) {
	accRep := mocks.NewAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(entity.Account{
		ID:        12,
		FirstName: "Ali",
	}, nil).Once()

	newAcc, savedAcc, err := s.UpdateOrCreate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "Ali",
	})

	assert.NoError(t, err)
	assert.Equal(t, savedAcc, false)
	assert.Equal(t, newAcc.FirstName, "Ali")

	accRep.AssertExpectations(t)
}
