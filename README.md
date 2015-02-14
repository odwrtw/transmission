# Transmission RPC client

This a basic implementation of transmission rpc API.

## Features

In all the exemple I assume you have a transmission instance. Ex:

```
package main
import "github.com/gregdel/transmission"

func main() {
	// New transmission
	t := transmission.New()
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
		fmt.Println("Torrent already added")
	default:
		log.Panic(err)
	}
} else {
	fmt.Printf("Torrent : %#v\n", torrent)
}
```
