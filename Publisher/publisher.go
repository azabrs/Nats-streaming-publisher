package publisher

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)


type Publisher struct{
	Client_ID string
	Cluster_ID string
	Server_URL string
	Channel_name string
	Size int
	Delay time.Duration
}

func New(Client_ID, Cluster_ID, Channel_name string, Size, Server_port, Delay int) Publisher{
	return Publisher{
		Client_ID: Client_ID,
		Cluster_ID: Cluster_ID,
		Channel_name: Channel_name,
		Server_URL: fmt.Sprintf("nats://localhost:%d", Server_port),
		Size: Size,
		Delay: time.Second * time.Duration(Delay),
	}

}

func (pub Publisher) Publish_data() (int, error){
	sc, err := stan.Connect(pub.Cluster_ID, pub.Client_ID, stan.NatsURL(pub.Server_URL))
	if err != nil{
		log.Fatal(err)
	}
	for i := 0; i < pub.Size; i++{
		time.Sleep(pub.Delay)
		byte_data, err := get_byte_data()
		if err != nil{
			return i, err
		}
		if err = sc.Publish(pub.Channel_name, byte_data); err != nil{
			return i, err
		}
		log.Printf("%d-st message was published", i)
	}
	log.Println("All messages have been published successfully")
	return pub.Size, nil
}

func get_struct_data() order{
	cities := []string{"Moscow", "Volgograd", "Zelenograd", "Kazan", "Ekaterinburg"}
    charset := "abcdefghijklmnopqrstuvwxyz"
	charset_b := "QWERTYUIOPASDFGHJKLZXCVBNM"
	integer_num := "123456789"

	OrderUID := make([]byte, 19) 
	Rid := make([]byte, 19) 
	Track_number := make([]byte, 14) 

	City := cities[rand.Intn(len(cities))]
	for i := 0; i < 19; i++{
		b := rand.Intn(2)
		if b == 0{
			OrderUID[i] = charset[rand.Intn(len(charset))]
		}else{
			OrderUID[i] = integer_num[rand.Intn(len(integer_num))]
		}
	}
	for i := 0; i < 19; i++{
		b := rand.Intn(2)
		if b == 0{
			Rid[i] = charset[rand.Intn(len(charset))]
		}else{
			Rid[i] = integer_num[rand.Intn(len(integer_num))]
		}
	}
	for i := 0; i < 14; i++{
		Track_number[i] = charset_b[rand.Intn(len(charset_b))]
	}
	Entry := Track_number[:4]

	Phone := "+"
	for i := 0; i < 10; i++{
		Phone += strconv.Itoa((rand.Intn(10))) 
	}

	Zip := ""
	for i := 0; i < 7; i++{
		Zip += strconv.Itoa((rand.Intn(10))) 
	}
	var l [9][]byte
	for i := 0; i < 9; i++{
		for j := 0; j < 15; j++{
			l[i] = append(l[i], charset[rand.Intn(len(charset))])
		}
	}
	Name := "+" + string(l[0]) + " " + string(l[0]) + "ov"
	Addres := l[1]
	Region := l[2]
	Email := string(l[3]) + "@gmail.com"
	Provider := l[4]
	Bank := l[5]
	Name_items := l[6]
	Brand := l[7]
	DeliveryServ := l[8]
	Price := rand.Intn(1000)
	Sale := rand.Intn(80)

	return order{
		OrderUID: string(OrderUID),
		TrackNumber: string(Track_number),
		Entry: string(Entry),
		Delivery: Delivery{
			Name : Name,
			Phone :Phone ,
			Zip : Zip,
			City : City,
			Address : string(Addres),
			Region : string(Region),
			Email : Email,
		},
		Payment: Payment{
			Transaction : string(OrderUID),
			RequestID : "",
			Currency : "USD",
			Provider : string(Provider),
			Amount : rand.Intn(10000),
			PaymentDt : rand.Intn(10000),
			Bank : string(Bank),
			DeliveryCost : rand.Intn(1000),
			GoodsTotal : rand.Intn(100),
			CustomFee : 0,
		},
		Items: []Items{{
			ChrtID : rand.Intn(10000),
			TrackNumber : string(Track_number),
			Price : Price,
			Rid : string(Rid),
			Name : string(Name_items),
			Sale : Sale,
			Size : strconv.Itoa(rand.Intn(5)),
			TotalPrice : Price * Sale / 100,
			NmID : rand.Intn(100000),
			Brand : string(Brand),
			Status : rand.Intn(300),
		}},
		Locale : "en",
		InternalSignature: "",
		CustomerID: "test",
		DeliveryService: string(DeliveryServ),
		Shardkey : strconv.Itoa(rand.Intn(20)),
		SmID : rand.Intn(100),
		DateCreated : randate(),
		OofShard : "1",
	}
}

func  get_byte_data() ([]byte, error){
	data := get_struct_data()
	res, err := json.Marshal(data)
	if err != nil{
		return nil, err
	}
	return res, nil
}



func randate() time.Time {
    min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
    max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
    delta := max - min

    sec := rand.Int63n(delta) + min
    return time.Unix(sec, 0)
}