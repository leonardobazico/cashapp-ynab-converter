# CLI to convert files into You Need a Budget (YNAB) format

[![ynab-converter-pipeline](https://github.com/leonardobazico/ynab-converter/actions/workflows/go.yml/badge.svg)](https://github.com/leonardobazico/ynab-converter/actions/workflows/go.yml)
![coverage](https://raw.githubusercontent.com/leonardobazico/ynab-converter/badges/.badges/main/coverage.svg)

This solution is based on the [Formatting a CSV File](https://support.ynab.com/en_us/formatting-a-csv-file-an-overview-BJvczkuRq)

## CashApp

The CashApp csv report is a bit different than the YNAB csv report. This tool converts the CashApp csv report to the YNAB csv report.

The CashApp csv input file looks like this:

```csv
"Transaction ID","Date","Transaction Type","Currency","Amount","Fee","Net Amount","Asset Type","Asset Price","Asset Amount","Status","Notes","Name of sender/receiver","Account"
"ywur7r","2023-10-07 16:12:36 EDT","Boost Payment","USD","$1","$0","$1","","","","CARD REFUNDED","Cash Reward","","Visa Debit 0987"
"rmgsrz","2023-10-06 23:59:59 EDT","Cash Card Debit","USD","-$2.90","$0","-$2.90","","","","CARD CHARGED","MTA*NYCT PAYGO","","Visa Debit 0987"
"s048op","2023-10-06 19:19:48 EDT","Cash Card Debit","USD","-$2.90","$0","-$2.90","","","","CARD CHARGED","MTA*NYCT PAYGO","","Visa Debit 0987"
"rnkyxvb","2023-06-13 00:00:57 EDT","Sent P2P","USD","-$10","$0","-$10",,"",,"PAYMENT SENT",,"Some business name","Visa Debit 1230"

```

The YNAB csv output file looks like this:

```csv
Date,Payee,Memo,Amount
10/07/2023,Cash Reward,CARD REFUNDED,1.00
10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90
10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90
06/13/2023,Some business name,PAYMENT SENT,-10.00

```

### Usage

```bash
ynabconverter cashapp -file cashapp_report.csv > ynab_transactions.csv
```

## Tech stack

This is a simple cli tool written in Go lang.
gotestsum is used for testing.

### Continuous Integration

This project uses GitHub Actions for CI.
In order to test the GIthub Actions locally, you can use [act](https://github.com/nektos/act).

```bash
act -s GITHUB_TOKEN="$(gh auth token)"
```
