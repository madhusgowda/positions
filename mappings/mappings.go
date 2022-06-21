package positions

import (
	"os"

)

func TokenContractData(chainName string, protocol string, currency string, stakedTokenAddress string, rewardTokenAddress string, stakeContractAddress string)(string, string, string){
	contractDataMapping := map[string]map[string]map[string]string{
		"ethereum": {
			"openDao": {
				"stakeContractAddress": "0x0d0b9d5e0c5b01e3b8bbeb9f0fe8ebf2497fd02b",
				"stakedTokenAddress": "0x0d0b9d5e0c5b01e3b8bbeb9f0fe8ebf2497fd02b",
				"stakedContractname": "OpenDAO",
				"stakedContractsymbol": "ODAO",
				"stakedContractdecimals": "18",
				"rewardTokenAddress": "0x0d0b9d5e0c5b01e3b8bbeb9f0fe8ebf2497fd02b",
				"rewardContractname": "OpenDAO",
				"rewardContractsymbol": "ODAO",
				"rewardContractdecimals": "18",
			},
		},
	},
		chainData := contractDataMapping[chainName]
		if chainData == nil {
			return "", "", ""
		}
		protocolData := chainData[protocol]
		if protocolData == nil {
			return "", "", ""
		}
		currencyData := protocolData[currency]
		if currencyData == nil {
			return "", "", ""
		}
		return currencyData["stakedContractname"], currencyData["stakedContractsymbol"], currencyData["stakedContractdecimals"]
}
