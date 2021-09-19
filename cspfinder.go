package thetagang

import (
	"github.com/sreeshas/thetagang/config"
	"fmt"
)

type CspFinder struct {
	SymbolMap map[string]float64
	TotalCapital float64
	GroupByDate bool
	GroupBySymbol bool
}

func NewCspFinder(configData config.Config) *CspFinder {
	return &CspFinder{SymbolMap:configData.SymbolMap,
	TotalCapital:configData.TotalCapital,
	GroupByDate:configData.OutputFormat.GroupByDate,
	GroupBySymbol:configData.OutputFormat.GroupBySymbol}
}


func(c *CspFinder) Execute() {

	c.getData()


}

func(c *CspFinder) getData() {
	groupByDateMap := map[string][]*ContractDetail{}
	groupBySymbolMap := map[string][]*ContractDetail{}
	for symbol, price := range c.SymbolMap {
		expiryDates, _ := getExpiryDatesForSymbol(symbol)
		for _, date := range expiryDates {
			contractInfo := getContract(symbol, price, date, c.TotalCapital)
			if contractDetails, ok := groupByDateMap[date]; ok {
				contractDetailsList := append(contractDetails, contractInfo)
				groupByDateMap[date] = contractDetailsList
			} else {
				contractDetails := []*ContractDetail{}
				contractDetailsList := append(contractDetails, contractInfo)
				groupByDateMap[date] = contractDetailsList
			}
			if contractDetails, ok := groupBySymbolMap[symbol]; ok {
				contractDetailsList := append(contractDetails, contractInfo)
				groupBySymbolMap[symbol] = contractDetailsList
			} else {
				contractDetails := []*ContractDetail{}
				contractDetailsList := append(contractDetails, contractInfo)
				groupBySymbolMap[symbol] = contractDetailsList
			}

		}
	}

	fmt.Println("Total Capital: ", c.TotalCapital)
	if c.GroupBySymbol {
		printBySymbol(groupBySymbolMap)
	}
	if c.GroupByDate {
		printByDate(groupByDateMap)
	}


}











