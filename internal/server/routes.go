package server

import (
	"time"

	"github.com/go-ping/ping"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/mullvad-servers", s.mullvadServersHandler)
	s.App.Get("/ping/:hostname", s.pingHandler)
}

func (s *FiberServer) mullvadServersHandler(c *fiber.Ctx) error {
	resp, err := resty.New().R().Get("https://api.mullvad.net/www/relays/all/")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	defer resp.RawResponse.Body.Close()

	return c.Send(resp.Body())
}

func (s *FiberServer) pingHandler(c *fiber.Ctx) error {
	hostname := c.Params("hostname")
	pinger, err := ping.NewPinger(hostname)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	pinger.Count = 1
	pinger.Timeout = time.Second * 2
	pinger.SetPrivileged(false)

	err = pinger.Run()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	stats := pinger.Statistics()
	type MyStat struct {
		*ping.Statistics
		LatencyMS int64 `json:"latency_ms"`
	}

	myStat := MyStat{
		Statistics: stats,
		LatencyMS:  stats.AvgRtt.Milliseconds(),
	}

	return c.JSON(myStat)
}
