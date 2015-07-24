Transmission JSON RPC API library
=================================

Implement all available methode, for details see the [documentation](https://trac.transmissionbt.com/browser/trunk/extras/rpc-spec.txt?rev=14463).


## Usage


```go
package main

import (
        "crypto/tls"
        "net/http"

        "github.com/kr/pretty"

        "github.com/odwrtw/transmission"
)

func main() {
        // Simple client
        conf := transmission.Config{
                Address: "http://localhost:9091/transmission/rpc",
        }
        t, err := transmission.New(conf)
        if err != nil {
                pretty.Println(err)
        }

        // With untrusted SSL
        tr := &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        httpClient := http.Client{Transport: tr}

        conf = transmission.Config{
                Address:    "http://localhost:9091/transmission/rpc",
                HTTPClient: &httpClient,
        }
        t, err = transmission.New(conf)
        if err != nil {
                pretty.Println(err)
        }

        // Get all torrents
        torrents, err := t.GetTorrents()
        if err != err {
                pretty.Println(err)
        }
        pretty.Println(torrents)

        // Add a torrent
        torrent, err := t.Add("http://cdimage.debian.org/debian-cd/8.1.0/amd64/bt-cd/debian-8.1.0-amd64-CD-1.iso.torrent")
        if err != nil {
                pretty.Println(err)
        }

        // Update is information
        torrent.Update()
        pretty.Println(torrent)

        // Remove it
        err = t.RemoveTorrents([]*transmission.Torrent{torrent}, true)
        if err != nil {
                pretty.Println(err)
        }

        // Get session informations
        t.Session.Update()
        pretty.Println(t.Session)

}
```
