# submask
 Substitutes tokens in wordlists using Hashcat masks.

 ```
# get a list of probable tokens

$ cat tokens.lst
Keywrd

# then take your favorite wordlist
$ cat words.lst
TheGreat123
TheGreats123
Thefats123
Greaty12345**

# and sub matching masks with your token
$ cat test.lst | submask tokens.lst
TheKeywrd123
Keywrds123
Keywrd12345**

 ```
 Install
```
go install -v github.com/jakewnuk/submask@latest
```

See [maskcat](https://github.com/jakewnuk/maskcat), [matchmask](https://github.com/jakewnuk/matchmask), and [submask](https://github.com/jakewnuk/submask), and [920mPasswordMasks](https://github.com/jakewnuk/920mPasswordMasks)
