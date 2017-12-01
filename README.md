# fluent-bit-stdout
Stdout Output Plugin for FluentBit

Info on FluentBit: https://github.com/fluent/fluent-bit

Info on Golang Output Plugins for FluentBit: https://github.com/fluent/fluent-bit-go

### Run locally for development:


Download [Fluent-bit](http://fluentbit.io/download/) locally. [Build Fluent-bit](http://fluentbit.io/documentation/0.12/installation/build_install.html). Test Fluent-bit is installed and built properly by running: `bin/fluent-bit -i random -o stdout` from your Fluent-bit build directory. This command starts Fluent-bit, tells it to use "random" for the input plugin and "stdout" for the output plugin.

Now we want to run Fluent-bit with our stdout output plugin. `bin/fluent-bit -v -e <path to ../fluent-bit-stdout/out_stdout.so> -i random -o rainbow_stdout`.
