package beaconchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	BeaconchainHoleksy = "https://holesky.beaconcha.in/api/v1/"
	BeaconchainMainnet = "https://beaconcha.in/api/v1/"
)

func GetBeaconchainData(nodeAddress common.Address, eth2Url string) (*Data, error) {
	data := &Data{}
	if err := getChainSpec(data, eth2Url); err != nil {
		return nil, errors.Join(errors.New("failed to get chain spec"), err)
	}

	if err := getSlotAndEpoch(data); err != nil {
		return nil, errors.Join(errors.New("failed to get slot and epoch"), err)
	}

	if err := getWithdrawalQueue(data); err != nil {
		return nil, errors.Join(errors.New("failed to get withdrawal queue"), err)
	}

	if err := getLatestWithdrawals(data); err != nil {
		return nil, errors.Join(errors.New("failed to get latest withdrawals"), err)
	}

	if err := getAllValidators(data, nodeAddress); err != nil {
		return nil, errors.Join(errors.New("failed to get all validators"), err)
	}

	if err := updateValidatorStatus(data, eth2Url); err != nil { 
		return nil, errors.Join(errors.New("failed to update validator status"), err)
	}

	return data, nil
}

func getAllValidators(data *Data, nodeAddress common.Address) error {
	baseUrl, err := getBeaconChainBaseUrl(data)
	if err != nil {
		return err
	}

	baseURL := baseUrl + "validator/eth1/" + nodeAddress.Hex()
	offset := 0

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	for {
		url := fmt.Sprintf("%s?offset=%d", baseURL, offset)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return errors.Join(errors.New("failed to create request"), err)
		}

		res, err := httpClient.Do(req)
		if err != nil {
			return errors.Join(errors.New("failed to send request"), err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.Join(errors.New("failed to read response body"), err)
		}

		var beaconchainError struct {
			Data   string `json:"data"`
			Status string `json:"status"`
		}
		if err := json.Unmarshal(body, &beaconchainError); err == nil {
			return errors.Join(errors.New("failed to decode error"), err)
		}

		if beaconchainError.Status != "OK" {
			return fmt.Errorf("failed to get beaconchain data %s: %s", beaconchainError.Status, beaconchainError.Data)
		}

		var beaconchainResponse struct {
			Data []struct {
				PublicKey      string `json:"publickey"`
				ValidSignature bool   `json:"valid_signature"`
				ValidatorIndex int    `json:"validatorindex"`
			} `json:"data"`
			Status string `json:"status"`
		}
		if err := json.Unmarshal(body, &beaconchainResponse); err != nil {
			return errors.Join(errors.New("failed to decode response"), err)
		}

		if len(beaconchainResponse.Data) == 0 {
			break
		}

		for _, validator := range beaconchainResponse.Data {
			minipool := Minipool{
				PublicKey: validator.PublicKey,
				Index:     validator.ValidatorIndex,
			}
			data.Minipools = append(data.Minipools, minipool)
		}

		// we did not hit the limit, so we are done
		if len(beaconchainResponse.Data) < 2000 {
			break
		}

		offset += 2000
	}

	return nil
}

