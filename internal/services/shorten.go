package service

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type Shorten struct {
	redis *redis.Client
}

func NewShorten(r *redis.Client) *Shorten {
	return &Shorten{
		redis: r,
	}
}

func (s *Shorten) CreateShorten(ctx context.Context, originalUrl string) (*string, error) {
	key := generateKey(6)

	err := s.redis.Set(ctx, key, originalUrl, 0).Err()

	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (s *Shorten) GetOriginalUrl(ctx context.Context, shortenKey string) (*string, error) {
	originalUrl, err := s.redis.Get(ctx, shortenKey).Result()

	if err != nil {
		return nil, err
	}

	return &originalUrl, nil
}

func (s *Shorten) Delete(ctx context.Context, shortenKey string) error {
	err := s.redis.Del(ctx, shortenKey).Err()

	if err != nil {
		return err
	}

	return nil
}

func generateKey(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}
