package thetagang

import (
	"github.com/piquette/finance-go"
	"math"
	"github.com/piquette/finance-go/quote"
)

type ContractDetail struct {
	Contract *finance.Contract
	TotalCapital float64
}




func (c *ContractDetail) getPercentageReturn() float64 {
	if c.Contract == nil {
		return 0.0
	}
	costForOneContract := c.Contract.Strike * 100
	totalNoContract := math.Floor(c.TotalCapital / costForOneContract)

	totalPremium := totalNoContract * c.Contract.Bid * 100.00

	totalPercentReturn := (totalPremium / c.TotalCapital) * 100.00
	return totalPercentReturn
}

func (c *ContractDetail) getTotalPremiumEarned() float64 {
	if c.Contract == nil {
		return 0.0
	}
	costForOneContract := c.Contract.Strike * 100
	totalNoContract := math.Floor(c.TotalCapital / costForOneContract)

	totalPremium := totalNoContract * c.Contract.Bid * 100
	return totalPremium
}

func (c *ContractDetail) getTotalNoOfContracts() float64 {
	if c.Contract == nil {
		return 0.0
	}
	costForOneContract := c.Contract.Strike * 100
	totalNoContract := math.Floor(c.TotalCapital/ costForOneContract)
	return totalNoContract
}

func (c *ContractDetail) getTwoHundredDayMA() float64 {
	if c.Contract == nil {
		return 0.0
	}

	q, _ := quote.Get(parseSymbol(c.Contract.Symbol))

	return q.TwoHundredDayAverage
}


type ContractDetails []*ContractDetail

func (c ContractDetails) Len() int {
	return len(c)
}

func (c ContractDetails) Less(i, j int) bool {
	return c[i].getPercentageReturn() < c[j].getPercentageReturn()
}