func getSlotAndEpoch(data *Data) error {
	baseUrl, err := getBeaconChainBaseUrl(data)
	if err != nil {
		return err
	}

	url := baseUrl + "slot/latest"

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Join(errors.New("failed to create request"), err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return errors.Join(errors.New("failed to send request"), err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Join(errors.New("failed to read response body"), err)
	}

	var beaconchainResponse struct {
		Data struct {
			Epoch int `json:"epoch"`
			Slot  int `json:"slot"`
		} `json:"data"`
		Status string `json:"status"`
	}
	if err := json.Unmarshal(body, &beaconchainResponse); err != nil {
		return errors.Join(errors.New("failed to decode response"), err)
	}

	if beaconchainResponse.Status != "OK" {
		return fmt.Errorf("failed to get beaconchain data %s", beaconchainResponse.Status)
	}

	data.CurrentEpoch = uint64(beaconchainResponse.Data.Epoch)
	data.CurrentSlot = beaconchainResponse.Data.Slot
	return nil
}

func updateValidatorStatus(data *Data, eth2Url string) error {
	for i, minipool := range data.Minipools {
		if err := getValidatorStatus(eth2Url, &minipool, data.CurrentSlot); err != nil {
			msg := fmt.Errorf("failed to get validator status for %s", minipool.PublicKey)
			return errors.Join(msg, err)
		}
		data.Minipools[i] = minipool
	}
	return nil
}

func getValidatorStatus(baseUrl string, minipool *Minipool, slot int) error {
	url := fmt.Sprintf("%s/eth/v1/beacon/states/%d/validators/%d", baseUrl, slot, minipool.Index)

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Join(errors.New("failed to create request"), err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return errors.Join(errors.New("failed to send request"), err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Join(errors.New("failed to read response body"), err)
	}

	var errResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &errResponse); err == nil && errResponse.Code != 0 {
		return fmt.Errorf("failed to get validator status %d: %s", errResponse.Code, errResponse.Message)
	}

	var validatorStatus struct {
		ExecutionOptimistic bool `json:"execution_optimistic"`
		Finalized           bool `json:"finalized"`
		Data                struct {
			Index     string `json:"index"`
			Balance   string `json:"balance"`
			Status    string `json:"status"`
			Validator struct {
				Pubkey                     string `json:"pubkey"`
				WithdrawalCredentials      string `json:"withdrawal_credentials"`
				EffectiveBalance           string `json:"effective_balance"`
				Slashed                    bool   `json:"slashed"`
				ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
				ActivationEpoch            string `json:"activation_epoch"`
				ExitEpoch                  string `json:"exit_epoch"`
				WithdrawableEpoch          string `json:"withdrawable_epoch"`
			} `json:"validator"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &validatorStatus); err != nil {
		return errors.Join(errors.New("failed to decode response"), err)
	}

	minipool.WithdrawalCredentials = common.HexToHash(validatorStatus.Data.Validator.WithdrawalCredentials)
	minipool.WithdrawalAddress = common.HexToAddress(validatorStatus.Data.Validator.WithdrawalCredentials[26:])
	minipool.Status = validatorStatus.Data.Status
	minipool.ExitEpoch, err = strconv.ParseUint(validatorStatus.Data.Validator.ExitEpoch, 10, 64)
	if err != nil {
		return errors.Join(errors.New("failed to convert exit epoch"), err)
	}
	minipool.WithdrawableEpoch, err = strconv.ParseUint(validatorStatus.Data.Validator.WithdrawableEpoch, 10, 64)
	if err != nil {
		return errors.Join(errors.New("failed to convert withdrawable epoch"), err)
	}
	minipool.ActivationEpoch, err = strconv.ParseUint(validatorStatus.Data.Validator.ActivationEpoch, 10, 64)
	if err != nil {
		return errors.Join(errors.New("failed to convert activation epoch"), err)
	}

	// fmt.Println("url: ", url)
	// fmt.Println("Validator status:", validatorStatus)

	return nil
}

func getWithdrawalQueue(data *Data) error {
	baseUrl, err := getBeaconChainBaseUrl(data)
	if err != nil {
		return err
	}

	url := baseUrl + "validators/queue"

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Join(errors.New("failed to create request"), err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return errors.Join(errors.New("failed to send request"), err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Join(errors.New("failed to read response body"), err)
	}

	var beaconchainResponse struct {
		Data struct {
			Entering int `json:"beaconchain_entering"`
			Exiting  int `json:"beaconchain_exiting"`
			Count    int `json:"validatorscount"`
		} `json:"data"`
		Status string `json:"status"`
	}
	if err := json.Unmarshal(body, &beaconchainResponse); err != nil {
		return errors.Join(errors.New("failed to decode response"), err)
	}

	if beaconchainResponse.Status != "OK" {
		return fmt.Errorf("failed to get beaconchain data %s", beaconchainResponse.Status)
	}

	data.ValidatorsEntering = beaconchainResponse.Data.Entering
	data.ValidatorsExiting = beaconchainResponse.Data.Exiting
	data.ValidatorsCount = beaconchainResponse.Data.Count
	return nil
}

func getLatestWithdrawals(data *Data) error {
	baseUrl, err := getBeaconChainBaseUrl(data)
	if err != nil {
		return err
	}

	i := 0
	for {
		url := fmt.Sprintf("%sslot/%d/withdrawals", baseUrl, data.CurrentSlot-i)
		i--

		httpClient := &http.Client{
			Timeout: time.Second * 5,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return errors.Join(errors.New("failed to create request"), err)
		}

		res, err := httpClient.Do(req)
		if err != nil {
			return errors.Join(errors.New("failed to send request"), err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.Join(errors.New("failed to read response body"), err)
		}

		var beaconchainResponse struct {
			Data []struct {
				ValidatorIndex int `json:"validatorindex"`
			} `json:"data"`
			Status string `json:"status"`
		}
		if err := json.Unmarshal(body, &beaconchainResponse); err != nil {
			return errors.Join(errors.New("failed to decode response"), err)
		}

		if beaconchainResponse.Status != "OK" {
			return fmt.Errorf("failed to get beaconchain data %s", beaconchainResponse.Status)
		}
		time.Sleep(100 * time.Millisecond)

		if beaconchainResponse.Status == "OK" && len(beaconchainResponse.Data) > 0 {
			data.LastValidatorIndex = beaconchainResponse.Data[len(beaconchainResponse.Data)-1].ValidatorIndex
			data.WithdrawalPerSlot = len(beaconchainResponse.Data)
			break
		}
	}
	return nil
}

func getChainSpec(data *Data, eth2Url string) error {
	url := eth2Url + "/eth/v1/config/spec"

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Join(errors.New("failed to create request"), err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return errors.Join(errors.New("failed to send request"), err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Join(errors.New("failed to read response body"), err)
	}

	var errResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &errResponse); err == nil && errResponse.Code != 0 {
		return fmt.Errorf("failed to get chain spec %d: %s", errResponse.Code, errResponse.Message)
	}

	var beaconchainResponse struct {
		Data struct {
			DepositContractAddress          string `json:"DEPOSIT_CONTRACT_ADDRESS"`
			DepositNetworkID                string `json:"DEPOSIT_NETWORK_ID"`
			DomainAggregateAndProof         string `json:"DOMAIN_AGGREGATE_AND_PROOF"`
			InactivityPenaltyQuotient       string `json:"INACTIVITY_PENALTY_QUOTIENT"`
			InactivityPenaltyQuotientAltair string `json:"INACTIVITY_PENALTY_QUOTIENT_ALTAIR"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &beaconchainResponse); err != nil {
		return errors.Join(errors.New("failed to decode response"), err)
	}

	data.NetworkId, err = strconv.Atoi(beaconchainResponse.Data.DepositNetworkID)
	if err != nil {
		return errors.Join(errors.New("failed to convert network id"), err)
	}

	return nil
}

func getBeaconChainBaseUrl(data *Data) (string, error) {
	switch data.NetworkId {
	case 1:
		return BeaconchainMainnet, nil
	case 17000:
		return BeaconchainHoleksy, nil
	default:
		return "", errors.New("only mainnet and holesky are supported")
	}
}