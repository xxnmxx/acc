package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Data struct {
	Base          Base          `json:"基礎情報"`
	TaxReport0    TaxReturn     `json:"申告書情報(直前期)"`
	TaxReport1    TaxReturn     `json:"申告書情報(直前々期)"`
	TaxReport2    TaxReturn     `json:"申告書情報(直前々々期)"`
	FinancialData FinancialData `json:"決算書情報"`
	Index1        Index         `json:"業種目別株価1"`
	Index2        Index         `json:"業種目別株価2"`
}

type Base struct {
	Name          string  `json:"会社名"`
	Date          string  `json:"課税時期(YYYY/MM/DD)"`
	IssuedStock   int     `json:"発行済株式数"`
	TreasuryStock int     `json:"自己株式"`
	Employee      int     `json:"直前期以前1年間の従業員数"`
	Size          float64 `json:"会社規模"`
}

type TaxReturn struct {
	FiscalYear string `json:"事業年度(YYYY/MM/DD)"`
	Div        Div    `json:"配当"`
	Income     Income `json:"所得"`
	Equity     Equity `json:"純資産"`
}

type Div struct {
	Normal int `json:"年配金当額(千円)"`
	Extra  int `json:"非経常的な配当金額(千円)"`
}

type Income struct {
	Income    int `json:"課税所得金額(千円)"`
	Extra     int `json:"非経常的な利益金額(千円)"`
	ExemptDiv int `json:"受取配当金等の益金不算入額(千円)"`
	Wht       int `json:"上記の所得税額(千円)"`
	Nol       int `json:"損金算入した繰越欠損金額(千円)"`
}

type Equity struct {
	Capital int `json:"資本金等の額(千円)"`
	Re      int `json:"利益積立金額(千円)"`
}

type FinancialData struct {
	Asset   Asset `json:"直前期末の総資産価額(千円)"`
	Sales   int   `json:"直前期末以前１年間の取引金額(千円)"`
	Capital int   `json:"直前期末の資本金(千円)"`
}

type Asset struct {
	TotalAsset          int `json:"総資産価額(千円)"`
	AllowanceForBadDebt int `json:"貸倒引当金(千円)"`
}

type Index struct {
	Industry   string  `json:"業種目"`
	IndustryNo int     `json:"番号"`
	A          A       `json:"株価(A)"`
	B          float32 `json:"配当金額(B)"`
	C          int     `json:"利益金額(C)"`
	D          int     `json:"簿価純資産価額(D)"`
}

type A struct {
	A1 int `json:"課税時期の属する月"`
	A2 int `json:"課税時期の属する月の前月"`
	A3 int `json:"課税時期の属する月の前々月"`
	A4 int `json:"前年平均株価"`
	A5 int `json:"課税時期の属する月以前2年間の平均株価"`
}

func (d *Data) InputSheet(s string) {
	f, err := os.Create(s)
	if err != nil {
		log.Fatal(err)
	}
	ipt, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, string(ipt))
}

func DataImport(p string) *Data {
	d := new(Data)
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	if json.Unmarshal(b, d); err != nil {
		log.Fatal(err)
	}
	return d
}
