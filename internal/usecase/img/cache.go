package imgUC

import "context"

const (
	logoCachePrefix  = "logo:*"
	allLogosCachekey = "logo:all"
)

func (uc *ImgUC) DeleteCache(ctx context.Context) {
	if err := uc.cache.Delete(context.Background(), logoCachePrefix); err != nil {
		uc.log.Error(err)
	}
}
