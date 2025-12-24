package payments

import (
	"fmt"
)

type IPayment interface {
	pay(total []int)
}

type PayWithBankService struct {
}

func (pb PayWithBankService) pay(bill []int) {
	total := 0
	for x := 0; x < len(bill); x++ {
		total += bill[x]
	}

	str := fmt.Sprintf("Total pembayaran Rp. %d dan dilakukan dengan menggunakan metode pembayaran via Bank\n=======================\n", total)
	fmt.Print(str)
}

type PayWithQrisService struct {
}

func (pq PayWithQrisService) pay(bill []int) {
	total := 0
	for x := 0; x < len(bill); x++ {
		total += bill[x]
	}
	str := fmt.Sprintf("Total pembayaran Rp. %d dan dilakukan dengan menggunakan metode pembayaran via QRIS\n=======================\n", total)
	fmt.Print(str)
}

type PayWithDcelupService struct {
}

func (pq PayWithDcelupService) pay(bill []int) {
	for x := 0; x < len(bill); x++ {
		fmt.Println(bill[x])
	}
}

type UserService struct {
}

func (us *UserService) RegisterPay(payService IPayment, total []int) {
	(payService.pay(total))
}
