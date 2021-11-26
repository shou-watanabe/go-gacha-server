package mock_repository

import (
	reflect "reflect"
	"techtrain-mission/src/domain/entity"
	"techtrain-mission/src/usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

// type UserUsecase interface {
// 	Create(string) (*entity.User, error)
// 	Get(string) (*entity.User, error)
// 	Update(string, string) (*entity.User, error)
// }

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *entity.User
	var err error
	name := "test"

	ur := NewMockUserRepository(ctrl)
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

func Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *entity.User
	var err error
	token := "test"

	ur := NewMockUserRepository(ctrl)
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

func Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *entity.User
	var err error
	name := "test"
	token := "test"

	ur := NewMockUserRepository(ctrl)
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
