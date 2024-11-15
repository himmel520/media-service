package adUC

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const advCachePrefix = "advs:*"

func (uc *AdUC) DeleteCache(ctx context.Context) error {
	return uc.cache.Delete(ctx, advCachePrefix)
}

func generateCacheKeyAdv(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return advCachePrefix + hash
}
