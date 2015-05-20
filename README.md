Transmission JSON RPC API library
=================================

Implement all available methode, for details see
https://trac.transmissionbt.com/browser/trunk/extras/rpc-spec.txt?rev=14463


Usage
-----

    package main

    import (
            "crypto/tls"
            "net/http"

            "github.com/kr/pretty"

            "gitlab.quimbo.fr/odwrtw/transmission-go"
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
            torrent, err := t.Add("http://cdimage.debian.org/debian-cd/8.0.0/amd64/bt-dvd/debian-8.0.0-amd64-DVD-2.iso.torrent")
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
