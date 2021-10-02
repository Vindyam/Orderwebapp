package main

import "testing"

// This is initial unit test for illustration purpose, far from complete.
// TODO: add more tests.

func TestBuildQuery(t *testing.T) {
	filters := []string{"order_name ilike $3"}
	idx := 1

	q, w := buildQuery(filters, idx)

	ansQ := `with some_orders as (
				select id, order_name, created_at, customer_id
				from orders
				where order_name ilike $3 order by created_at desc
				limit $1 offset $2),
			some_order_items as (
				select A.id, order_name, created_at, customer_id, B.id as order_item_id, price_per_unit, quantity
				from some_orders as A left join order_items as B
				on A.id = B.order_id
			)
			select order_name, created_at, customer_id, sum(price_per_unit*quantity) as amount, sum(coalesce(price_per_unit*delivered_quantity,0)) as delivered_amount
			from some_order_items as A left join delivery as B
			on A.order_item_id = B.order_item_id
			group by order_name, created_at, customer_id
			order by created_at desc`
	if q != ansQ {
		t.Errorf("q is %s;%s", q, ansQ)
	}

	ansW := `where order_name ilike $3`
	if w != ansW {
		t.Errorf("w is %s; want %s", w, ansW)
	}
	t.Log(q, w)
}
