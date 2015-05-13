
    package main

    import (
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

        // resp, err := t.Post("torrent-get")
        // if err != nil {
        // 	pretty.Println("==========")
        // 	pretty.Println(err)
        // 	pretty.Println("==========")
        // }

        // body, err := ioutil.ReadAll(resp.Body)
        // if err != nil {
        // 	pretty.Println(err)
        // }

        // r := transmission.Response{Arguments: &transmission.Torrents{}}
        // err = json.Unmarshal(body, &r)
        // if err != nil {
        // 	pretty.Println(err)
        // }
        // // fmt.Println(f)
        torrents, err := t.GetTorrents()
        if err != err {
                pretty.Println(err)
        }
        pretty.Println(torrents)

    }
