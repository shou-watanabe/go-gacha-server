package mock_usecase

import (
	context "context"
	"testing"

	"github.com/golang/mock/gomock"

	"techtrain-mission/src/domain/entity"
	mock_repository "techtrain-mission/src/mock/repository"
	"techtrain-mission/src/usecase"
)

// type CharaUsecase interface {
// 	List(ctx context.Context, token string) ([]*entity.UserChara, error)
// }

func TestCharaList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ue *entity.User
	// var uces []*entity.UserChara
	var err error
	token := "test"
	ctx := context.Background()

	ur := mock_repository.NewMockUserRepository(ctrl)
	ucr := mock_repository.NewMockUserCharaRepository(ctrl)
	ur.EXPECT().Get(ctx, token).Return(ue, nil)
	// ucr.EXPECT().List(ctx, &ue).Return(uces, nil)

	cu := usecase.NewCharaUsecase(ur, ucr)
	_, err = cu.List(ctx, token)

	if err != nil {
		t.Error("Actual List(ctx context.Context, token string) ([]*entity.UserChara, error) is not same as expected")
	}

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Errorf("Actual List(ctx context.Context, token string) ([]*entity.UserChara, error) is not same as expected")
	// }
}
