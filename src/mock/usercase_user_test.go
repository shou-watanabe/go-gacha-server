package mock_repository

import (
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"techtrain-mission/src/domain/entity/user"
	"techtrain-mission/src/usecase"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *user.Entity
	var err error
	name := "test"

	ur := NewMockRepository(ctrl)
	ur.EXPECT().Create(name).Return(expected, err)

	uu := usecase.NewUserUsecase(ur)
	result, err := uu.Create(name)

	if err != nil {
		t.Error("Actual Create(name string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create(name string) is not same as expected")
	}
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *user.Entity
	var err error
	token := "test"

	ur := NewMockRepository(ctrl)
	ur.EXPECT().Get(token).Return(expected, err)

	uu := usecase.NewUserUsecase(ur)
	result, err := uu.Get(token)

	if err != nil {
		t.Error("Actual Get(token string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Get(token string) is not same as expected")
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *user.Entity
	var err error
	name := "test"
	token := "test"

	ur := NewMockRepository(ctrl)
	ur.EXPECT().Update(name, token).Return(expected, err)

	uu := usecase.NewUserUsecase(ur)
	result, err := uu.Update(name, token)

	if err != nil {
		t.Error("Actual Update(name string, token string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Update(name string, token string) is not same as expected")
	}
}
