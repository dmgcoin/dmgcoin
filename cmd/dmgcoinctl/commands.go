package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dmgcoin/dmgcoin/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.DmgcoinMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.DmgcoinMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.DmgcoinMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.DmgcoinMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.DmgcoinMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.DmgcoinMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.DmgcoinMessage_BanRequest{}),
	reflect.TypeOf(protowire.DmgcoinMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
