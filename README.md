# Protonmail Reader

Right now this has enough logic hacked together from a broken project that it
will actually read the email out of a box and print the context of the latest
message. 

#### What is our goal?
We want to minimize the proton mail api as much as possible to create an
ultra-light version that will allow us to build ultra-light clients for very
specific purposes; allowing us great customization without much overhead
ideally.

#### Development
Lets start by merging down some of the files like manager; that are spread
across like 10 files unnecessarily. 

#### References:

https://github.com/johannesdrescher/go-protonmail **This is the project we
modified to get working**

https://github.com/emersion/hydroxide/blob/master/cmd/hydroxide/main.go
https://github.com/NeoHBz/protonmail_hydroxide/blob/master/cmd/hydroxide/main.go

https://github.com/ljanyst/peroxide/tree/master/pkg
