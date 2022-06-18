#!/bin/bash

set -ev

$GOPATH/bin/godepgraph -horizontal -s -o github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet | dot -Tpng -o depgraph.png