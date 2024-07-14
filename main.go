package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "bytes"
)

type BlockHeader struct {
    RawData struct {
        Number    int64 `json:"number"`
        Timestamp int64 `json:"timestamp"`
    } `json:"raw_data"`
}

type Block struct {
    BlockHeader BlockHeader `json:"block_header"`
}

func getLatestBlock() (*Block, error) {
    resp, err := http.Get("https://api.trongrid.io/wallet/getnowblock")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var block Block
    if err := json.NewDecoder(resp.Body).Decode(&block); err != nil {
        return nil, err
    }
    return &block, nil
}

func getBlockByNumber(blockNumber int64) (*Block, error) {
    url := "https://api.trongrid.io/wallet/getblockbynum"
    requestBody, err := json.Marshal(map[string]int64{"num": blockNumber})
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var block Block
    if err := json.NewDecoder(resp.Body).Decode(&block); err != nil {
        return nil, err
    }
    return &block, nil
}

func calculateAverageBlockTime(numBlocks int) (float64, error) {
    latestBlock, err := getLatestBlock()
    if err != nil {
        return 0, err
    }
    latestBlockNumber := latestBlock.BlockHeader.RawData.Number

    var blockTimes []int64
    for i := latestBlockNumber - int64(numBlocks); i < latestBlockNumber; i++ {
        block, err := getBlockByNumber(i)
        if err != nil {
            return 0, err
        }
        blockTime := block.BlockHeader.RawData.Timestamp
        blockTimes = append(blockTimes, blockTime)
    }

    var blockIntervals []int64
    for i := 0; i < len(blockTimes)-1; i++ {
        interval := blockTimes[i+1] - blockTimes[i]
        blockIntervals = append(blockIntervals, interval)
    }

    var sumIntervals int64
    for _, interval := range blockIntervals {
        sumIntervals += interval
    }

    averageBlockTime := float64(sumIntervals) / float64(len(blockIntervals)) / 1000
    return averageBlockTime, nil
}

func main() {
    averageBlockTime, err := calculateAverageBlockTime(10)
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
        return
    }
    fmt.Printf("Среднее время блока: %.2f секунд\n", averageBlockTime)
}

