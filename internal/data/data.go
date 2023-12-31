package data

import (
	"context"
	iV1 "sample/api/interface/v1"
	"sample/helper/utils"
	"sample/internal/conf"
	"sample/internal/data/ent"
	"sample/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-kratos/kratos/v2/log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	// postgres driver
	_ "github.com/lib/pq"
)

// Data .
type Data struct {
	ec *ent.Client
	rc redis.Cmdable

	log *log.Helper
}

// NewData .
func NewData(entClient *ent.Client, rc redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "sample/data"))

	d := &Data{
		ec:  entClient,
		rc:  rc,
		log: logHelper,
	}

	cleanup := func() {
		logHelper.Info("closing the data resources")
		if err := d.ec.Close(); utils.IsNotNil(err) {
			logHelper.Error(err)
		}
	}

	return d, cleanup, nil
}

// NewEntClient - new ent client
func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	logHelper := log.NewHelper(log.With(logger, "module", "sample/data/ent"))
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if utils.IsNotNil(err) {
		logHelper.Fatalf("failed opening connection to db: %v", err)
	}
	if conf.Database.Migrate {
		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); utils.IsNotNil(err) {
			logHelper.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return client
}

// NewServiceClient - connect sample service
func NewServiceClient(r registry.Discovery, logger log.Logger) iV1.GreeterClient {
	logHelper := log.NewHelper(log.With(logger, "module", "sample/data/discovery"))
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///sample.tone"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			logging.Client(logger),
			metadata.Client(),
		),
	)

	if utils.IsNotNil(err) {
		logHelper.Fatalf("failed discovering sample service: %v", err)
		panic(err)
	}

	return iV1.NewGreeterClient(conn)
}

// NewRedisCmd - new redis.
func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logHelper := log.NewHelper(log.With(logger, "module", "sample/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Username:     conf.Redis.Username,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), conf.Redis.DialTimeout.AsDuration())
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if utils.IsNotNil(err) {
		logHelper.Fatalf("redis connect error: %v", err)
	}
	return client
}

// getCacheKey - define cache key of the sample service.
// @param key - format: entity:%s
func (d *Data) getCacheKey(key string) string {
	return "sample_cache_" + key
}

// deleteCache - del cache.
// @param ctx - context
// @param key - format: entity:%s
func (d *Data) deleteCache(ctx context.Context, key string) {
	err := d.rc.Del(ctx, key).Err()
	if utils.IsNotNil(err) {
		d.log.Errorf("fail to delete cache:redis.Del(%v) error(%v)", key, err)
	}
}
