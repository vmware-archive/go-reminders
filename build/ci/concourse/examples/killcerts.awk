BEGIN {
    nope=1
}

/#### Added by/ { nope=0 }

{ if (nope) {
     print
  }
}
