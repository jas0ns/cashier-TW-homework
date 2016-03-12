package cashier

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Barcode string
type Category string

type Commodity struct {
	name     string
	unit     string
	price    float64
	category Category
	p_head   *Promotions
	p_tail   *Promotions
}

type ShoppingItem struct {
	commodity *Commodity
	amount    int
	subtotal  float64
	allowance float64
	/* indicate promotions this ShoppingItem got
	0x0001 << 0				 DiscountPromotions
	0x0001 << 1				 ThreeForTwoPromotions
	...
	*/
	promotions_flags int32
}

type Input struct {
	Barcode string
	Amount  int
}

type Request struct {
	inputs           []Input
	si_list          []*ShoppingItem
	promotions_flags int32 //same as ShoppingItem
}

const (
	DISCOUNTPROMOTIONS    = 0x0001 << 0
	THREEFORTWOPROMOTIONS = 0x0001 << 1
)

func (r *Request) InitRequest(
	inputStr *string, comMap_p *map[Barcode]*Commodity) {
	dec := json.NewDecoder(strings.NewReader(*inputStr))
	inputs := make([]Input, 0)
	var in Input
	for {
		if err := dec.Decode(&in); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, in)
	}
	r.inputs = inputs
	r.si_list = make([]*ShoppingItem, 0)
	for _, v := range inputs {
		r.si_list = append(r.si_list,
			&ShoppingItem{(*comMap_p)[Barcode(v.Barcode)],
				v.Amount, 0, 0, 0})
	}
}

func PrintInvoice(r *Request) {
	var total float64
	var totalAllowance float64
	fmt.Println("***<没钱赚商店>购物清单***")
	for _, si := range r.si_list {
		CalcSubtotal(si, r)
		PrintOneShoppingItem(si)
		total += si.subtotal
		totalAllowance += si.allowance
	}
	fmt.Println("----------------------")
	ExtraPrint(r)
	fmt.Printf("总计：%.2f(元)\n", total)
	if totalAllowance > 0 {
		fmt.Printf("节省：%.2f(元)\n", totalAllowance)
	}
	fmt.Println("**********************")

}

func PrintOneShoppingItem(si *ShoppingItem) {
	fmt.Printf("名称：%v，数量：%v%v，单价：%.2f(元)，小计：%.2f(元)",
		si.commodity.name, si.amount, si.commodity.unit,
		si.commodity.price, si.subtotal)

	if si.promotions_flags|DISCOUNTPROMOTIONS ==
		si.promotions_flags {
		fmt.Printf("，节省%.2f(元)\n", si.allowance)
	} else {
		fmt.Println()
	}

}

func CalcSubtotal(si *ShoppingItem, r *Request) {
	si.subtotal = float64(si.amount) * si.commodity.price
	orig := si.subtotal
	si.DoPromotions(r)
	si.allowance = orig - si.subtotal
}

func ExtraPrint(r *Request) {
	if r.promotions_flags|THREEFORTWOPROMOTIONS ==
		r.promotions_flags {
		fmt.Println("买二赠一商品：")
		for _, si := range r.si_list {
			if si.promotions_flags|THREEFORTWOPROMOTIONS ==
				si.promotions_flags {
				fmt.Printf("名称：%v，数量：%v%v\n",
					si.commodity.name, si.allowance/si.commodity.price, si.commodity.unit)
			}
		}
		fmt.Println("----------------------")
	}
}
