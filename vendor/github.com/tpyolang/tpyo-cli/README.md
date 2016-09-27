# tpyo-cli
:ok_hand: Mkae tpyos in yuor tarmniel

```console
$ go get github.com/tpyolang/tpyo-cli/...
```

## Empaxels

```console
$ echo 'Hello World!' | typo
Hlleo Wrold!
```

```console
$ xmllint --xpath 'string(//div[@id="mw-content-text"]/p[1])'     \
    <(curl -s https://en.wikipedia.org/wiki/Typographical_error)  \
    | typo                                                        \
    | fold -w 80 -s
A taphgyiarpocl eorrr (oetfn setnhreod to tpyo) is a mtiakse mdae in the tpynig
psroecs (scuh as a seplling mkstaie)[2] of pntried mtieraal. Hloiairlscty, tihs
rfrreeed to makitess in mnaaul type-sitnteg (tpygpoarhy). The trem inucelds
errros due to mnhacecial flriaue or silps of the hnad or finegr,[2] but
edcxules eorrrs of inorcagne, scuh as slplenig errros. Borfee the aarrvil of
ptirning, the "cipysot's maitske" or "sbrcail erorr" was the eluianqvet for
mairtsnpcus. Msot topys ilvonve spmile doituaplcin, osiiosmn, tostpsraoiinn, or
sttbsoiiutun of a slaml nbemur of catrrechas.
```
