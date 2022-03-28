package cache

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	address  string
	password string
	options  *redis.Options
	dbs      map[int]*redis.Client
}

func NewRedis(addr, passwd string) *Redis {
	return &Redis{
		address:  addr,
		password: passwd,
		dbs:      make(map[int]*redis.Client),
	}
}

func (r *Redis) SetOptions(o *redis.Options) {
	r.options = o
	if r.options.Addr == "" {
		r.options.Addr = r.address
	}
	if r.options.Password == "" {
		r.options.Password = r.password
	}
}

func (r *Redis) GetClient(db int) (*redis.Client, error) {
	if cli, ok := r.dbs[db]; ok {
		return cli, nil
	}
	if r.options == nil {
		r.options = &redis.Options{
			Addr:     r.address,
			Password: r.password,
			DB:       db,
		}
	}
	cli := redis.NewClient(r.options)
	r.dbs[db] = cli
	_, err := cli.Ping().Result()
	return cli, err
}

/*********************************RedisClient*******************************************************/
const defualtDB = 0

var (
	RedisClient   *redis.Client
	ClusterClient *redis.ClusterClient
	Rediser       *Redis
)

func InitRedis(address, password string) (err error) {
	Rediser = NewRedis(address, password)
	RedisClient, err = Rediser.GetClient(defualtDB)
	return err
}

/*********************************RedisOperationFuncttion*********************************************/
// 普通模式
func NewRedisClient(addr, passwd string, db int) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	_, err := cli.Ping().Result()
	return cli, err
}

// 哨兵模式
func NewSentienlClient(sentinelAddrs []string, masterName string) (*redis.Client, error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: sentinelAddrs,
	})
	_, err := rdb.Ping().Result()
	//RedisClient = rdb
	return rdb, err
}

// 集群模式
func NewClusterClient(clusterAddrs []string) (*redis.ClusterClient, error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: clusterAddrs,
	})
	_, err := rdb.Ping().Result()
	//ClusterClient = rdb
	return rdb, err
}
