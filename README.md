# shield-race
Code example to show race condition we bumped into while working on SHIELD.

To reproduce this issue, please run the following:

```
mkdir -p $GOPATH/src/github.com/starkandwayne 
cd !$ && git clone https://github.com/starkandwayne/shield-race.git && cd shield-race && go build -race
```

And then execute the resulting binary.
