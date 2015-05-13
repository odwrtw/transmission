
    package main

    import (
            "encoding/json"
            "io/ioutil"

            "github.com/kr/pretty"

            "gitlab.quimbo.fr/odwrtw/transmission-go"
    )

    func main() {
            conf := transmission.Config{}
            t, err := transmission.New(conf)
            if err != nil {
                    pretty.Println(err)
            }

            pretty.Println(t)

            resp, err := t.Post("torrent-get")
            if err != nil {
                    pretty.Println("==========")
                    pretty.Println(err)
                    pretty.Println("==========")
            }

            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                    pretty.Println(err)
            }
            var f interface{}
            err = json.Unmarshal(body, &f)
            if err != nil {
                    pretty.Println(err)
            }

            pretty.Println(f)

    }
