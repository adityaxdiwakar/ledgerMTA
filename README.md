# ledgerMTA
A ledger-cli tool to track MTA expenses (assuming usage of CashApp Boost & Fare-Capping). It will automatically determine (using your ledger file) how much you have saved using CashApp, FareCapping, etc. It will also produce a per-ride cost figure for any specified period.

It does require three accounts: `Expenses:Travel:Subway`, `Expenses:Travel:Subway:Boost`, and `Expenses:Travel:Subway:Farecap` where the full `$2.75` is always billed to the first account, paid via the `...:Farecap`when using farecap. If a CashApp boost applies, `$1` is paid via `...:Boost`

## Installation
```sh
git clone git@github.com:adityaxdiwakar/ledgerMTA.git
cd ledgerMTA/
go build .
mv ledgerMTA /usr/local/bin/.
```

## Usage
Simply run `ledgerMTA` with any additional arguments (i.e. `--begin` or `--end`) as they will be passed to ledger-cli. You also must pass your records file using the `LEDGER_PATH` environment variable.
```sh
➜  LEDGER_PATH=records.ldg ledgerMTA
Total Rides Taken:      34      $93.50
Paid Rides Taken:       25      $49.75
Full Cost Rides:        6       $16.50
Cost per Ride:          N/A     $1.46
➜  LEDGER_PATH=records.ldg ledgerMTA --begin 2022/06/06      
Total Rides Taken:      17      $46.75
Paid Rides Taken:       12      $23.00
Full Cost Rides:        2       $5.50
Cost per Ride:          N/A     $1.35
➜  LEDGER_PATH=records.ldg ledgerMTA --begin 2022/05/30 --end 2022/06/05 
Total Rides Taken:      16      $44.00
Paid Rides Taken:       13      $26.75
Full Cost Rides:        4       $11.00
Cost per Ride:          N/A     $1.67
```

## Pro Tip
In your shell's RC file, add the line:
```sh
export LEDGER_PATH="..."
```
then run `source ~/.<shell>rc` and then you can run `ledgerMTA` without needing to provide the location of your records file each time.
