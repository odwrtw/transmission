# Transmission RPC client

This a basic implementation of transmission rpc API.

## Features

In all the exemple I assume you have a transmission instance. Ex:

```
package main
import "gitlab.quimbo.fr/odwrtw/transmission-go"

func main() {
	// New transmission
	t := transmission.New("http://mytransmission.com/transmisson/rpc")
	// Or with auth
	tWithAuth := transmission.NewWithAuth("http://mytransmission.com/transmisson/rpc", "MyUser", "MyPassword")
}
```

### List all the torrents

```
list, err := t.GetList()
if err != nil {
	log.Panic(err)
}
```

### Remove torrents from transmission

```
ids := []int{1,2,3}
err := t.RemoveTorrents(ids)
if err != nil {
	log.Panic(err)
}
```

### Add a new torrent

```
torrent, err := t.AddTorrent("http://myfile.torrent")
if err != nil {
	switch err {
	case transmission.ErrDuplicateTorrent:
		log.Println("Torrent already added")
	default:
		log.Panic(err)
	}
} else {
	log.Printf("Torrent : %#v\n", torrent)
}
```
