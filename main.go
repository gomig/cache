package cache

import (
	"os"
	"path"
	"time"

	"github.com/go-redis/redis/v8"
)

// NewRedisCache create a new redis cache manager instance
func NewRedisCache(prefix string, opt redis.Options) Cache {
	rc := new(rCache)
	rc.init(prefix, opt)
	return rc
}

// NewFileCache create a new file cache manager instance
func NewFileCache(prefix string, dir string) Cache {
	fc := new(fCache)
	fc.init(prefix, dir)
	return fc
}

// NewRateLimiter create a new rate limiter
func NewRateLimiter(key string, maxAttempts uint32, ttl time.Duration, cache Cache) (RateLimiter, error) {
	limiter := new(rLimiter)
	if err := limiter.init(key, maxAttempts, ttl, cache); err != nil {
		return nil, err
	} else {
		return limiter, nil
	}
}

// NewVerificationCode create a new verification code manager instance
func NewVerificationCode(key string, ttl time.Duration, cache Cache) (VerificationCode, error) {
	vc := new(vcDriver)
	if err := vc.init(key, ttl, cache); err != nil {
		return nil, err
	} else {
		return vc, nil
	}
}

// CleanFileExpiration clean file cache expired records
func CleanFileExpiration(dir string) error {
	files, err := os.ReadDir("./")
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		bytes, err := os.ReadFile(path.Join(dir, f.Name()))
		if err != nil {
			continue
		}

		rec := record{}
		if err := rec.Deserialize(string(bytes)); err != nil {
			continue
		}

		if rec.IsExpired() {
			os.Remove(path.Join(dir, f.Name()))
		}
	}

	return nil
}
