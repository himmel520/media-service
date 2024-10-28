package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateCacheKey(limit, offset int, posts, priority []string) string {
	key := fmt.Sprintf("%d:%d:%s:%s", limit, offset, strings.Join(posts, ","), strings.Join(priority, ","))

	// Создаем хеш
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return "advs:" + hash
}
