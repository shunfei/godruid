package godruid

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Client struct {
    Url      string
    EndPoint string

    Response []byte

    Debug        bool
    LastRequest  string
    LastResponse string
}

func (c *Client) DoQuery(query Query) (err error) {
    if c.EndPoint == "" {
        c.EndPoint = "/druid/v2"
    }
    query.setup()
    var reqJson []byte
    if c.Debug {
        reqJson, err = json.MarshalIndent(query, "", "  ")
        c.EndPoint += "?pretty"

        c.LastRequest = string(reqJson)
    } else {
        reqJson, err = json.Marshal(query)
    }
    if err != nil {
        return err
    }

    resp, err := http.Post(c.Url+c.EndPoint, "application/json", bytes.NewBuffer(reqJson))
    if err != nil {
        return
    }
    defer func() {
        resp.Body.Close()
    }()

    rawBytes, _ := ioutil.ReadAll(resp.Body)
    if c.Debug {
        c.LastResponse = string(rawBytes)
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("%s: %s", resp.Status, string(rawBytes))
    }

    return query.onResponse(rawBytes)
}
