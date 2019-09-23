## Meetup Terraform provider

Welcome! This repo contains the source code and compilation/utilization instructions for the custom terrafor provider described in Billtrust's 09/23/2019 tech meetup. In here you can find a simple NodeJS server that supports the basic CRUD operations and a terraform provider that interacts with said service, along with a Dockerfile to be able compile and run the code.

### Build

A Dockerfile is provider. All further instructions assume the image was built using the tag `meetup-tf-provider`

```shell
docker build -t meetup-tf-provider -f Dockerfile .
```

### Server
The terraform provider communicates with a custom made server written in NodeJS.
To run it, execute the following command: 

```shell
docker run -p 3000:3000 meetup-tf-provider
```

### Provider

The provider requires Go to be installed to be able to compile it. Go is included in the Dockerfile, so to compile the provider, just run the following:

```shell
docker run -v $(pwd)/terraform-provider-btmascot/bin:/root/go/src/github.com/billtrust/meetup-terraform-provider/terraform-provider-btmascot/bin -it --entrypoint /bin/bash meetup-tf-provider /scripts/build-provider.sh
```

You'll have your compiled terraform provider in the `./terraform-provider-btmascot/bin` folder, which you'll have to copy to the proper folder following the instructions [here](https://www.terraform.io/docs/plugins/basics.html#installing-plugins)

### Contributing

To add fixes or features, just create a PR. No strong enforcement of standars is implemented as of yet, just keep it readable and simple.
