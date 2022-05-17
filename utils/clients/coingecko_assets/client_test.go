//go:build integration && vpn
// +build integration,vpn

package coingeckoAssets_test

import (
	"context"

	"os"
	"testing"

	coingeckoAssets "github.com/brave-intl/bat-go/utils/clients/coingecko_assets"
	appctx "github.com/brave-intl/bat-go/utils/context"
	logutils "github.com/brave-intl/bat-go/utils/logging"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/suite"
)

type CoingeckoAssetsTestSuite struct {
	suite.Suite
	redisPool *redis.Pool
	client    coingeckoAssets.Client
	ctx       context.Context
}

func TestCoingeckoTestSuite(t *testing.T) {
	suite.Run(t, new(CoingeckoAssetsTestSuite))
}

var (
	coingeckoAssetsService string = "https://assets.coingecko.com/"
)

func (suite *CoingeckoAssetsTestSuite) SetupTest() {
	// setup the context
	suite.ctx = context.Background()

	// setup debug for client
	suite.ctx = context.WithValue(suite.ctx, appctx.DebugLoggingCTXKey, false)
	// setup debug log level
	suite.ctx = context.WithValue(suite.ctx, appctx.LogLevelCTXKey, "info")

	// setup a logger and put on context
	suite.ctx, _ = logutils.SetupLogger(suite.ctx)

	// setup server location
	suite.ctx = context.WithValue(suite.ctx, appctx.CoingeckoAssetsServerCTXKey, coingeckoAssetsService)

	var redisAddr string = "redis://grant-redis"
	if len(os.Getenv("REDIS_ADDR")) > 0 {
		redisAddr = os.Getenv("REDIS_ADDR")
	}

	suite.redisPool = &redis.Pool{
		MaxIdle:   50,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialURL(redisAddr)
			suite.Require().NoError(err, "failed to connect to redis")
			return conn, err
		},
	}

	rConn := suite.redisPool.Get()
	defer rConn.Close()
	s, err := redis.String(rConn.Do("PING"))
	suite.Require().NoError(err, "failed to connect to redis")
	suite.Require().True(s == "PONG", "bad response from redis")

	// setup the client under test, no redis, will test redis interactions in ratios service
	suite.client, err = coingeckoAssets.NewWithContext(suite.ctx, suite.redisPool)
	suite.Require().NoError(err, "Must be able to correctly initialize the client")
}

func (suite *CoingeckoAssetsTestSuite) TestFetchImageAsset() {
	// PNG
	responseBundle, t1, err := suite.client.FetchImageAsset(suite.ctx, "662", "large", "logo_square_simple_300px.png")
	suite.Require().NoError(err, "should be able to fetch the coin markets")
	suite.Require().True(len(responseBundle.ImageData) > 0, "should have some image data")
	suite.Require().Equal(responseBundle.ContentType, "image/png")

	// Cache works
	responseBundle, t2, err := suite.client.FetchImageAsset(suite.ctx, "662", "large", "logo_square_simple_300px.png")
	suite.Require().NoError(err, "should be able to fetch the coin markets")
	suite.Require().True(t1.Equal(t2), "should have the same last updated timestamp because of cache usage")

	// JPG
	// https://assets.coingecko.com/coins/images/24383/large/apecoin.jpg?1647476455
	responseBundle, _, err = suite.client.FetchImageAsset(suite.ctx, "24383", "large", "apecoin.jpg")
	suite.Require().NoError(err, "should be able to fetch the coin markets")
	suite.Require().True(len(responseBundle.ImageData) > 0, "should have some image data")
	suite.Require().Equal(responseBundle.ContentType, "image/jpeg")
}
