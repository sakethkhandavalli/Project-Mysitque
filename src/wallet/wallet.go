package wallet

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	UserID        string       `json:"userID"` //Primary Key
	UserName      string       `json:"userName"`
	WalletId      string       `json: "walletId"` //Secondary Key
	AssociateList []Associates `json:"associates, omitempty"`
}

type Associates struct {
	AssociateID       string `json:'associateID'`
	AssociateName     string `json:"associateName"`
	AssociateWalletId string `json:"associateWalletAddress"`
}
type WalletInfo struct {
	WalletId   string     `json:"walletId"`
	WalletKeys WalletKeys `json:"walletkeys"`
}

type BalanceInfo struct {
	Balance          float64 `json:"balance"`
	AvailableBalance float64 `json:"availableBalance"`
	LockedBalance    float64 `json:"lockedBalance"`
	UnderClearance   float64 `json:"underClearance"`
}

type Wallet struct {
	User            UserInfo       `json: "userInfo"`
	WalletInfo      WalletInfo     `json: "walletInfo"`
	Status          string         `json:"status"` //Active Frozen INACTIVE
	Balance         BalanceInfo    `json:"balance"`
	TransactionList []Transactions `json:"transactionList, omitempty"`
}

type WalletKeys struct {
	PublicKey  string `json:"publickey"`
	PrivateKey string `json:"privatekey"`
}
type Transactions struct {
	TransactionId       string     `json:"transactionId"`
	TimeOfTransaction   time.Time  `json:"time"`
	Amount              float64    `json:"amount"`
	NatureOfTransaction string     `json:"nature"` //Credited, debited, Credit Under Process, Debit Under Process etc.
	SourceAddress       WalletInfo `json:"sourceAddress"`
	DestinationAddress  WalletInfo `json:"destinationAddress"`
}

func RenderWalletPage(wallet *Wallet) string {
	var htmlPage string
	htmlPage = "<html><body>This is wallet page</body></html>"
	return htmlPage
}

var FetchWalletPage = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/wallet.html")
})
var FetchWalletPageWithUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userid"]
	for _, item := range wallet {
		if item.User.UserID == userId {
			json.NewEncoder(w).Encode(item)
		}
	}

})

var GetBalanceByUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userid"]
	for _, item := range wallet {
		if item.User.UserID == userId {
			json.NewEncoder(w).Encode(item.Balance)
		}
	}

})

var ListAssociatesWithUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userid"]
	for _, item := range wallet {
		if item.User.UserID == userId {
			json.NewEncoder(w).Encode(item.User.AssociateList)
			return
		}
	}

})

var AddAssociateWithUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userid"]

	for i := 0; i < len(wallet); i++ {
		if wallet[i].User.UserID == userId {
			var associate Associates
			_ = json.NewDecoder(r.Body).Decode(&associate)
			wallet[i].User.AssociateList = append(wallet[i].User.AssociateList, associate)
			fmt.Println(wallet[i].User.AssociateList)
			json.NewEncoder(w).Encode("SUCCESS")
			break
		}
	}

})

var GetLatestTransactionsByUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

})

var GetAllTransactionsByUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userid"]
	for _, item := range wallet {
		if item.User.UserID == userId {
			json.NewEncoder(w).Encode(&item.TransactionList)
			break
		}
	}

})

func updateBalance(index int, transaction *Transactions) {
	//wallet[index]
	//TODO update this index or whatever in databases
}

var AddTransactionByUserId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userid"]

	for i := 0; i < len(wallet); i++ {
		if wallet[i].User.UserID == userId {
			var transaction Transactions
			_ = json.NewDecoder(r.Body).Decode(&transaction)
			wallet[i].TransactionList = append(wallet[i].TransactionList, transaction)
			fmt.Println(wallet[i].TransactionList)
			json.NewEncoder(w).Encode("SUCCESS")
			updateBalance(i, &transaction)
			break
		}
	}

})

func GetWallet(userId string) *Wallet {

	wallet := Wallet{}
	return &wallet

}

var wallet []Wallet

func getHexId() string {
	return uuid.NewV4().String()
}

func populateAndAppendWallet(username string, userid string) {
	exwallet := Wallet{}

	exwallet.User.UserID = userid
	exwallet.User.UserName = username
	exwallet.User.WalletId = getHexId()

	exwallet.WalletInfo.WalletId = getHexId()
	exwallet.WalletInfo.WalletKeys.PublicKey = getHexId()
	exwallet.WalletInfo.WalletKeys.PrivateKey = getHexId()

	exwallet.Status = "ACTIVE"
	exwallet.Balance.Balance = rand.Float64() * 1000
	exwallet.Balance.AvailableBalance = rand.Float64() * 500
	exwallet.Balance.LockedBalance = rand.Float64() * 50
	exwallet.Balance.UnderClearance = exwallet.Balance.Balance - (exwallet.Balance.AvailableBalance + exwallet.Balance.LockedBalance)
	wallet = append(wallet, exwallet)
}

func populateAndAppendAssociates(userId string,id string, name string, walletId string) {
	for i := 0; i < len(wallet); i++ {
		if wallet[i].User.UserID == userId {
			var associate Associates
			associate.AssociateID = id
			associate.AssociateName = name
			associate.AssociateWalletId = walletId
			wallet[i].User.AssociateList = append(wallet[i].User.AssociateList, associate)
			fmt.Println(wallet[i].User.AssociateList)
			break
		}
	}
}
func populateAndAppendTransactionList(userId string) {
	for i := 0; i < len(wallet); i++ {
		if wallet[i].User.UserID == userId {
			var transaction Transactions
			//TODO: Populate the struct here.
			wallet[i].TransactionList = append(wallet[i].TransactionList, transaction)
			updateBalance(i, &transaction)
			break
		}
	}
}
func InitWallets() {
	populateAndAppendWallet("Ashwini Patankar", "1234567890")
	populateAndAppendWallet("Sid Chakravarty", "1234567891")
	populateAndAppendWallet("Anirudh", "1234567892")
	populateAndAppendWallet("Gulshan", "1234567893")
	populateAndAppendWallet("Preetham", "1234567894")
	populateAndAppendWallet("Sai Kamal", "1234567895")
	populateAndAppendWallet("Saketh", "1234567896")
	populateAndAppendTransactionList("1234567890")
	populateAndAppendAssociates("1234567890","123","kamal","456")
	populateAndAppendAssociates("1234567890","1231","preetham","1131")

	for _, item := range wallet {
		//fmt.Println(item)
		sample, _ := json.Marshal(item)
		fmt.Println(string(sample))
	}
	fmt.Println("total dummy record created: ")
	fmt.Println(len(wallet))
}

