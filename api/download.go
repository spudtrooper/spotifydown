package api

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/request"
)

var (
	// id="hash" name="hash" value="SPFiIsrwRHu9u6oh44df6fgaw8qDbwDgM4D3VlR9OEB+4WgBSYRmmFFuuoMmhED2MladMUX/UkNaKdVWKhI3M3G8bAye1gJNW++kdsMXswTSIbVOUxrqE4MIUDoaPjolI57vW421UmygjDz8mRrUHtOGrQR8WiYRCZZdYFEP7CdL9WHGOJ4T0YWJZyi8SI7801/NW45vH8c2GnZoYWuE8Ru05qCqMvhnkFPVp/ZmRw/XMMTbPdTcY0vMbw9UAFye/BXeEtARDqniw/UGmvXYs7ncUgsZqgsrUH+RwIzzVXmB6pfjxHspIbjv9qhVuphUUDsYjpWhZ6ECsd7qgx+yosEM3vNAl7LLf5J5lpSg840iVYyra27LEv894D5LmN7mPejpv55uaASTQWRmVJAd1NvEDjYJys/HL6rRkWUixPycfiwfTrql0NDpMQWQkMNUiLKqBBuRnYnKVQniiOAVkNrKQg5s5ZX2gedR0g1Y78SJw8ANBbFPIPpv7P/LQSmLM17fVB+ozrhRSBVRsAPodCDF7e6LLV1qclVu8ojhhmcEORfLoLqGwD+GHUbV4cjJyxrVL2+S4+36MKIgRKoqlW7BsgPFEV57x+AgFOdNnbYPgiyAvXbWk9gPCUWOvDvQi0QSgIwpFfnvYD5Lv+B3GhG4A2Ixt72150PYXg1EXwGU09n86dcKvOMe/E4nZgEI6S2iGSzPoDm7ZOAc7EQ5lS0u0ivZgFsD8+G6N4BrJQFZJCeDAasEHvNZ1CrNaen0A0fmY2OPvQN4MnXqJTVLWmKcvYqwLhmJ6HEWZvhO3kEvZkui/73gXmTiJLsk8YitRKEr4ZKp1CBA2Eu2zlT1+Wly2sbwKeej1He3Mkf6gScmuWyXPzJuSjkINMBwW0jIcfinHcRtUMd0VFTey3UD6hXLmhGHKdMmLEhxX5lkVIoGuE/xX3P3OmpaD8iku2HfHnYX9DSL5gEBt3GQRZcYnSdJEXH4L0Fm2/ZFAVJX4JZc2TnU3HK+Y9Z7+aCW9NT6VqaAH4XATq2pLu60k1bZwB+fnXGZsVqWcT17ccWtGWvDNDEqGJiUSjEhKh7Du7h5vGuSK9t3lXlATb6wuDw5Ogy+2CRdrw7IAVJBCXw5IerXgbBZ7xmINUaHrhXa15/zQOwgKESDshEuX/t6EHT8r2SmXcCsW1Nw5J2q3y/P7mlVOZHyz465L0XtS4Ce293JSbQCWAoYklWuQRSLLeF4UhueeGYJv/LAnwwouJ9YK9R9f3f+h5eTrFKUOnTa923enX47V/+oEyaz8r+7SLTqUfdfI1Dkfb19syOlgzziuKFZaQ7ZcvKYYm5o8q4ar/ikzWDAcQGo1K/9ghRsWJkZdOksuBYKOqRZLyrTO3SW5WJTIxZIRVLrS0/KTIAQobxMRp0WP2yQ4MF3tLWTjZtN2oDQckzPpqwPch4ACnBn2MJYtKhHPXUcYi/EzjewGz2CnTob4L/gcLBSeHJtN9OspQ=="
	hashRE = regexp.MustCompile(`id="hash" name="hash" value="([^"]+)"`)
	// "taskId": "d10e1f29-64d7-4ee7-826b-1687e071f74a"
	taskIDRE = regexp.MustCompile(`"taskId": "([^"]+)"`)
)

func findHash(vid string) (string, error) {
	uri := request.CreateRoute("https://ytmp3api.net/iframe",
		request.MakeParam("color", "green"),
		request.MakeParam("vid", vid),
	)

	res, err := request.Get(uri, nil)
	if err != nil {
		return "", err
	}

	m := hashRE.FindStringSubmatch(string(res.Data))
	if len(m) != 2 {
		return "", errors.Errorf("could not find hash in %s", string(res.Data))
	}
	hash := m[1]

	return hash, nil
}

func startConvert(hash string) (string, error) {
	const uri = "https://ytmp3api.net/iframe/convert.php"

	data := url.Values{}
	data.Set("hash", hash)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	res, err := request.Post(uri, nil, strings.NewReader(data.Encode()), request.RequestExtraHeaders(headers))
	if err != nil {
		return "", err
	}

	m := taskIDRE.FindStringSubmatch(string(res.Data))
	if len(m) != 2 {
		return "", errors.Errorf("could not find task ID for hash %s", hash)
	}
	taskID := m[1]

	return taskID, nil
}

type DownloadInfo struct {
	Hash   string
	TaskID string
}

//go:generate genopts --params --function Download --extends Base id:string
func (c *Client) Download(optss ...DownloadOption) (*DownloadInfo, error) {
	opts := MakeDownloadOptions(optss...)
	verbose := opts.Verbose()

	hash, err := findHash(opts.Id())
	if err != nil {
		return nil, err
	}
	if verbose {
		c.logger.Printf("found hash: %+v", hash)
	}

	taskID, err := startConvert(hash)
	if err != nil {
		return nil, err
	}
	if verbose {
		c.logger.Printf("found taskID: %+v", taskID)
	}

	res := &DownloadInfo{
		Hash:   hash,
		TaskID: taskID,
	}

	return res, nil
}
