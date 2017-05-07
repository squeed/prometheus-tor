package collector

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/yawning/bulb"
)

func getInfoFloat(c *bulb.Conn, val string) (float64, error) {
	resp, err := c.Request("GETINFO " + val)
	if err != nil {
		return -1, err
	}

	if len(resp.Data) != 1 {
		return -1, errors.Errorf("GETINFO %s returned unknown response", val)
	}

	vals := strings.SplitN(resp.Data[0], "=", 2)
	if len(vals) != 2 {
		return -1, errors.Errorf("GETINFO %s returned invalid response", val)
	}

	return strconv.ParseFloat(vals[1], 64)
}
