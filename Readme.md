Build
```shell
docker build -t meetup-tf-provider -f Dockerfile .
```

Run server
```shell
docker run -p 3000:3000 meetup-tf-provider
```

Compile provider
```shell
docker run -v $(pwd)/terraform-provider-btmascot/bin:/root/go/src/github.com/billtrust/meetup-terraform-provider/terraform-provider-btmascot/bin -it --entrypoint /bin/bash meetup-tf-provider /scripts/build-provider.sh
```

You'll have your compiled terraform provider in the `./terraform-provider-btmascot/bin` folder.
