package rds

import (
	"BalanceRange/internal/config"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	opts, err := getRedisOptions(cfg)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	result := client.Ping(context.Background())
	if result.Err() != nil {
		return nil, result.Err()
	}

	res := client.FlushDB(context.Background())
	if res.Err() != nil {
		panic(res.Err())
	}
	client.Ping(context.Background()) // TODO: remove this

	return client, nil
}

func getRedisOptions(cfg config.RedisConfig) (opts *redis.Options, err error) {
	if cfg.UseCertificates {
		certs := make([]tls.Certificate, 0)
		if cfg.CertificatesPaths.Cert != "" && cfg.CertificatesPaths.Key != "" {
			cert, err := tls.LoadX509KeyPair(cfg.CertificatesPaths.Cert, cfg.CertificatesPaths.Key)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"certPath: %v, keyPath: %v",
					cfg.CertificatesPaths.Cert,
					cfg.CertificatesPaths.Key,
				)
			}
			certs = append(certs, cert)
		}
		caCert, err := os.ReadFile(cfg.CertificatesPaths.Ca)
		if err != nil {
			return nil, errors.Wrapf(err, "ca load path: %v", cfg.CertificatesPaths.Ca)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		opts = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			MinIdleConns: cfg.MinIdleConns,
			PoolSize:     cfg.PoolSize,
			PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
			Password:     cfg.Password,
			DB:           cfg.DB,
			TLSConfig: &tls.Config{
				InsecureSkipVerify: cfg.InsecureSkipVerify,
				Certificates:       certs,
				RootCAs:            caCertPool,
			},
		}
	} else {
		opts = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			MinIdleConns: cfg.MinIdleConns,
			PoolSize:     cfg.PoolSize,
			PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
			Password:     cfg.Password,
			DB:           cfg.DB,
		}
	}

	return opts, err
}
