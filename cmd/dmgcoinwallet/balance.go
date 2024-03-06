package main

import (
	"context"
	"fmt"

	"github.com/dmgcoin/dmgcoin/cmd/dmgcoinwallet/daemon/client"
	"github.com/dmgcoin/dmgcoin/cmd/dmgcoinwallet/daemon/pb"
	"github.com/dmgcoin/dmgcoin/cmd/dmgcoinwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatDmgc(addressBalance.Available), utils.FormatDmgc(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, DMGC %s %s%s\n", utils.FormatDmgc(response.Available), utils.FormatDmgc(response.Pending), pendingSuffix)

	return nil
}
