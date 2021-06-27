package handler

import "github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"

func getPromoMock() []promo.Promo {
	return []promo.Promo{
		{ID: 1, Name: "Free Rasp Pi", Description: "Each sale of a MacBook Pro comes with a free Raspberry Pi B", Formula: "1*{43N23P}=1*{234234}", Enabled: true},
		{ID: 2, Name: "Buy 2 get 1 free Google Home", Description: "Buy 3 Google Homes for the price of 2", Formula: "3*{120P90}=1*{120P90}", Enabled: true},
		{ID: 3, Name: "Alexa Spreaker 10% discount", Description: "Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa Speakers", Formula: "3*{A304SD}=0.1n*{A304SD}", Enabled: true},
	}
}
