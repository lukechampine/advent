|=  ms=(list @)
=<
:-  (roll (turn ms fuel) add)
  (roll (turn ms recfuel) add)
|%
++  fuel
  |=  m=@
  (sub (max (div m 3) 2) 2)
++  recfuel
  |=  m=@
  ?:  =(m 0)  0
  =.  m  (fuel m)
  (add m $)
--
