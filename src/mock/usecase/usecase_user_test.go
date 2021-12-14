package mock_usecase

import (
	context "context"
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"techtrain-mission/src/domain/entity"
	mock_repository "techtrain-mission/src/mock/repository"
	"techtrain-mission/src/usecase"
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
	ctx := context.Background()

	ur := mock_repository.NewMockUserRepository(ctrl)
	ur.EXPECT().Create(ctx, name).Return(expected, err)

	uu := usecase.NewUserUsecase(ur)
	result, err := uu.Create(ctx, name)

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

	var expected *entity.User
	var err error
	token := "test"
	ctx := context.Background()

	ur := mock_repository.NewMockUserRepository(ctrl)
	ur.EXPECT().Get(ctx, token).Return(expected, err)

	uu := usecase.NewUserUsecase(ur)
	result, err := uu.Get(ctx, token)

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

	var expected *entity.User
	var err error
	name := "test"
	token := "test"
	ctx := context.Background()

	ur := mock_repository.NewMockUserRepository(ctrl)
	ur.EXPECT().Update(ctx, name, token).Return(expected, err)

	uu := usecase.NewUserUsecase(ur)
	result, err := uu.Update(ctx, name, token)

	if err != nil {
		t.Error("Actual Update(name string, token string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Update(name string, token string) is not same as expected")
	}
}
