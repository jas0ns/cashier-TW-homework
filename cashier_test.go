package cashier

import ()

func InitExample() (*map[Barcode]*Commodity, *string) {
	comMap := make(map[Barcode]*Commodity)
	comMap["ITEM000001"] = &Commodity{
		name:     "可口可乐",
		unit:     "瓶",
		price:    3,
		category: Category("饮料"),
	}

	comMap["ITEM000005"] = &Commodity{
		name:     "羽毛球",
		unit:     "个",
		price:    1,
		category: Category("球类"),
	}

	comMap["ITEM000003"] = &Commodity{
		name:     "苹果",
		unit:     "斤",
		price:    5.5,
		category: Category("水果"),
	}

	input := `
	{"barcode" : "ITEM000001", "amount" : 3}
	{"barcode" : "ITEM000005", "amount" : 5}
	{"barcode" : "ITEM000003", "amount" : 2}
	`

	return &comMap, &input
}

func Example1() {
	comMap, input := InitExample()

	threeForTwoPromotions := &ThreeForTwoPromotions{}
	(*comMap)["ITEM000001"].AddPromotions(threeForTwoPromotions)
	(*comMap)["ITEM000005"].AddPromotions(threeForTwoPromotions)

	var request Request
	request.InitRequest(input, comMap)
	PrintInvoice(&request)

	// Output:
	// ***<没钱赚商店>购物清单***
	// 名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：6.00(元)
	// 名称：羽毛球，数量：5个，单价：1.00(元)，小计：4.00(元)
	// 名称：苹果，数量：2斤，单价：5.50(元)，小计：11.00(元)
	// ----------------------
	// 买二赠一商品：
	// 名称：可口可乐，数量：1瓶
	// 名称：羽毛球，数量：1个
	// ----------------------
	// 总计：21.00(元)
	// 节省：4.00(元)
	// **********************
}

func Example2() {
	comMap, input := InitExample()

	var request Request
	request.InitRequest(input, comMap)
	PrintInvoice(&request)

	// Output:
	// ***<没钱赚商店>购物清单***
	// 名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：9.00(元)
	// 名称：羽毛球，数量：5个，单价：1.00(元)，小计：5.00(元)
	// 名称：苹果，数量：2斤，单价：5.50(元)，小计：11.00(元)
	// ----------------------
	// 总计：25.00(元)
	// **********************

}

func Example3() {
	comMap, input := InitExample()

	discountPromotions := &DiscountPromotions{0.95}
	(*comMap)["ITEM000003"].AddPromotions(discountPromotions)

	var request Request
	request.InitRequest(input, comMap)
	PrintInvoice(&request)

	// Output:
	// ***<没钱赚商店>购物清单***
	// 名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：9.00(元)
	// 名称：羽毛球，数量：5个，单价：1.00(元)，小计：5.00(元)
	// 名称：苹果，数量：2斤，单价：5.50(元)，小计：10.45(元)，节省0.55(元)
	// ----------------------
	// 总计：24.45(元)
	// 节省：0.55(元)
	// **********************
}

func Example4() {
	comMap, input := InitExample()

	*input = `
	{"barcode" : "ITEM000001", "amount" : 3}
	{"barcode" : "ITEM000005", "amount" : 6}
	{"barcode" : "ITEM000003", "amount" : 2}
	`

	threeForTwoPromotions := &ThreeForTwoPromotions{}
	(*comMap)["ITEM000001"].AddPromotions(threeForTwoPromotions)
	(*comMap)["ITEM000005"].AddPromotions(threeForTwoPromotions)

	discountPromotions := &DiscountPromotions{0.95}
	(*comMap)["ITEM000003"].AddPromotions(discountPromotions)

	var request Request
	request.InitRequest(input, comMap)
	PrintInvoice(&request)

	// Output:
	// ***<没钱赚商店>购物清单***
	// 名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：6.00(元)
	// 名称：羽毛球，数量：6个，单价：1.00(元)，小计：4.00(元)
	// 名称：苹果，数量：2斤，单价：5.50(元)，小计：10.45(元)，节省0.55(元)
	// ----------------------
	// 买二赠一商品：
	// 名称：可口可乐，数量：1瓶
	// 名称：羽毛球，数量：2个
	// ----------------------
	// 总计：20.45(元)
	// 节省：5.55(元)
	// **********************
}
