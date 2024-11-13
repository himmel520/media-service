package adUC

import "context"

func (uc *AdUC) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
