package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seatingCap = 10
var arrivalRate = 100
var cutDur = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false
		fmt.Println(barber, " goes to waiting room to check for clients")
		for {
			// if no client barber sleeps
			if len(shop.ClientChan) == 0 {
				fmt.Println("No client so sleeping")
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientChan
			if shopOpen {
				if isSleeping {
					fmt.Printf("%s wakes %s up\n", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed
				shop.sendBarberHome(barber)
			}

		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s hair\n", barber, client)
	time.Sleep(shop.HairCutDuration)
	fmt.Printf("%s finished cutting %s hair\n", barber, client)
}
func (shop *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is going home\n", barber)
	shop.BarberDoneChan <- true
}
func (shop *BarberShop) closeShopForDay() {
	fmt.Print("Closing shop\n")
	close(shop.ClientChan)
	shop.Open = false
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarberDoneChan
	}
	close(shop.BarberDoneChan)
	fmt.Print("Closed and client gone home\n")
}
func (shop *BarberShop) addClient(client string) {
	if shop.Open {
		select {
		case shop.ClientChan <- client:
			fmt.Printf("adding %s to waitingroom\n", client)
		default:
			fmt.Printf("waiting room full so %s leaves\n", client)
		}
	} else {
		fmt.Printf("Shop is closed so %s leaves\n", client)
	}
}

func main() {
	// create channel
	clientChan := make(chan string, seatingCap)
	doneChan := make(chan bool)
	// create barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCap,
		HairCutDuration: cutDur,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}

	// add barber
	shop.addBarber("frank")
	shop.addBarber("gerard")
	shop.addBarber("meghan")
	shop.addBarber("billy")
	shop.addBarber("tony")
	// start barbershop
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()
	// add client
	i := 1
	go func() {

		for {
			randomMs := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMs)):
				shop.addClient(fmt.Sprintf("Client-%d", i))
				i++

			}
		}
	}()
	// wait for barbershop to close
	<-closed
}
