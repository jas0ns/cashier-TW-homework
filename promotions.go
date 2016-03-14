package cashier

type Promotioner interface {
	PromotionsHandler(*ShoppingItem, *Request)
}

type Promotions struct {
	promotioner Promotioner
	pre         *Promotions
	next        *Promotions
}

func (si *ShoppingItem) DoPromotions(r *Request) {
	p_cur := si.commodity.p_head
	for p_cur != nil {
		p_cur.promotioner.PromotionsHandler(si, r)
		p_cur = p_cur.next
	}
}

func (commodity *Commodity) AddPromotions(
	promotioner Promotioner) *Promotions {
	newp := &Promotions{promotioner, nil, nil}
	// first node
	if commodity.p_head == nil {
		commodity.p_head = newp
		commodity.p_tail = newp
	} else {
		newp.pre = commodity.p_tail
		commodity.p_tail.next = newp
		commodity.p_tail = newp
	}
	return newp
}

func (commodity *Commodity) DeletePromotions(p *Promotions) {
	if p.pre == nil {
		commodity.p_head = p.next
		p.next.pre = nil
		p.next = nil
		return
	}

	if p.next == nil {
		commodity.p_tail = p.pre
		p.pre.next = nil
		p.pre = nil
		return
	}

	p.pre.next = p.next
	p.next.pre = p.pre
	p.pre = nil
	p.next = nil
}

type DiscountPromotions struct {
	percent float32
}

func (dp *DiscountPromotions) PromotionsHandler(
	si *ShoppingItem, r *Request) {
	if si.promotions_flags&THREEFORTWOPROMOTIONS != 0 {
		return
	}

	si.subtotal *= float64(dp.percent)
	r.promotions_flags |= DISCOUNTPROMOTIONS
	si.promotions_flags |= DISCOUNTPROMOTIONS
}

type ThreeForTwoPromotions struct{}

func (fp *ThreeForTwoPromotions) PromotionsHandler(
	si *ShoppingItem, r *Request) {
	for c := si.amount; c > 2; c = c - 3 {
		si.subtotal -= si.commodity.price
	}
	r.promotions_flags |= THREEFORTWOPROMOTIONS
	si.promotions_flags |= THREEFORTWOPROMOTIONS
}
