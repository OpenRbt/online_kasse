package svcstorage

import (
	"net"
	"net/http"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"

	"github.com/OpenRbt/lea-central-wash/storageapi/client"
	"github.com/OpenRbt/lea-central-wash/storageapi/model"
	"github.com/OpenRbt/online_kasse/cmd/web/device"
	"github.com/OpenRbt/online_kasse/cmd/web/fptr10"
)

// Client for storage.
type Client struct {
	*client.Storage
}

// NewClient creates and return new client for storageapi.
func NewClient(endpoint string) device.ConfigSvc {
	basePath := client.DefaultBasePath
	schemes := client.DefaultSchemes

	transport := httptransport.New(endpoint, basePath, schemes)
	transport.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	c := &Client{
		Storage: client.New(transport, nil),
	}
	return c
}

// GetConfig getting kasse settings
func (c *Client) GetConfig() (*device.Config, error) {
	res, err := c.Op.Kasse(nil)
	if err != nil {
		return nil, err
	}
	return newConfig(res.Payload)
}

func newConfig(cfg *model.KasseConfig) (*device.Config, error) {
	tax := 0
	switch cfg.Tax {
	case "TAX_VAT110":
		tax = fptr10.LIBFPTR_TAX_VAT110
	case "TAX_VAT0":
		tax = fptr10.LIBFPTR_TAX_VAT0
	case "TAX_NO":
		tax = fptr10.LIBFPTR_TAX_NO
	case "TAX_VAT120":
		tax = fptr10.LIBFPTR_TAX_VAT120
	case "TAX_VAT105":
		tax = fptr10.LIBFPTR_TAX_VAT105
	case "TAX_VAT107":
		tax = fptr10.LIBFPTR_TAX_VAT107
	default:
		return nil, errors.New("unknown tax")
	}
	if cfg.Cashier != "" && len(cfg.CashierINN) != 12 {
		return nil, errors.New("cashier INN required")
	}
	return &device.Config{
		Cashier:         cfg.Cashier,
		CashierINN:      cfg.CashierINN,
		ReceiptItemName: cfg.ReceiptItemName,
		Tax:             tax,
		Timezone:        int(cfg.Timezone),
	}, nil
}
