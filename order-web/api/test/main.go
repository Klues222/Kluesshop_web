package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

var client *alipay.Client

const (
	kAppId       = "9021000135619858"
	kPrivateKey  = "MIIEpgIBAAKCAQEA1yG4QFa1sCHQQS+nI7NdfaDHd6ORch1URZGbhsot+dzvP8In6Vx+jGm/ylHkbtDw7BlGXNs3cHetK17ATdR+FTXIV80JaWS4/DNEfzVo2BWTLMWsHapRNOSnXAVnIPsZCWz/Gi9O1BLJ2E8Dvetg6way2iAzwyMUSTl9/xCIUMlBF7WM+vdRlZGbzsXkIr5rTbfMdwe62NAFA47xN6Io9Bdb8F20ZqAfira/p92ztyEx72xwGybXcIhqyRgZNKDSWIfEjyD6nxzMJilmuYvTwRKt7YkrKm8ErBLtw01Cx1/cu2dKfrlx5KQrASq1CMjk53svfL7XIlSyPPXZSSucBwIDAQABAoIBAQCMs7odIlj8VHEvYSpQmCwqUTCEWA91cic+xOfSdYMya0RzD9oor1z5GuTcP0lHDRK1aCZz7tBIl41D01m0zaiU2LMqcaiZLM4r7J/9DwJ+aUXzv2k9kbpA/NdUwRfdIz0RxB7okk5dN9iZx7vrx/mAXcY2EV+pXdUTuy1j1/0/z93vdYwdj0m9ti38QXZzHBRHM5kYCzUsTWdwEGAhUJeTV4PW4kedeSo+kI/ylC+ormt1pVWyCi4alGwm9hEG+2cxZaMmVvf5ZUZPURY0G2yDZ288Dc5NeP4ewETDCC4vCoGtaP2EfeDyRihxpiVzz6bCQJPlVEeXWyvl0SbVIyXZAoGBAOt1eU6sVxkAgvfrIoqXXZJRC3Cogz/mtDN6K/7QsF7JSJRgwNxZTELorKfuTqDGIOQTgirz6RFK8qpyN3xPRWh+vE8k03P5pvucbcwT4eZHJWYjB/dCKnlW0QYpguzvy0UHqhd5Rd6WDuVyNJ+6Dp6NkDHKBSsSKkDHYt6FDXr7AoGBAOnmRv8X5zuCQA15iJzSSW7ajAuoXP6hdtvwIHpCxUb2nozlg8GjgHptsQP3h2djOygO/9EiiSPtwDAKgvCi5HCc41boY0/UEXjboa+s2ddoqqc3jPjzb0GUd96J5FIJSqL3fkd+AR2QWInzbWtjwZmE5Od1ARYDBY4v6RBD8pVlAoGBAJTAS1nLgN7XtuXfE5xQ0hmMv9h1bS2ilzdqOH8r4jCPox2yHkVW2NnwWptg7yWc5cyREowGObjmC3Zo3+rVvbitUFQDpN7A2qBci/UAnpc3XUYwXWj00RaFJVpqQT9koptCo09fGyfqzxBfXSWHipLaRj1eLnOubykrjaeckQ6fAoGBANAspPm7c5FSXuHfTkiNHNBt8QEbKxFx8dDUioNmVCDRtNGgIMFCXUIZyNfIAhpxhiAkIWkofLiejVP0tw+nWvwjlm8uS92r2JUhnWk2xXfj7yb+2RhlyZqronhNcAnXvTcIYbsNmb6PT04Qe01+LrwNPJIIYBwQOmyPlbZlU03FAoGBAOIc8rg+ANZmPP1dtuBNlRDG/LAEVW0CmuyQwSXX66QvQaiocFkGGK5/JtHoR1o+TQNGFJFyvW+Me4BSgpKaTgIeaNXeKlEX5L7XeSlL+PyNmXY3wDlYZsqXhR+i4Uwzm7iPBemZW7h7QKjB/BMBbwbYhsfa4iH7iTvMB58iCHft"
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1yG4QFa1sCHQQS+nI7NdfaDHd6ORch1URZGbhsot+dzvP8In6Vx+jGm/ylHkbtDw7BlGXNs3cHetK17ATdR+FTXIV80JaWS4/DNEfzVo2BWTLMWsHapRNOSnXAVnIPsZCWz/Gi9O1BLJ2E8Dvetg6way2iAzwyMUSTl9/xCIUMlBF7WM+vdRlZGbzsXkIr5rTbfMdwe62NAFA47xN6Io9Bdb8F20ZqAfira/p92ztyEx72xwGybXcIhqyRgZNKDSWIfEjyD6nxzMJilmuYvTwRKt7YkrKm8ErBLtw01Cx1/cu2dKfrlx5KQrASq1CMjk53svfL7XIlSyPPXZSSucBwIDAQAB"
	// TODO 设置回调地址域名
	kServerDomain = "https://87763681pa.vicp.fun"
)

func main() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		panic(err)
	}
	if err = client.LoadAliPayPublicKey(aliPublicKey); err != nil {
		panic(err)
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:"
	p.OutTradeNo = "bobby_klues"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err := client.TradePagePay(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())

}
