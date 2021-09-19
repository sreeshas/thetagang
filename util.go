package thetagang

import (
	"os"
	"fmt"
	"time"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/options"
	"github.com/piquette/qtrn/utils"
	"math"
	"sort"
	"regexp"
	tw "github.com/olekukonko/tablewriter"
)


var (
	dateLayout = "2006-01-02"
)


var tickerSymbol = regexp.MustCompile("([A-Z]+)([0-9]+)")

func parseSymbol(s string) (letters string) {
	matches := tickerSymbol.FindStringSubmatch(s)
	return matches[1]
}


func getContract(symbol string, strikeprice float64, date string, totalCapital float64) *ContractDetail {
	t, _ := time.Parse(dateLayout, date)
	dt := datetime.New(&t)
	straddle, err := getStraddleForStrike(symbol, strikeprice, dt)
	if err != nil {
		fmt.Println(err)
	}
	contract := getPutContract(straddle)
	cd := ContractDetail{
		Contract: contract,
		TotalCapital:totalCapital,
	}
	return &cd
}

func getPutContract(straddle *finance.Straddle) *finance.Contract {
	if straddle == nil {
		return nil
	}
	return straddle.Put
}

func getStraddleForStrike(symbol string, strikeprice float64, datetime *datetime.Datetime) (*finance.Straddle, error) {
	p := &options.Params{
		UnderlyingSymbol: symbol,
		Expiration:       datetime,
	}
	iter := options.GetStraddleP(p)

	for iter.Next() {
		straddle := iter.Straddle()
		if straddle.Strike == strikeprice {
			return straddle, nil
		}

	}
	if iter.Err() != nil {
		return nil, iter.Err()
	}

	return nil, nil
}

func buildContractDetails(cd []*ContractDetail) (tbl [][]string) {


	for _, contractDetail := range cd {
		if contractDetail == nil || contractDetail.Contract == nil {
			continue
		}
		row := []string{}
		row = append(row, contractDetail.Contract.Symbol)
		row = append(row, utils.ToStringF(contractDetail.Contract.Strike))
		row = append(row, utils.ToStringF(contractDetail.Contract.Bid))
		row = append(row, utils.ToStringF(contractDetail.getTotalPremiumEarned()))

		row = append(row, getCorrectedExpiryDate(contractDetail.Contract.Expiration))
		row = append(row, utils.ToStringF(contractDetail.getTotalNoOfContracts()))
		row = append(row, utils.ToStringF(contractDetail.getPercentageReturn()))
		row = append(row, utils.ToStringF(contractDetail.getTwoHundredDayMA()))
		tbl = append(tbl, row)

	}
	return tbl
}

func printTableByDate(contractDetails []*ContractDetail) {
	//utils.DateFS(iter.Meta().ExpirationDate + 86400)
	table := tw.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tw.ALIGN_LEFT)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetHeader([]string{
		"Symbol",
		"Strike",
		"Bid",
		"Total Premium Earned",
		"Expiry Date: ",
		"Total No of Contracts",
		"Percentage Return",
		"200MA"})
	table.AppendBulk(buildContractDetails(contractDetails))
	table.Render()
}

func getExpiryDatesForSymbol(s string) ([]string, error) {

	p := &options.Params{
		UnderlyingSymbol: s,
	}
	iter := options.GetStraddleP(p)
	meta := iter.Meta()
	if meta == nil {
		return nil, fmt.Errorf("could not retrieve dates")
	}
	dates := []string{}
	for _, stamp := range meta.AllExpirationDates {
		// set the day to friday instead of EOD thursday..
		// weird math here..
		stamp = stamp + 86400
		t := time.Unix(int64(stamp), 0)

		dates = append(dates, t.Format("2006-01-02"))
	}

	return dates[0:int(math.Min(float64(len(dates)), 10))], nil

}

func printBySymbol(groupBySymbol map[string][]*ContractDetail) {
	for _, value := range groupBySymbol {
		printTableByDate(value)
	}
}

func printByDate(groupByDate map[string][]*ContractDetail) {
	keys := make([]string, 0, len(groupByDate))
	for k := range groupByDate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		printTableByDate(groupByDate[key])
	}
}

func getCorrectedExpiryDate(stamp int) string {
	stamp = stamp + 86400
	t := time.Unix(int64(stamp), 0)
	return t.Format(dateLayout)

}


