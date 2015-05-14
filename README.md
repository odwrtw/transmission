
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

            torrents, err := t.GetTorrents()
            if err != err {
                    pretty.Println(err)
            }
            pretty.Println(torrents)

            torrent, err := t.AddTorrent("http://cdimage.debian.org/debian-cd/8.0.0/amd64/bt-dvd/debian-8.0.0-amd64-DVD-2.iso.torrent", "")
            if err != nil {
                    pretty.Println(err)
            }
            pretty.Println(torrent)

            err = t.RemoveTorrents([]*transmission.Torrent{torrent}, true)
            if err != nil {
                    pretty.Println(err)
            }

    }
