package accounts

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	ERC20 "test/ERC20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func createClient() *ethclient.Client {
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func AccountBalancesCoinbase() {
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	if err != nil {
		log.Fatal(err)
	}
	AddressWallet := common.HexToAddress("0x3F85FfF902356Bf1594B1dA9f950097eD99a838b")
	balance, err := client.BalanceAt(context.Background(), AddressWallet, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
}

func AccountTokenBalances() {
	client := createClient()

	// Golem (GNT) Address

	tokenAddress := common.HexToAddress("0xbab6d3E8b9cC5CcAF8dcce8E6b28E792cF1ac175")

	AddressWallet := common.HexToAddress("0x3F85FfF902356Bf1594B1dA9f950097eD99a838b")
	instance, err := ERC20.NewERC20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	bal, err := instance.BalanceOf(&bind.CallOpts{}, AddressWallet)
	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value)
}

/*"Khối" đề cập đến dữ liệu và trạng thái được lưu trữ trong các nhóm liên tiếp được gọi là "khối".
Nếu bạn gửi ETH cho người khác, dữ liệu giao dịch cần được thêm vào một khối để thành công.*/

/*"Chuỗi" đề cập đến thực tế là mỗi khối liên quan đến mật mã của cha mẹ.
Nói cách khác, các khối bị xích lại với nhau.
Dữ liệu trong một khối không thể thay đổi mà không thay đổi tất cả các khối tiếp theo,
điều này sẽ yêu cầu sự đồng thuận của toàn bộ mạng.
Mỗi máy tính trong mạng phải đồng ý với từng khối mới và toàn bộ chuỗi.
Những máy tính này được gọi là "nút".
Các nút đảm bảo tất cả mọi người tương tác với blockchain có cùng một dữ liệu.
Để thực hiện thỏa thuận phân tán này, các blockchain cần một cơ chế đồng thuận.*/

/*Ethereum sử dụng cơ chế đồng thuận dựa trên bằng chứng. \
Bất cứ ai muốn thêm các khối mới vào chuỗi đều phải đặt cược ETH
- loại tiền gốc trong Ethereum -
làm phần mềm Validator tài sản thế chấp và chạy.
Các "trình xác nhận" này sau đó có thể được chọn ngẫu nhiên để đề xuất các khối mà các trình xác nhận khác kiểm tra
và thêm vào blockchain. Có một hệ thống phần thưởng và hình phạt khuyến khích người tham gia trung thực
và có sẵn trực tuyến càng nhiều càng tốt.*/

/*Nếu bạn muốn xem dữ liệu blockchain được băm như thế nào và sau đó được thêm vào lịch sử của các tài liệu tham khảo khối,
hãy chắc chắn xem bản demo này (mở trong một tab mới) của Anders Brownworth và xem video đi kèm bên dưới.*/
