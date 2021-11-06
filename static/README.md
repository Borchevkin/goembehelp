# Project name

Short overview

## Firmware Overview

Short overview

## Hardware requirements

Suitable boards

## Software requirements

This part assumes the development environment is being built on a clean install of Ubuntu 16.04 or Windows 10.

To build the project, you need [Rowley CrossWorks for ARM IDE](https://www.rowley.co.uk/arm/index.htm)
One will also need to install the following packages:

* Generic ARM CPU Support Package
* STMicroelectronics STM32 CPU Support Package
* CMSIS 5 CMSIS-Core(M) Support Package
* CMSIS 5 Support Package
* Generic ARM CPU Support Package

One can install them via the CrossWorks system menu: `Tools->Packages->Install Packages...`.

## Dependencies

Dependency list

## Build application

Clone or copy project from the repo.

After that, open solution in CrossWorks environment. It's placed at `./solution/crossworks/stm32_test.hzp`.

You can build it via `Build Solution` command.

## Document

In the `./docs` folder exist `doxygen_doc` file for generation Doxygen docs. For generating docs run `doxygen doxygen_doc`

## UI and serial port settings

## More info

* [The project wiki page](#)
