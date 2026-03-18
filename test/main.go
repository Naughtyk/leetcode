package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

// PaymentProcessor - интерфейс для обработки платежей
type PaymentProcessor interface {
	Process(amount float64) error
	Verify(amount float64) bool
}

// CreditCardProcessor - реализация интерфейса для кредитной карты
type CreditCardProcessor struct {
	limit float64
}

// Process обрабатывает платеж
func (c CreditCardProcessor) Process(amount float64) error {
	if amount > c.limit {
		return errors.New("credit card limit exceeded")
	}
	fmt.Printf("Process payment of $%.2f using CreditCard\n", amount)
	return nil
}

// Verify проверяет возможность обработки платежа
func (c CreditCardProcessor) Verify(amount float64) bool {
	return amount <= c.limit
}

// PayPalProcessor - реализация интерфейса для PayPal
type PayPalProcessor struct {
	balance float64
}

// Process обрабатывает платеж
func (p PayPalProcessor) Process(amount float64) error {
	if amount > p.balance {
		return errors.New("not enough balance in PayPal")
	}
	fmt.Printf("Processed payment of $%.2f using PayPal\n", amount)
	return nil
}

// Verify проверяет возможность обработки платежа
func (p PayPalProcessor) Verify(amount float64) bool {
	return amount <= p.balance
}

// ExecutePayment вызывает методы Process и Verify
func ExecutePayment(processor PaymentProcessor, amount float64) {
	if processor.Verify(amount) {
		err := processor.Process(amount)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Verification failed for amount:", amount)
	}
}

func main() {
	creditCard := CreditCardProcessor{limit: 100.0}
	payPal := PayPalProcessor{balance: 200.0}

	ExecutePayment(creditCard, 50.0)
	ExecutePayment(&creditCard, 50.0)
	ExecutePayment(&payPal, 150.0)
	ExecutePayment(payPal, 150.0)
}

// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: <url>, code: <statusCode>
// В случае ошибки выводит в отдельной строке url: <url>, code: <statusCode>
// Функция должна завершаться при отмене контекста.
// Доп. задание: реализовать ограничение количества одновременно запущенных горутин.
func fetchParallel(ctx context.Context, urls []string) {
	const concurrentLimit = 10
	httpClient := &http.Client{}
	sem := make(chan struct{}, concurrentLimit)
	wg := sync.WaitGroup{}
	mch := make(chan string)
	defer close(mch)
	go func() {
		for msg := range mch {
			fmt.Println(msg)
		}
	}()
	wg.Add(len(urls))
	for _, u := range urls {
		select {
		case <-ctx.Done():
			return
		case sem <- struct{}{}:
		}
		u := u
		go func() {
			defer wg.Add(-1)
			defer func() { <-sem }()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				fmt.Printf("http.NewRequestWithContext: %v\n", err)
				return
			}
			resp, err := httpClient.Do(req)
			if err != nil {
				fmt.Printf("http client Do: %v\n", err)
				return
			}

			defer resp.Body.Close()

			select {
			case mch <- fmt.Sprintf("url: %s, code: %d", u, resp.StatusCode):
			case <-ctx.Done():
				return
			}
		}()
	}
	wg.Wait()
}

// задача 1
/*
package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
}

func main() {
	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}
*/

// задача 2

/*
func FindMaxProblem() {
	var maxNum int
	for i := 1000; i > 0; i-- {
		go func() {
			if i%2 == 0 && i > maxNum {
				maxNum = i
			}
		}()
	}
	fmt.Printf("Maximum is %d", maxNum)
}
*/

// задача 3
/*
// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: <url>, code: <statusCode>
// В случае ошибки выводит в отдельной строке url: <url>, code: <statusCode>
// Функция должна завершаться при отмене контекста.
// Доп. задание: реализовать ограничение количества одновременно запущенных горутин.
func fetchParallel(ctx context.Context, urls []string) {

}
*/
