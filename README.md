# filegea

Easy file server. 

## Requirements

- Go version
    - 1.15.x
    - 1.14.x
    - 1.13.x

## Installation & usage

```shell
$ git clone https://github.com/WhaleMountain/filegea.git
$ cd filegea
$ go build
$ mkdir -p /opt/filegea/Data
$ mv filegea /opt/filegea
$ cp config.toml /opt/filegea
```

### Example Config

```toml
Port = "1270"
DataPath = "/opt/filegea/Data"
```

* Save the data in the path set in **DataPath**

### Start filegea

```shell
$ cd /opt/filegea
$ ./filegea
```
