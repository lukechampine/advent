package main

const input = `The first floor contains a strontium generator, a strontium-compatible microchip, a plutonium generator, and a plutonium-compatible microchip.
The second floor contains a thulium generator, a ruthenium generator, a ruthenium-compatible microchip, a curium generator, and a curium-compatible microchip.
The third floor contains a thulium-compatible microchip.
The fourth floor contains nothing relevant.`

var part1 = `
F4 .  .  .  .  .  .  .  .  .  .  .
F3 .  .  .  .  .  .  TM .  .  .  .
F2 .  .  .  .  .  TG .  RG RM CG CM
F1 E  SG SM PG PM .  .  .  .  .  .   (0)

F4 .  .  .  .  .  .  .  .  .  .  .
F3 .  .  .  .  .  .  TM .  .  .  .
F2 E  SG .  PG .  TG .  RG RM CG CM
F1 .  .  SM .  PM .  .  .  .  .  .   (1)

F4 .  .  .  .  .  .  .  .  .  .  .
F3 .  .  .  .  .  .  TM .  .  .  .
F2 .  SG .  PG .  TG .  RG RM CG .
F1 E  .  SM .  PM .  .  .  .  .  CM  (2)

F4 .  .  .  .  .  .  .  .  .  .  .
F3 E  .  .  .  PM .  TM .  .  .  CM
F2 .  SG .  PG .  TG .  RG RM CG .
F1 .  .  SM .  .  .  .  .  .  .  .   (4)

F4 .  .  .  .  .  .  .  .  .  .  .
F3 .  .  .  .  PM .  TM .  .  .  .
F2 .  SG .  PG .  TG .  RG RM CG .
F1 E  .  SM .  .  .  .  .  .  .  CM  (6)

F4 E  SG SM .  .  .  .  .  .  .  .
F3 .  .  .  .  PM .  TM .  .  .  .
F2 .  .  .  PG .  TG .  RG RM CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (9)

F4 .  SG .  .  .  .  .  .  .  .  .
F3 .  .  SM .  .  .  TM .  .  .  .
F2 E  .  .  PG PM TG .  RG RM CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (11)

F4 E  SG .  PG PM .  .  .  .  .  .
F3 .  .  SM .  .  .  TM .  .  .  .
F2 .  .  .  .  .  TG .  RG RM CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (13)

F4 .  SG .  PG .  .  .  .  .  .  .
F3 .  .  SM .  PM .  .  .  .  .  .
F2 E  .  .  .  .  TG TM RG RM CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (15)

F4 .  SG .  PG .  .  .  .  .  .  .
F3 E  .  SM .  PM .  TM .  RM .  .
F2 .  .  .  .  .  TG .  RG .  CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (16)

F4 .  SG .  PG .  .  .  .  .  .  .
F3 .  .  SM .  PM .  .  .  RM .  .
F2 E  .  .  .  .  TG TM RG .  CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (17)

F4 E  SG .  PG .  TG TM .  .  .  .
F3 .  .  SM .  PM .  .  .  RM .  .
F2 .  .  .  .  .  .  .  RG .  CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (19)

F4 .  SG .  PG .  TG .  .  .  .  .
F3 .  .  SM .  PM .  TM .  .  .  .
F2 E  .  .  .  .  .  .  RG RM CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (21)

F4 .  SG .  PG .  TG .  .  .  .  .
F3 E  .  SM .  PM .  TM .  RM .  CM
F2 .  .  .  .  .  .  .  RG .  CG .
F1 .  .  .  .  .  .  .  .  .  .  .   (22)

F4 .  SG .  PG .  TG .  .  .  .  .
F3 .  .  SM .  PM .  TM .  .  .  CM
F2 E  .  .  .  .  .  .  RG RM CG .
F1 .  .  .  .  .  .  .  .  .  .  .   (23)

F4 E  SG .  PG .  TG .  RG RM .  .
F3 .  .  SM .  PM .  TM .  .  .  CM
F2 .  .  .  .  .  .  .  .  .  CG .
F1 .  .  .  .  .  .  .  .  .  .  .   (25)

F4 .  SG .  PG .  TG .  RG .  .  .
F3 .  .  SM .  PM .  TM .  RM .  .
F2 E  .  .  .  .  .  .  .  .  CG CM
F1 .  .  .  .  .  .  .  .  .  .  .   (27)

F4 E  SG .  PG .  TG .  RG .  CG CM
F3 .  .  SM .  PM .  TM .  RM .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (29)

F4 .  SG .  PG .  TG .  RG .  CG .
F3 E  .  SM .  PM .  TM .  RM .  CM
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (30)

F4 E  SG .  PG .  TG .  RG RM CG CM
F3 .  .  SM .  PM .  TM .  .  .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (31)

F4 .  SG .  PG .  TG .  RG .  CG CM
F3 E  .  SM .  PM .  TM .  RM .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (32)

F4 E  SG .  PG .  TG TM RG RM CG CM
F3 .  .  SM .  PM .  .  .  .  .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (33)

F4 .  SG .  PG .  TG .  RG RM CG CM
F3 E  .  SM .  PM .  TM .  .  .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (34)

F4 E  SG .  PG PM TG TM RG RM CG CM
F3 .  .  SM .  .  .  .  .  .  .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (35)

F4 .  SG .  PG .  TG TM RG RM CG CM
F3 E  .  SM .  PM .  .  .  .  .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (36)

F4 E  SG SM PG PM TG TM RG RM CG CM
F3 .  .  .  .  .  .  .  .  .  .  .
F2 .  .  .  .  .  .  .  .  .  .  .
F1 .  .  .  .  .  .  .  .  .  .  .   (37)
`

