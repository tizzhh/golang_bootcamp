package main

import "C"

/*
#include "cow.c"
*/

import (
	"candyShop/client"
	"candyShop/client/operations"
	"candyShop/restapi"
	"context"
	"strconv"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"

	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"

	"net/http"
	"os"
	"time"
)

const (
	NOT_SPEC_CANDY_COUND int64  = -420
	NOT_SPEC_CANDY_TYPE  string = "blank"
	NOT_SPEC_MONEY       int64  = -420
)

const (
	CERT_PATH string = "cert/client/cert.pem"
	KEY_PATH  string = "cert/client/key.pem"
	CA_PATH   string = "cert/minica.pem"
)

var candyCount = flag.Int64("c", -420, "Amount of candy")
var candyType = flag.String("k", "blank", "Type of candy: CE, AA, NT, DE or YR")
var money = flag.Int64("m", -420, "Amount of money")

func main() {
	flag.Parse()
	if *candyCount == -420 || *money == -420 {
		flag.Usage()
		os.Exit(1)
	}
	if _, ok := restapi.ShopPrices[*candyType]; !ok {
		flag.Usage()
		os.Exit(1)
	}

	cert, err := tls.LoadX509KeyPair(CERT_PATH, KEY_PATH)
	if err != nil {
		panic(err)
	}
	caCert, err := os.ReadFile(CA_PATH)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	transport := httptransport.NewWithClient(
		"candy.tld:3333",
		"",
		[]string{"https"},
		&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	)

	c := client.New(transport, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer cancel()
	params := operations.BuyCandyParams{Context: ctx, Order: operations.BuyCandyBody{CandyCount: candyCount, CandyType: candyType, Money: money}}

	res, err := c.Operations.BuyCandy(&params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error()[strings.Index(err.Error(), "&{")+2:len(err.Error())-1])
		os.Exit(1)
	}

	fmt.Println(res.Payload.Thanks + " Your change is " + strconv.FormatInt(res.Payload.Change, 10))
}
