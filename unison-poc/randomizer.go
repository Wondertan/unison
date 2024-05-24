package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"time"

	"github.com/iykyk-syn/unison/bapl"
)

func RandomBatches(ctx context.Context, pool bapl.BatchPool, batchSize int, batchTime time.Duration) {
	ticker := time.NewTicker(batchTime)
	defer ticker.Stop()

	log := slog.With("module", "randomizer")
	for {
		select {
		case <-ticker.C:
			go func() {
				batchData := make([]byte, batchSize)
				rand.Read(batchData)
				batch := &bapl.Batch{Data: batchData}
				err := pool.Push(ctx, batch)
				if err != nil {
					log.ErrorContext(ctx, "error pushing batch", "err", err)
				}
				log.DebugContext(ctx, "pushed batch", "hash", base64.StdEncoding.EncodeToString(batch.Hash()))
			}()

		case <-ctx.Done():
			return
		}
	}
}
