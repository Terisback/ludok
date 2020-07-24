## Ludok - Experiments, sketches... ***Proof of concept** <sup>definitely an excuse</sup>*

I tried to do smth like Love2D in Go

I couldn't handle SDL2 gfx package (cause of compilation of original that can be done on unix system I believe)  
So if you didn't cope as I am, you need to pass `static` tag to go compilator

```bash
go run -tags static github.com/Terisback/ludok/cmd/ludok
```

Good luck, brav ✋