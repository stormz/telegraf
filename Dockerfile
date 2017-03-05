FROM golang:1.8

RUN go get github.com/influxdata/telegraf

WORKDIR $GOPATH/src/github.com/influxdata/telegraf

RUN git remote add stormz https://github.com/stormz/telegraf \
		&& git fetch stormz \
	  && git checkout deploy

RUN make

CMD ["telegraf"]
