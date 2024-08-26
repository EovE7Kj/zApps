# zukd
## preliminary
[install docker](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://docs.docker.com/engine/install/&ved=2ahUKEwi58d7I7fSHAxVjODQIHeFNALIQFnoECAgQAQ&usg=AOvVaw3oxUtu6GW_HNWz3ZCPMLU_)
## setup
1. clone repo: `git clone https://github.com/EovE7Kj/zukd`
2. copy desired BINARY to /zuk/img: `cp /path/to/binary ./zukd/img`
3. _optional:_ modify target image format (qcow2/vmdk/etc) [default = qcow2]
```sh 
vim ./zukd/Dockerfile
```
```docker 
ARG FORMAT="<target_format>"
```
3. build Docker image
```sh
cd zukd && docker build ./zukd -t zukd
```
5. build unikernel image:
```sh
mkdir ./bin && docker run --rm -v ./bin:/home/bin zukd
```
