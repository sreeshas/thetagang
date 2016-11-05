package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"regexp"
)

type Stock struct {
	C       string `json:"c,omitempty"`  //Change
	CFix    string `json:"c_fix,omitempty"`
	Ccol    string `json:"ccol,omitempty"`
	Cp      string `json:"cp,omitempty"`
	CpFix   string `json:"cp_fix,omitempty"`
	E       string `json:"e,omitempty"`  //Exchange
	ID      string `json:"id,omitempty"`  //
	L       string `json:"l,omitempty"`
	LCur    string `json:"l_cur,omitempty"` //Last Trade with Currency
	LFix    string `json:"l_fix,omitempty"`
	Lt      string `json:"lt,omitempty"`   //Last Trade
	LtDts   string `json:"lt_dts,omitempty"` //Last Trade time stamp.
	Ltt     string `json:"ltt,omitempty"`    //Last trade time
	PclsFix string `json:"pcls_fix,omitempty"`
	S       string `json:"s,omitempty"`
	T       string `json:"t,omitempty"`  //Ticker
}

func getStockValue(symbol string) (Stock , error) {
	resp, err := http.Get("https://www.google.com/finance/info?q="+symbol)

	if err !=nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var stocks []Stock
	//remove first four bytes.
	body = bytes.TrimPrefix(body,[]byte{10,47,47,32})
	//replace newline characters
	body_wo_newline := string(body)
	re := regexp.MustCompile("\\n")
	body_wo_newline = re.ReplaceAllString(body_wo_newline,"")

	error :=json.Unmarshal([]byte(body_wo_newline), &stocks)
	if error!= nil {
		fmt.Println("error:", error)
		return  Stock{}, error
	}
	return stocks[0], nil
}
func main() {

	stockVMW, error := getStockValue("VMW");
	if (error!=nil){
		fmt.Println(error)
	}else{
		fmt.Println(stockVMW.L)
	}

	stockAAPL, error := getStockValue("AAPL");
	if (error != nil) {
		fmt.Println(error)
	} else{
		fmt.Println(stockAAPL.L)
	}

}





