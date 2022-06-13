package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	path := os.Getenv("LEDGER_PATH") + " "

	defaultArgs := strings.Split("-f "+path+"bal Expenses:Travel:Subway", " ")

	withoutSavings := append(defaultArgs,
		strings.Split("and not ("+
			"Expenses:Travel:Subway:Boost or "+
			"Expenses:Travel:Subway:Farecap)", " ")...)
	cmd := exec.Command("ledger", append(withoutSavings, os.Args[1:]...)...)

	out, err := cmd.Output()
	if err != nil {
		log.Fatalln("could not perform analysis, got", err)
	}

	dollarRegex, _ := regexp.Compile("\\$[-?0-9]+\\.[0-9]+")
	totalValue := dollarRegex.FindString(string(out))[1:]

	spentAmountCmd, err := exec.Command("ledger", append(defaultArgs,
		os.Args[1:]...)...).Output()
	spentAmount := dollarRegex.FindString(string(spentAmountCmd))[1:]

	farecapArgs := strings.Split("-f "+path+"bal "+
		"Expenses:Travel:Subway:FareCap", " ")

	farecapAmt, err := exec.Command("ledger", append(farecapArgs,
		os.Args[1:]...)...).Output()
	farecapSavings := dollarRegex.FindString(string(farecapAmt))[1:]

	boostArgs := strings.Split("-f "+path+"bal "+
		"Expenses:Travel:Subway:Boost", " ")

	boostAmt, err := exec.Command("ledger", append(boostArgs,
		os.Args[1:]...)...).Output()
	boostSavings := dollarRegex.FindString(string(boostAmt))[1:]

	fullAmt, erra := strconv.ParseFloat(totalValue, 64)
	spentAmt, errb := strconv.ParseFloat(spentAmount, 64)
	farecapFloat, errc := strconv.ParseFloat(farecapSavings, 64)
	boostValue, errd := strconv.ParseFloat(boostSavings, 64)

	if erra != nil || errb != nil || errc != nil || errd != nil {
		log.Fatalln("error: could not convert prices to numbers")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t',
		tabwriter.AlignRight)

	fmt.Fprintf(writer, "Total Rides Taken:\t%.0f\t$%.2f\n",
		fullAmt/2.75, fullAmt)
	fmt.Fprintf(writer, "Paid Rides Taken:\t%.0f\t$%.2f\n",
		(fullAmt+farecapFloat)/2.75, spentAmt)

	fullCostRides := (fullAmt+farecapFloat)/2.75 + boostValue
	fmt.Fprintf(writer, "Full Cost Rides:\t%.0f\t$%.2f\n",
		fullCostRides, fullCostRides*2.75)

	fmt.Fprintf(writer, "Cost per Ride:\t%s\t$%.2f\n", "N/A",
		spentAmt/(fullAmt/2.75))
	writer.Flush()
}
