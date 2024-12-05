package adUC

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/himmel520/media-service/internal/infrastructure/cache"
)

func (uc *AdUC) DeleteCache(ctx context.Context) error {
	return uc.cache.Delete(ctx, cache.AdvPrefixKey)
}

func generateCacheKeyAdv(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return cache.AdvPrefixKey + hash
}
