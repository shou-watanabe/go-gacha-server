package mock_usecase

import (
	context "context"
	"testing"

	"github.com/golang/mock/gomock"

	"techtrain-mission/src/domain/entity"
	mock_repository "techtrain-mission/src/mock/repository"
	"techtrain-mission/src/usecase"
)

// type GachaUsecase interface {
// 	Draw(ctx context.Context, times int, token string) ([]*entity.Chara, error)
// }

func TestGachaDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ces []*entity.Chara
	var ue *entity.User
	var err error
	token := "test"
	times := 2
	ctx := context.Background()

	ur := mock_repository.NewMockUserRepository(ctrl)
	cr := mock_repository.NewMockCharaRepository(ctrl)
	ucr := mock_repository.NewMockUserCharaRepository(ctrl)
	cr.EXPECT().List(ctx).Return(ces, nil)
	ur.EXPECT().Get(ctx, token).Return(ue, nil)
	ucr.EXPECT().Store(ctx, *ue, ces).Return(nil)

	gu := usecase.NewGachaUsecase(ur, cr, ucr)
	_, err = gu.Draw(ctx, times, token)

	if err != nil {
		t.Error("Actual Draw(ctx context.Context, times int, token string) is not same as expected")
	}

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Errorf("Actual Draw(ctx context.Context, times int, token string) is not same as expected")
	// }
}