var part2 = `
F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 .  .  .  .  .  .  TM .  .  .  .  .  .  .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG SM PG PM .  .  .  .  .  .  EG EM DG DM  (0)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 E  SG SM .  .  .  TM .  .  .  .  .  .  .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  .  .  PG PM .  .  .  .  .  .  EG EM DG DM  (2)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 .  .  SM .  .  .  TM .  .  .  .  .  .  .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG .  PG PM .  .  .  .  .  .  EG EM DG DM  (4)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 E  .  SM PG PM .  TM .  .  .  .  .  .  .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  SG .  .  .  .  .  .  .  .  .  EG EM DG DM  (6)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  .  .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG .  PG .  .  .  .  .  .  .  EG EM DG DM  (8)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 E  .  SM .  PM .  TM .  .  .  .  EG EM .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  SG .  PG .  .  .  .  .  .  .  .  .  DG DM  (10)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  . 
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG .  PG .  .  .  .  .  .  .  EG .  DG DM  (12)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 E  .  SM .  PM .  TM .  .  .  .  .  EM DG DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  SG .  PG .  .  .  .  .  .  .  EG .  .  .   (14)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG .  PG .  .  .  .  .  .  .  EG .  DG .   (16)

F4 E  .  .  .  .  .  .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  SG .  PG .  .  .  .  .  .  .  .  .  .  .   (19)

F4 .  .  .  .  .  .  .  .  .  .  .  .  .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG .  PG .  .  .  .  .  .  .  EG .  .  .   (22)

F4 E  .  .  PG .  .  .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  SG .  .  .  .  .  .  .  .  .  .  .  .  .   (25)

F4 .  .  .  .  .  .  .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 E  SG .  PG .  .  .  .  .  .  .  .  .  .  .   (28)

F4 E  SG .  PG .  .  .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (31)

F4 .  .  .  PG .  .  .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 E  SG .  .  .  TG .  RG RM CG CM .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (33)

F4 E  SG .  PG .  TG .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  .  .  RG RM CG CM .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (35)

F4 .  .  .  PG .  TG .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  .  .  .  .  EM .  DM
F2 E  SG .  .  .  .  .  RG RM CG CM .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (37)

F4 .  .  .  PG .  TG .  .  .  .  .  EG .  DG . 
F3 E  .  SM .  PM .  TM RG RM .  .  .  EM .  DM
F2 .  SG .  .  .  .  .  .  .  CG CM .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (38)

F4 .  .  .  PG .  TG .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  RM .  .  .  EM .  DM
F2 E  SG .  .  .  .  .  RG .  CG CM .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (39)

F4 .  .  .  PG .  TG .  .  .  .  .  EG .  DG . 
F3 E  .  SM .  PM .  TM .  RM CG CM .  EM .  DM
F2 .  SG .  .  .  .  .  RG .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (40)

F4 .  .  .  PG .  TG .  .  .  .  .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  RM .  CM .  EM .  DM
F2 E  SG .  .  .  .  .  RG .  CG .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (41)

F4 E  SG .  PG .  TG .  .  .  CG .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  RG .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (43)

F4 .  .  .  PG .  TG .  .  .  CG .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  RM .  CM .  EM .  DM
F2 E  SG .  .  .  .  .  RG .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (45)

F4 E  SG .  PG .  TG .  RG .  CG .  EG .  DG . 
F3 .  .  SM .  PM .  TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (47)

F4 .  .  .  PG .  TG .  RG .  CG .  EG .  DG . 
F3 E  SG SM .  PM .  TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (48)

F4 E  SG SM PG .  TG .  RG .  CG .  EG .  DG . 
F3 .  .  .  .  PM .  TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (49)

F4 .  SG SM .  .  TG .  RG .  CG .  EG .  DG . 
F3 E  .  .  PG PM .  TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (50)

F4 E  SG SM PG PM TG .  RG .  CG .  EG .  DG . 
F3 .  .  .  .  .  .  TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (51)

F4 .  SG SM PG PM .  .  RG .  CG .  EG .  DG . 
F3 E  .  .  .  .  TG TM .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (52)

F4 E  SG SM PG PM TG TM RG .  CG .  EG .  DG . 
F3 .  .  .  .  .  .  .  .  RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (53)

F4 .  SG SM PG PM TG TM .  .  CG .  EG .  DG . 
F3 E  .  .  .  .  .  .  RG RM .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (54)

F4 E  SG SM PG PM TG TM RG RM CG .  EG .  DG . 
F3 .  .  .  .  .  .  .  .  .  .  CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (55)

F4 .  SG SM PG PM TG TM RG RM .  .  EG .  DG . 
F3 E  .  .  .  .  .  .  .  .  CG CM .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (56)

F4 E  SG SM PG PM TG TM RG RM CG CM EG .  DG . 
F3 .  .  .  .  .  .  .  .  .  .  .  .  EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (57)

F4 .  SG SM PG PM TG TM RG RM CG CM .  .  DG . 
F3 E  .  .  .  .  .  .  .  .  .  .  EG EM .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (58)

F4 E  SG SM PG PM TG TM RG RM CG CM EG EM DG . 
F3 .  .  .  .  .  .  .  .  .  .  .  .  .  .  DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (59)

F4 .  SG SM PG PM TG TM RG RM CG CM EG EM .  . 
F3 E  .  .  .  .  .  .  .  .  .  .  .  .  DG DM
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (60)

F4 E  SG SM PG PM TG TM RG RM CG CM EG EM DG DM
F3 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F2 .  .  .  .  .  .  .  .  .  .  .  .  .  .  . 
F1 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .   (61)
`
