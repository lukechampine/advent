package main

import (
	"fmt"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `jlbcwrl (93)
fzqsahw (256) -> lybovx, pdmhva
rxivjo (206) -> mewof, hrncqs, qgfstpq
jhldwxy (26)
bpzxjg (62)
lahain (70)
enbnfw (39)
uzsytm (45)
gmcsy (16)
rsqyvy (99)
lsbsfge (163) -> ldxgz, mksan
husmkc (29)
ootidjt (63)
pjhry (38)
kvlbq (22)
rdrad (6) -> gwyfm, fozyip, uotzz, fmkkz
oqbfkud (470) -> rnbqhk, mepez, mnksdxf, mjsck, bbfaxid, nglea
pgchejz (54) -> ifelr, rdkvtq
zzjyw (91)
yftjdo (12)
eqnvty (87)
adbolgz (38)
rcjqzp (65) -> mkyam, apdfe, avzljw, hxgidsw, fkgxak, wzsbsf, woczl
ksrsmjx (72)
wfdlusw (49)
rpoep (38)
jesiypo (78)
jjxvns (56)
syyfs (35)
otplae (91)
epsjj (17)
utgxfsh (959) -> mupbrv, borbd, jmieet
pxzdv (15)
ksybvgt (213)
tywzhc (243)
sbdja (16)
ctynr (63)
vwbjuvx (99)
aidknm (97)
qlgme (21) -> ehjnzn, cdbkci
zetvslt (99)
ferzy (65)
dssdiok (97)
gexdb (6)
bbzmsv (87)
mepez (126) -> uqvyau, witovp
pubtsp (70) -> jlvwibm, uvvdw, okmqiy
gjcxbx (44)
mogwb (84)
qwiekvk (65)
kwbovke (74)
offjbb (15) -> dvoja, jteju, wuybunc, qzpzi
cwdmlj (86)
ojpok (88)
lytcy (2662) -> bkmvwp, uyrwi
antcinm (178) -> dmthn, ycacj, wkhggv
kstvq (69)
ibysnvc (79)
xulwh (71)
zgqzrc (459) -> wlajxb, mfywm, jqmlxr
uyrwi (11)
qirjqsm (96)
fnoteku (2482) -> mbezs, kcuygjx, bymkeq, opsqjx
leeyi (88)
pcodkpi (95)
tetdx (224) -> nshbwn, rpoep, fbqpk
ajhctd (18)
yhzuh (72)
dmvjz (39)
zdgwmi (24)
vprkzya (37)
ipryqov (24)
pdxylvu (86) -> etscle, bqtawf
ehrqn (23) -> njabgq, fyzeso, jrgwnfh, fxasat
ekszr (148) -> gnmydtk, wchxl
izkaqg (26)
lovypfn (53)
sztqzfl (98)
ckwopo (43) -> yurfpj, bgrxlxe, tohrp
muksdck (48) -> gwtcgyo, tfpuqgs
tlfsn (21)
hrvztx (57)
psulm (1838) -> rhbouej, urhlfju, obfet
dbufjl (95) -> faihy, oyamw
ucfhxsv (65)
ietfhkq (31)
psvjtka (29)
wzxei (51)
swurfm (64)
oybwhs (18)
dmdiy (1601) -> lazjlns, ygvol, rljjjo, whnjtp, jilomb
jteju (54)
rdnpoms (177) -> eskdbe, fbhidp, xtkxwd
rdhwx (62)
hxgidsw (332) -> fukgfu, skkatyg
pcewixh (109) -> iekurr, xspxxb
gsiypfp (1146) -> bhwca, qhcolun
igpabio (53) -> mlqxueo, lhsncv
vdjmhb (39)
pwdpe (42) -> leeyi, rhlpt, dtexqt, skpcerk
ciejakr (43) -> hcqnhm, anmeg, melsryt
yehxck (2391) -> boygd, kayqqz, iajslqp
sofve (139)
asifeu (278) -> rtsajcq, dcouu
mgqfa (75) -> cipgxee, jscjw
kbppnh (99)
apwsmv (31)
dzjyzd (39)
uobgj (488) -> akidv, sofve, wblhx
qngfv (8)
kledju (95)
besihsl (86)
zqnmsyb (73)
csqfv (14)
ubgrma (1059) -> ymhxg, yvxuom, aeyykn
ufyqf (77)
llventw (308) -> lsbsfge, itxycu, nddymc
zfhwfsw (53)
kigdvl (31)
fiufzkb (1194) -> ysigfem, bchfd, hgsmlsi
vaubjz (233) -> erfnrn, gqzva, goxfpk
yhpiqn (99)
wzsbsf (222) -> mwztduj, hkpvwpa
mjaol (281) -> dnazkh, jkamtqv
mufnw (106) -> yxdld, obkhwze, nkssh
mhapqqy (16)
brztp (27)
ebmjz (68) -> xfydt, eqnvty
hkjwio (322) -> hdzxuz, zdgwmi, ipryqov
eszxg (18)
qwzac (908) -> uiioczf, qjdpk, ylpuaf
ndsub (75)
xcqwjrm (63)
glsrg (74) -> maiimn, ufyqf
mtcxdod (80)
ygmhxm (129) -> pljyyn, njdmj
ijcojo (1042) -> dxboqki, ikplh, pubtsp, omergh
urhlfju (249) -> csqfv, rnddnj
lgefkp (17)
bmmmp (90)
rjzrh (360) -> hbkujf, mzwbtsa, oywob, dmxhbe
khiom (117) -> hpaggeh, lqumiws
zlgpmfs (143) -> ilxfa, nhpcp
fozyip (293) -> kvlbq, pfqbht
ylpuaf (64) -> mdzkkjf, tfdbdq, kiauwpn
xekimqs (65)
bekxoxk (87)
zybeuh (197) -> wonxzkm, jzuvrtp
pudyrbw (76)
bcyarwn (65)
saowgr (367) -> gbnxgny, krmphks, yftjdo, zmpwz
tgmle (73) -> prcjhey, thzwwh, cxhipq, tgvpi
ezxsv (90) -> qzqyveb, dfmlo
rayez (17)
ujjpxe (56)
efsrz (93)
xaifomj (53)
ayury (23)
zavef (69)
qonfiro (190) -> cotvk, evqkdq
puurcu (1689) -> awagc, ajhctd
omergh (208)
padxpdx (192) -> psvjtka, husmkc
cxhipq (92)
jhgsgy (39)
leyiz (74) -> fvfek, njrfnvt
kdcvwyf (52)
zyympz (887) -> pxteg, amnoi, amzwex
jbbmcxg (34)
uvmqarq (751) -> muyodod, nclwgga, oeqvt
duepwu (683) -> hbueku, zbcra, yxtjx
aatagmv (44)
zsvuw (11)
fynniwm (35)
fjzpjk (88) -> uvsny, aatagmv
rulhhsl (90)
fcscdg (276) -> twvib, skjtvyz, oybwhs, rdmggka
vwotlbl (61)
ijyoy (24)
jpenh (186) -> kdcvwyf, rjxvbkb
qzpzi (54)
nshbwn (38)
foyfb (50)
kbyot (337) -> jhldwxy, izkaqg
bhxdhmr (65)
netsm (53)
tgffx (75) -> kstvq, cjmfqbo
msthql (7)
hgscccv (62)
dbnelj (212) -> jzcmf, sqiac, ijyoy, jjqguew
jxfqev (99)
elmog (32)
ygvol (202) -> mszzq, tzzdbi
gjrqs (159) -> iprkb, cgouh
rabxkov (84)
wenii (79) -> qsrpqe, zdhqhh, jpenh, hwrxn, vtvnn, mpgixa, fbjbwv
jkqvg (28)
kzdugfh (90)
mhkcp (17)
tfhdccw (93)
qzakfh (62)
hrncqs (12)
tzmndkd (221) -> lixsvp, ofyxhb
cjmfqbo (69)
ikplh (98) -> bqifoq, pjedtl
berje (87)
ikbqz (9)
avzljw (234) -> vdbihjp, zavef
ibiwh (26) -> ndsub, moihbcu
vdvadz (38)
yirnxk (158) -> lgefkp, rayez
vbnlfuo (169) -> pppxrv, rdhwx
lgxjr (238) -> mhkba, bsrkr
ynayo (71)
uvvdw (46)
udkyxw (48)
zotwsb (170) -> wlufwnr, frksv
tohdgsa (30)
bqtawf (45)
wrfxdc (25)
vjxmbzm (69)
opmmmd (32) -> pcodkpi, xhonf
hkpvwpa (75)
uflldd (39)
oelrlp (23)
lptkbzc (151) -> unvje, bzsiwp
ecdieu (23)
pxhwecw (57)
ryulu (61)
uplweb (127) -> bpzxjg, utivde
wblhx (97) -> xayglgm, pddllsa
grcox (91)
xergqq (99)
tgujvov (59) -> rhpco, ojpok, trbnpf
sdnkg (1381) -> gyutarg, gcwcs, mfjeyx
oydxsh (67)
pxihdrd (50)
cizehi (7)
zhopeqm (80)
frksv (89)
qvbfob (266) -> oejqkp, ocgkcxp
ldfsxw (17)
wltpv (260) -> nlndxah, etyyw
pddteg (52) -> lwwyx, mhnlgyl, mvfhc, dzggd, opqoq, mufrkl, inghbu
kybsigz (80)
hgsmlsi (31) -> bqqwmry, lqavsk
cbyvzkp (99)
tjpatnk (44)
srneoo (11) -> dhlrq, ivwcuc, laxoeu
piouhkc (95)
tgfgem (33)
egrfvic (49)
jmieet (137) -> ckwooig, stkodp
cldbz (15)
gylsj (52)
ecoqyk (35)
pnhibed (75) -> gmdsb, gijtd
fksxl (5)
rhpco (88)
eklahxm (51)
ftzht (8102) -> dmdiy, sfrbkzf, hlcnxe, zwsrt
cykvxs (84)
ccckie (201) -> tgkusb, alztgi
hbueku (10) -> ohcszyk, szutvca
cztikzk (174) -> wxdpcvv, lpbayb
lprmau (27) -> rqymw, dssdiok, dydwqz, eyale
zorvynv (176) -> mqybmn, jaxidl
laxoeu (88)
nvvxl (93)
duophdw (72)
qjwfsfk (11)
tzzdbi (8)
kwmam (184) -> wdybaj, cyielpd, hhifd, gexdb
ujktzrv (71)
dvoja (54)
opsqjx (71)
hlrjgro (63)
oqjkafl (32) -> iwyjkz, auneqme, awccm
vuyeiv (65)
qhmyi (130) -> dvfuz, scruak
tnayomw (62)
ezdhr (179)
ypfsjd (60) -> rdmrzdv, yhpiqn, cbyvzkp
auneqme (86)
kabixkr (73)
jntohm (41)
oyamw (65)
utivde (62)
qhcolun (76)
qjcbye (535) -> fetkt, pcewixh, vaubjz, ojhlp, mnvgaqh, rcjiwg
rfdpm (80) -> wcevmtt, tlayc
ovvrx (84)
zdhqhh (200) -> ylyef, onogiv, tohdgsa
ofidu (349) -> pjxqt, cytlm
zqmlv (59)
uzuawb (47)
unvje (71)
osbbt (1214) -> gqmscj, vyriv, bkkop
tmyhhql (51)
zxson (61)
rhlllw (11)
xtkxwd (17)
lijeejc (57) -> mgkkyx, thzyz
crhho (1255) -> rfcfgg, chnfm, tuekps
jqkaucw (53)
rerckg (63)
kgevjdx (84)
muyodod (19) -> wolet, zzjyw
zjzoo (65)
evqkdq (91)
mecyei (75)
bvfcsls (227) -> knpwsi, ypfsjd, eilzvpr
ntabiof (365) -> rfohgya, yoqbgjb
gqmscj (155) -> kjlmdi, scaec
bkmvwp (11)
tuekps (169)
mksan (18)
apmdcz (16)
plumb (89)
gcmio (126) -> ujjpxe, jjxvns
uqvyau (6)
bdkoa (125) -> ctynr, tvnco
xseshzl (76)
mnpvutr (39)
lghzki (27)
citugfl (14)
vxgwku (16)
jwaskb (251)
olkopyn (66) -> qjcbye, cstgeia, uojcup, ycctk, dkhuccj
jscjw (72)
xatyg (71)
xpfxwd (8)
jjjks (7331) -> qpefm, dlhiqh, gtervu, pcnroj, jijwa, bgbnlki
ewvvi (53)
ycbgx (1531) -> tzmndkd, fpynzz, tywzhc
cstgeia (613) -> ckwopo, sjiel, akwfu, ehvqyl, wtyfb, gcmio
ursjc (45)
mabkbom (57)
lafoho (250)
xmvbzka (49)
oasspz (67)
epumc (86) -> pcdjo, rerckg, dwknfpc
cyielpd (6)
lmqaxz (146) -> ghobhlm, qvvcx
ydqjoa (84)
zfkfhfn (33) -> txapm, pygfz, ekszr, nbivp, wltpv, jsjaxjk
sslltlv (45)
tifqyde (2264) -> koxiwk, psulm, rcjqzp
vonzkut (76) -> bdafpm, nvlxqp, gxsbt
blagy (338)
cluxnul (15)
kdevmnr (77)
cmxulaj (44) -> mnkqevh, mkbgt, nrcbc
tygwtg (25)
wpnqifq (11)
jilomb (68) -> uduan, mecyei
kytnz (52)
gynfwly (66) -> lynvd, dxhcxko, xaatl, leulsg, zworz, fkbrmim, jjjks
msfxkn (130) -> ietfhkq, kigdvl
wewzb (164) -> yzptgez, ctytfj
hdzxuz (24)
ghbfyqe (5)
hbkujf (133) -> ukghke, aplbnbx
iwsjla (38)
dnalt (35)
gmdsb (75)
pcnroj (2553) -> oljci, losdis, sdnkg, zchulv, crhho
dzrkq (94)
cjcapa (292)
dzohgiq (43)
rhlpt (88)
dkvzre (99) -> mieecfc, nvdouwn, dbnelj, onlwq, ayaouxu, xrhsy, bvrlaep
zpntow (72)
vohta (58)
jqmlxr (173) -> eiyxf, fydjnl
lxhkgs (85)
qoiuwmf (1008) -> vbnlfuo, wjyreb, sdbksb, lptkbzc, wopfs, khiom
btgrhku (24)
nnhsuja (16)
jgwvp (84)
vdpmvm (28)
iimfrx (59)
oyfma (21)
sqypc (7)
txapm (272) -> isggu, yookz
zhbiaq (45) -> rqzfef, kplegas, ayejq, xevhcxq
bkcghv (35)
yjakrqa (70)
lmwrf (51)
uwhjd (94)
bpphw (49)
vtvnn (114) -> uadnb, huunb
blmfs (52)
rtsajcq (6)
lazjlns (190) -> xbyjur, edjwkkk
rnddnj (14)
vobeup (41)
kifer (53)
jdzbqlz (15)
wlufwnr (89)
bqznix (75) -> stiln, duophdw, yhzuh
aovlz (45)
dyrvvfn (340) -> prxseo, vxgwku
ukghke (28)
lnczb (69) -> tzntd, cfnce
qllluo (57)
asbzj (89) -> yjxyb, hsifo, fhasacf, vwojb, gcbcyn
sruyra (47)
ohplzu (58)
fmkkz (319) -> owigu, ikbqz
zimrsr (223) -> dxympe, fhpaqmd, ayury
cdbkci (79)
tchbf (93)
wdybaj (6)
bexvn (39)
rcsfkrb (11)
bgbnlki (2277) -> oqbfkud, qsqis, xhyqtqz, qorkez, qwzac, oewlch, gsiypfp
zworz (41633) -> ootkqm, wfovakv, inldh
dwknfpc (63)
xjcprt (87)
ghobhlm (41)
erfndpf (89) -> kwbovke, adxhax, cipwzoy
setrric (31)
erggh (197) -> fksxl, ghbfyqe
pzksun (873) -> vyozfv, jxfqev, kbppnh
pddllsa (21)
xeomb (44) -> vuyeiv, bwrpskq, qwiekvk, gxzkde
dfmlo (80)
guvuihw (39)
khqhd (42)
fphgbca (59)
fhasacf (269) -> fovilf, rjnzany
oyamsv (38)
kjjyu (65)
pfqbht (22)
amzwex (218) -> vprkzya, wxixtqj, oktfn
lwljcg (85)
hpeaki (35)
rcjiwg (35) -> odckjtb, jlfgvr, tdbne
ktazkb (57)
tgvpi (92)
sdhqlxw (1239) -> eklahxm, ejzeqxz, kabcmf
aonfaf (52)
mlqxueo (67)
akidv (21) -> cnvyxm, fphgbca
bozlg (67)
ewpbnsk (64)
thzwwh (92)
tuieolg (7624) -> ldnrw, cfuqhip, rjzrh
tzrppo (51) -> tfhdccw, kbses, jlbcwrl, efsrz
fbjbwv (290)
dmthn (25)
witovp (6)
ugjzag (24)
agliie (844) -> qjaywg, rridl, myaiu, antcinm, izhdt
ebgdslc (31)
abmqz (31)
hsfjd (21)
ootkqm (9535) -> uplweb, bdkoa, ehrqn, fpqudx, assuj, rjguqr, jwaskb
ldxgz (18)
atfhgyu (57)
hbzju (71)
rrywqx (69)
dxqqhhd (188) -> pzewv, oelrlp
ixtrjm (92)
njeff (28)
dxboqki (78) -> nstdz, ferzy
mnwefbn (65)
bugfhhx (357) -> abbnh, intfltq
qorkez (1280) -> euqiok, ibvaekj
anbcmt (17)
iprkb (26)
vflyupn (34)
ruwzkz (362)
xrxseqy (94)
mszzq (8)
thzyz (77)
xyxtm (92)
qeubhb (65)
fmtyy (35)
hpowmso (1854) -> jhysc, xeomb, nzwxd
ywvmghy (63)
rridl (131) -> nqvflzq, vwotlbl
dydwqz (97)
mhnlgyl (1185) -> qntstgd, qzpgle, aozygag
uycjl (292) -> xcqvne, ruxnrix
ohcszyk (32)
gtervu (88) -> dkvzre, awufxne, osbbt, ycbgx, wdjzjlk
xcqvne (35)
moihbcu (75)
wpoga (57)
rjxvbkb (52)
bihsr (21) -> kyjkjy, hgscccv, yonjpz
vmmjgy (742) -> vdkxk, yhyxv, cpfbmv
gwnipjd (24)
brcpgz (57)
dczcz (1862) -> wszghwt, navhthp, lsfggz
wmaywuo (87)
vrfkofp (49)
nrcbij (64) -> pudyrbw, ghdime, xseshzl
yxtjx (74)
dnzdqz (179)
gxsbt (58)
oqrmida (222) -> lixqwd, dnalt
nddymc (93) -> ewvvi, netsm
iekurr (71)
tcghqr (43) -> mnhojtk, ruwzkz, veksns, wrochvi, uycjl, umtfn, qgvsuqv
ikfihh (140) -> tygwtg, vlmhyer
ziypsz (84)
ehjnzn (79)
exwxepj (175) -> jszpe, guvuihw, ykruy
capxp (68)
nhpcp (34)
qzpgle (77) -> gwnipjd, mrcoxnt
edjwkkk (14)
uteosk (65)
ofyxhb (11)
tulxgem (213) -> mabkbom, btcjg, ktazkb, evcdf
dkttrfw (219) -> ahqfoz, kytnz
mttvnpg (9)
tzntd (59)
euqiok (9)
wgypwo (290) -> btgrhku, aqkdlo
chnfm (169)
vlmhyer (25)
urjwj (78)
miijab (49)
faihy (65)
skkatyg (20)
zfnoo (18)
dcmxxd (35)
evnlr (1175) -> erfndpf, hicuenj, zybeuh
qzqyveb (80)
wpdbejb (90)
trbnpf (88)
yxxpgv (70)
wyomko (184) -> tgfgem, clnkvox
dxhcxko (45) -> ftzht, ypsme, rmtbjn, pjyld
lixsvp (11)
mofkvq (126) -> ejuiv, abmqz, xqobg
zqtkn (79) -> ugjzag, dtzyg
xhonf (95)
kiauwpn (22)
nmstp (44)
hsifo (210) -> wfdlusw, myonh, qunldzi
whnjtp (146) -> zswyy, bmugsve
txkgx (60)
icjur (76) -> lwaeqm, rhdudt
fpynzz (24) -> kepkgvf, kabixkr, jbexk
qunldzi (49)
ucxedq (84)
wndpfac (55)
hicuenj (122) -> ootidjt, hlrjgro, ywvmghy
kkdaw (65)
dmxhbe (141) -> myhrxc, jbbivz, behncf
borbd (21) -> ujktzrv, hbzju, xulwh, xatyg
adxhax (74)
zwsrt (2544) -> xmvbzka, egrfvic, fovkc
hhqlfj (81) -> xqgwwr, zmlmsuf
jiuuhl (78)
dcouu (6)
yetcvhx (71)
rfcfgg (28) -> sruyra, bqmqbi, uzuawb
pygfz (92) -> kledju, upevl
etscle (45)
pzjbbdd (93)
fjpgybt (21)
mjsck (90) -> mkkpihe, fmqjogk
cfuqhip (57) -> ixkicz, yqnihs, vifwkwa
jkamtqv (80)
ulvncs (172) -> bexvn, jzsmy
cytlm (46)
axikbt (9) -> tjffhqh, mogwb, cykvxs, ydqjoa
lageym (228)
jmlznjh (50)
wszghwt (160) -> brcpgz, wryfov
yxdld (38)
fukgfu (20)
mjlnaz (72) -> zcgfj, jiuuhl
bchfd (109) -> ccityq, nmvvgr
ogvod (1281) -> bihsr, erggh, dqgfb, xguhm
gcxrx (91)
ttnnyy (92) -> lhsccbq, dpdjpal
kxflpnl (16)
ehvqyl (192) -> zjgok, ecdieu
lsfggz (94) -> itttb, wpdbejb
aacegb (8)
wxdpcvv (8)
viufoq (25) -> bekxoxk, wmaywuo
pqnte (70)
rmtbjn (78) -> lytcy, aiunuvn, hfvhp, dczcz, kqaoir, ekihql, qkrydu
imjutr (187) -> wgeig, wqbhby
swpfak (21)
vmvxwar (38)
uxrrjqx (205)
mddejh (441)
fbhidp (17)
vunam (90)
rnbqhk (62) -> rdwkvr, oyamsv
bezvr (55)
kbses (93)
dqgfb (43) -> hpuyku, rycmr
uadnb (88)
dnrfjv (55)
wykkwa (67)
kyjkjy (62)
wrochvi (150) -> kifer, xaifomj, usodrg, jqkaucw
krmphks (12)
jbzaips (36)
qjaywg (94) -> khpat, jcpavmj, bchlcqs
kayqqz (77) -> kdqjj, sbdja, gmcsy
zjgok (23)
mrcoxnt (24)
wopfs (159) -> oasspz, zgssy
herqgqi (36)
zcdzv (11)
assuj (137) -> atfhgyu, pxhwecw
cvgbj (48)
lywkfsn (127)
cpfbmv (204)
scruak (30)
lsteeu (162) -> tatubv, rprjk, tgblta, uxrrjqx, pweea, sgieno
hhlxxzt (96) -> ixtrjm, tknmww, cnbekc
gmwtl (49)
sjiel (238)
pweea (51) -> eggmj, clpekkm
cnnrj (78)
eilzvpr (213) -> ksrsmjx, zpntow
cipwzoy (74)
apdfe (184) -> xrxseqy, leegl
bkkop (347)
cuhrgp (81) -> ohtply, vrfkofp
kepkgvf (73)
odkzxae (91)
qmqrpcl (92)
bgrxlxe (65)
xqobg (31)
awccm (86)
slhitu (27)
dihzy (79)
jfdscql (362) -> amrbhgv, rfdpm, ecxfenj, dxqqhhd
eqsxuaq (49)
hlcnxe (1998) -> fcpde, zyniqni, offjbb
pdmhva (18)
dtzyg (24)
xpker (36)
gqzva (6)
thqkxpl (38)
avelbqj (31)
mrigsjh (55)
ltbpi (17) -> vwcygm, herqgqi
odckjtb (72)
tdniuai (39)
tohrp (65)
wryfov (57)
vhrtptr (139) -> bpqhqbg, pacsxn
ohrraic (94)
eyale (97)
beraw (14)
mfywm (79) -> erkarx, vscjrx
knpwsi (261) -> cvgbj, uzjejte
wjyreb (41) -> rabxkov, rxqfv, gcomv
rdwkvr (38)
mmerpi (5)
cbgyk (43)
oxyof (44)
tjhmz (51)
zmqom (42) -> grazu, yxkldyi
rxanas (210) -> ctfjb, ifbxvs
lynvd (42593) -> tuieolg, pddteg, pixmx
scaec (96)
zbcra (38) -> rjeunph, edkigpl
ciogjt (375) -> tygnc, vhrtptr, ccckie
uvsny (44)
mpgixa (110) -> bmmmp, btxepbv
aqkdlo (24)
yjxyb (77) -> tceog, pqnte, yxxpgv, aokpx
tlayc (77)
kjlmdi (96)
rqhhyc (214) -> cizehi, sqypc
tgkusb (32)
xguhm (102) -> syyfs, hpeaki, fynniwm
koane (8) -> rawuoi, hkjwio, vpynuf, exxcrj, ljhtov, pwdpe, bdymucv
phmtqf (175) -> aodnlv, jancm
rjguqr (183) -> fgdqr, ccsrkny
mnhojtk (218) -> kvdrnf, nkjgwn
ejuiv (31)
rijipom (107) -> ijmfnt, ymduw, vdpmvm, njeff
bbfaxid (138)
yoqbgjb (25)
bymzb (68) -> zneumoh, zhopeqm
qntstgd (103) -> bbkfiz, zonni
ahqfoz (52)
gfqtz (98)
yvxuom (154) -> jbbmcxg, ppiiyuy
zxkvs (12) -> bclicjl, yfruc, axikbt, nglji
vwojb (97) -> qeubhb, kkdaw, ucfhxsv, rythrvz
akpnwc (90)
rawuoi (166) -> dzouqwl, vztyfp, dqgivj, cssov
eggmj (77)
isggu (5)
jszpe (39)
kmarvbs (90)
btxepbv (90)
hjjfdj (11)
sfrbkzf (45) -> tgmle, mddejh, tulxgem, ofidu, mjaol, dhqzt
ibvaekj (9)
leegl (94)
lfjtmkg (6095) -> lsteeu, zxkvs, sdhqlxw
itttb (90)
wlajxb (201) -> tgyavjg, eszxg
jlvwibm (46)
hcqnhm (58)
iqygrur (44)
bqifoq (55)
fovkc (49)
aozygag (125)
prcjhey (92)
fetkt (203) -> nnhsuja, kxflpnl, xumsm
qjzol (15)
rufvv (162) -> qzcbtes, xekimqs
dhlrq (88)
mwztduj (75)
ydumax (61)
boygd (13) -> wiayzvp, mdhhd, jkqvg, wouprxh
uyrght (80) -> hvcii, lxhkgs
nglji (311) -> cfaniht, anbcmt
pfutotv (44)
zvwkjew (60)
miftchq (21)
xaatl (56147) -> dgjls, qoiuwmf, koane, fnoteku, pavwo, hpowmso, yehxck
oejqkp (13)
oewlch (659) -> tgffx, eiwjrxx, ksybvgt
dwpqell (35)
mnksdxf (138)
obfet (87) -> iolmgs, piouhkc
kazqnjr (391) -> qngfv, aacegb
kmogwi (1139) -> hkjtym, tgujvov, dkttrfw
behncf (16)
ofosgz (80)
xejner (239) -> jmlznjh, foyfb, pxihdrd
ylyef (30)
lqumiws (88)
jancm (58)
rdkvtq (77)
alztgi (32)
myhrxc (16)
ycctk (1381) -> qtgibw, lkorngt, mufnw
wfovakv (48) -> lppvxfk, tznwmc, utgxfsh, zyympz, asbzj, ijcojo
lhsccbq (42)
tglnkgk (81) -> wrxiwgy, wrfxdc
ptkqcl (41)
cipgxee (72)
ecjdzpq (35)
ykpdiav (51)
wdjzjlk (1631) -> iplyhhc, pgchejz, kwmam
ymhxg (48) -> vohta, ohplzu, edpjix
vuetjdb (157) -> pbxjvnl, jdzbqlz, xhnmlcw, vipurf
skpcerk (88)
hfvhp (2018) -> wewzb, opmmmd, zmqom
afrywlt (80)
amrbhgv (126) -> xpker, gkkrlg, jbzaips
qsrpqe (236) -> brztp, kwwsxez
cdpfk (92)
oksoj (51)
eiwjrxx (45) -> hbsmlf, dlabh, rcjxcou
pzewv (23)
zonni (11)
nkssh (38)
nmvvgr (34)
iteizf (21)
dvfuz (30)
scmiv (54)
qqishhq (14)
egsil (38)
iipqh (299)
icqjyqd (16)
zktnjab (87)
nkskfx (94)
leulsg (44696) -> tifqyde, olkopyn, lfjtmkg
eskdbe (17)
dkhuccj (96) -> kbyot, zhbiaq, hhmavd, xejner, cqlwzk
mkyam (372)
wxixtqj (37)
ilxfa (34)
woczl (60) -> okseah, afeqhu, cnnrj, cmaxybh
rjnzany (44)
lppvxfk (1001) -> nrnmjo, phmtqf, bqznix
uzjxd (196) -> zfnoo, wlaslo, tijkvua
ekihql (1184) -> ezxsv, vonzkut, dkyswuu, uyrght, uzjxd, yjomyth
rljjjo (192) -> sobzsd, ykljt
mfjeyx (49) -> tdniuai, vdjmhb
qjdpk (28) -> wzxei, jopyit
liamld (25)
rjeunph (18)
vscjrx (79)
gwyfm (287) -> liamld, ucqdz
gnmydtk (67)
xspxxb (71)
mwirmq (188) -> mnpvutr, dmvjz
myonh (49)
mupbrv (218) -> phkwq, hrjgaqj, bwvemo
kztkiqt (13) -> egsil, mjugqpu
khpat (53)
wchlcja (190)
tjffhqh (84)
geqwvx (129) -> acxlck, zqnmsyb, ojnziip
ufyavk (1838) -> vunam, kmarvbs, kzdugfh
nvlxqp (58)
izhdt (57) -> gfqtz, sztqzfl
zfhxg (345) -> srneoo, ygmhxm, epumc
wkhggv (25)
jjvxxtt (194) -> ldfsxw, mhkcp
mhzgkxx (156) -> qzakfh, tnayomw
bchlcqs (53)
ymduw (28)
grazu (90)
fgdqr (34)
swrkuc (199) -> gylsj, cyzzsc, blmfs, aonfaf
zrpqdzn (62)
dkyswuu (120) -> zjzoo, bhxdhmr
mjugqpu (38)
gijtd (75)
huhoda (191) -> bpphw, eqsxuaq, gmwtl
vdkxk (90) -> hrvztx, fjhqmnv
huunb (88)
stiln (72)
exxcrj (316) -> dzjyzd, pkoww
bdymucv (304) -> uzsytm, sslltlv
dyscpqo (49)
tremw (94)
uotzz (61) -> xyxtm, cdpfk, qmqrpcl
vlpop (64) -> hshyx, tchbf
rdmggka (18)
owigu (9)
bpqhqbg (63)
aeyykn (66) -> lrkfnoy, ltdrusl
kwwsxez (27)
jgtpw (84) -> cxnjv, zelucu, ygurya, mrsrl
oljci (892) -> asifeu, aoehbj, oqjkafl
xqgwwr (23)
ctfjb (20)
cxnjv (66)
nrnmjo (105) -> pzjbbdd, nvvxl
wydbqai (65) -> ryulu, ydumax
ghdime (76)
cnvyxm (59)
xffvy (59)
qtgibw (52) -> ovvrx, ziypsz
ukkaxy (211)
yurfpj (65)
qvvcx (41)
ygurya (66)
zchulv (72) -> cmxulaj, tetdx, huhoda, blagy, wgypwo
rqzfef (86)
drgtn (84)
goxfpk (6)
bcjecw (80)
njabgq (57)
gcwcs (93) -> dokgk, epsjj
fhzhqie (65)
ccsrkny (34)
onlwq (59) -> ekxxsqa, jlfho, ekabfx
ldnrw (760) -> plumb, yvhilh, kztkiqt, ltbpi
nnxfo (63) -> xffvy, zqmlv, krcoaft, iimfrx
mrsrl (66)
vztyfp (57)
pacsxn (63)
maiimn (77)
usodrg (53)
rzphpv (48)
pavwo (1890) -> mofkvq, fmxtg, rijipom, mgqfa
gcomv (84)
pjxqt (46)
oywob (47) -> ynayo, ixxkvtz
fovilf (44)
ypsme (11966) -> zfkfhfn, fiufzkb, ubgrma, puurcu
etyyw (11)
ccityq (34)
hrjgaqj (29)
dhqzt (305) -> capxp, pttij
dpdjpal (42)
fbqpk (38)
qkrydu (1886) -> leyiz, mwirmq, hhwngc
mufrkl (60) -> wyomko, lafoho, rxanas, vlpop, ulvncs, padxpdx
cssov (57)
wgeig (15)
rythrvz (65)
juptypm (14)
dxympe (23)
ckwooig (84)
zgimdwb (107) -> znucug, mrigsjh
bbhyth (53) -> xcqwjrm, kgmwfq
vyriv (50) -> vwbjuvx, xergqq, wlpfcsr
qzcbtes (65)
kgmwfq (63)
losdis (1165) -> vhmijln, lteyo, viufoq
iplyhhc (122) -> cbgyk, dzohgiq
fpqudx (187) -> iuxgzr, icqjyqd, apmdcz, mhapqqy
jaxidl (86)
xhyqtqz (77) -> swrkuc, ncfuru, kazqnjr
rznnmp (70)
jvhxfl (39)
gxzkde (65)
uiwaf (97)
cfnce (59)
fkgxak (328) -> wpnqifq, xbucnh, qjwfsfk, rcsfkrb
nrcbc (98)
aodnlv (58)
yvhilh (89)
cyzzsc (52)
pjedtl (55)
bvrlaep (143) -> dnrfjv, wndpfac, bezvr
clpekkm (77)
wnyxznj (65)
wonxzkm (57)
vifwkwa (173) -> zvwkjew, txkgx, vvqpffs
qroirmg (45)
mdzkkjf (22)
pmfkdn (14)
qgfstpq (12)
mhydkla (7)
yekjlfd (106) -> ecjdzpq, dcmxxd
gxiwcqv (186) -> zfhwfsw, lovypfn
ifbxvs (20)
lwwyx (41) -> vuetjdb, ciejakr, imjutr, zgimdwb, sdnlegj, gzixhdc, tlvkwlx
tfdbdq (22)
dzggd (663) -> iipqh, nnxfo, veggtf
xwltxk (2001) -> yirnxk, baawtw, msfxkn
tznwmc (1199) -> dnfvldg, dbufjl, pnhibed
uiioczf (130)
rcjxcou (56)
tatubv (205)
vvqpffs (60)
bugwblt (14)
umtfn (80) -> hbylau, dzrkq, rugltaa
cqlwzk (217) -> besihsl, cwdmlj
pxteg (56) -> odkzxae, gcxrx, cotpovw
nbivp (106) -> ndnku, gjcxbx, iqygrur, oxyof
twvib (18)
fyzeso (57)
nlndxah (11)
wolet (91)
pixmx (4482) -> hrgbkby, bvfcsls, tzvawqb, jfdscql, gqggcxb
ctytfj (29)
lwaeqm (83)
cnbekc (92)
ekzvjj (35)
zgssy (67)
hbylau (94)
yonjpz (62)
btcjg (57)
tdbne (72)
edkigpl (18)
amnoi (141) -> noxvvva, dfeyr
gwtcgyo (24)
zrnlo (42)
ndnku (44)
etwwy (16)
hsoxt (62)
kdqjj (16)
cgouh (26)
mnvgaqh (128) -> jntohm, vobeup, ptkqcl
ocgkcxp (13)
ayaouxu (200) -> ksyewek, gpfrztg
baawtw (50) -> nnguamj, yetcvhx
ykljt (13)
znucug (55)
ypqxs (31)
pmgrf (21)
anmeg (58)
fmxtg (84) -> ursjc, eqhxqxm, qroirmg
dgjls (1926) -> lgxjr, hrphtyk, mhzgkxx
prxseo (16)
vhxlart (70)
zyniqni (147) -> khqhd, zrnlo
dtexqt (88)
mqybmn (86)
pljyyn (73)
jzibybz (248) -> zrpqdzn, hsoxt
dokgk (17)
bclicjl (217) -> dxufd, jhcsmc
tlvkwlx (133) -> fjpgybt, miftchq, oyfma, ytivjxk
qsqis (1208) -> aovlz, mxusu
mhkba (21)
kvdrnf (72)
fcpde (175) -> beraw, qqishhq, citugfl, pmfkdn
jcpavmj (53)
bbkfiz (11)
wqbhby (15)
mfvkd (21)
hhifd (6)
ibonrqn (50)
jzuvrtp (57)
zelucu (66)
zneumoh (80)
ljhtov (394)
okmqiy (46)
cmaxybh (78)
phkwq (29)
mvfhc (420) -> wchlcja, ikfihh, ybnvm, cztikzk, qhmyi, uebnns
jzsmy (39)
edpjix (58)
myaiu (129) -> pqmrmke, iizbmoz, rhkbrsr, apwsmv
mnkqevh (98)
noxvvva (94)
ncfuru (87) -> bcjecw, kybsigz, mtcxdod, ofosgz
pttij (68)
eutcl (77)
mzwbtsa (119) -> eceocsy, ecoqyk
wrxiwgy (25)
fydjnl (32)
bsixe (80)
oxcuf (80) -> crobb, xpfxwd
yqnihs (56) -> wqoucl, rsqyvy, zetvslt
mqhkd (100) -> pjhry, ljhxxd
bwvemo (29)
hbsmlf (56)
sduuog (9)
nnguamj (71)
xrhsy (266) -> qbrrjg, juptypm, bugwblt
rxqfv (84)
bqqwmry (73)
gyutarg (51) -> adbolgz, vdvadz
vyozfv (99)
njrfnvt (96)
ruxnrix (35)
sgieno (107) -> miijab, zryrfnw
vpynuf (214) -> akpnwc, rulhhsl
mkbgt (98)
iwyjkz (86)
itxycu (45) -> eutcl, kdevmnr
ytspbx (184) -> nkskfx, ohrraic
ayejq (86)
fhpaqmd (23)
jjqguew (24)
mnfqc (789) -> zqtkn, hhqlfj, lywkfsn
aplbnbx (28)
hshyx (93)
ozfzz (11)
rhbouej (181) -> udkyxw, rzphpv
wchxl (67)
inghbu (1167) -> vcvypf, ljqmiu, tglnkgk
fticuc (1360) -> vleydj, lnczb, igpabio, wydbqai
ivwcuc (88)
peuppj (29) -> oqrmida, txxnutu, fzqsahw
ppiiyuy (34)
mgkkyx (77)
akwfu (50) -> tremw, uwhjd
tceog (70)
acxlck (73)
woves (32)
veksns (308) -> slhitu, lghzki
ljhxxd (38)
dqgivj (57)
sdbksb (171) -> hbkjjtt, zxson
erfnrn (6)
nclwgga (201)
mdhhd (28)
rhzimzq (74) -> drgtn, raqjoxn
rqymw (97)
nkjgwn (72)
vwcygm (36)
acfyjc (85)
iajslqp (32) -> avelbqj, ebgdslc, vzqnfs
dlhiqh (7301) -> rdrad, vmmjgy, uvmqarq
koxiwk (929) -> fcscdg, geqwvx, jgtpw, zorvynv, zotwsb
yjomyth (97) -> oksoj, lmwrf, tjhmz
hhwngc (138) -> ewpbnsk, swurfm
ikdsvc (1609) -> icjur, ebmjz, rxivjo, rhzimzq
mewof (12)
bmugsve (36)
zcgfj (78)
dyimc (54)
iuxgzr (16)
fxasat (57)
ixkicz (353)
fvfek (96)
rycmr (82)
xfydt (87)
ybnvm (50) -> dwpqell, fmtyy, bkcghv, ekzvjj
abbnh (33)
pbxjvnl (15)
lrkfnoy (78)
lpbayb (8)
oeqvt (19) -> grcox, otplae
bdafpm (58)
upevl (95)
jijwa (3632) -> xwltxk, ikdsvc, tcghqr
tgyavjg (18)
bwrpskq (65)
lhsncv (67)
kqaoir (80) -> ytspbx, dyrvvfn, bkldmro, qonfiro, hhlxxzt, jzibybz, slrfd
dlabh (56)
drfwgmu (249) -> zktnjab, cmfkxo
qbrrjg (14)
tijkvua (18)
stkodp (84)
zunhob (79)
dnfvldg (85) -> yjakrqa, lahain
pjyld (10676) -> zfhxg, oenxsfm, ciogjt, ebmcu, mnfqc, zgqzrc, pzksun
dxufd (64)
zryrfnw (49)
ycacj (25)
smmgkir (31)
bymkeq (71)
bzsiwp (71)
wqoucl (99)
njdmj (73)
tzvawqb (1010) -> qirjqsm, muksdck, oxcuf
ljqmiu (31) -> ibonrqn, imezigo
pppxrv (62)
xbyjur (14)
hkjtym (247) -> iwsjla, thqkxpl
ucqdz (25)
ekxxsqa (83)
jbexk (73)
mkkpihe (24)
evcdf (57)
sobzsd (13)
rfohgya (25)
ahmitv (11)
wtyfb (238)
inldh (4965) -> ogvod, agliie, wenii
mieecfc (138) -> lwljcg, acfyjc
yrlks (44)
gcbcyn (289) -> vflyupn, mbwenqu
xevhcxq (86)
xayglgm (21)
hdlovco (5)
awagc (18)
zqnul (38)
rhkbrsr (31)
ejzeqxz (51)
onogiv (30)
ltdrusl (78)
yzptgez (29)
bkldmro (212) -> bsixe, afrywlt
zksmijj (189) -> zsvuw, hjjfdj
pqmrmke (31)
gzixhdc (124) -> ypqxs, setrric, smmgkir
ytivjxk (21)
wiayzvp (28)
hbkjjtt (61)
rhdudt (83)
eqhxqxm (45)
gpfrztg (54)
slrfd (340) -> etwwy, qldijf
iolmgs (95)
ckhip (88) -> vhxlart, rznnmp
pcdjo (63)
ojnziip (73)
nvdouwn (161) -> llibmc, jkfob, dyscpqo
gbnxgny (12)
vzqnfs (31)
lixqwd (35)
crobb (8)
vdbihjp (69)
llibmc (49)
lkorngt (132) -> tjpatnk, nmstp
ecxfenj (78) -> urjwj, jesiypo
ekabfx (83)
opqoq (648) -> mjlnaz, bymzb, lmqaxz, lageym
yfruc (303) -> pmgrf, hsfjd
fkbrmim (71889) -> peuppj, uobgj, llventw, duepwu
nzwxd (103) -> wykkwa, oydxsh, bozlg
kcuygjx (71)
hvcii (85)
fjhqmnv (57)
txxnutu (97) -> bcyarwn, uteosk, kjjyu
hrgbkby (53) -> saowgr, lprmau, ntabiof
xhnmlcw (15)
clnkvox (33)
bhwca (76)
uojcup (985) -> ttnnyy, yekjlfd, pdxylvu, fjzpjk, mqhkd, ibiwh
mxusu (45)
ixxkvtz (71)
gqggcxb (29) -> tzrppo, bugfhhx, drfwgmu
bsrkr (21)
erkarx (79)
kabcmf (51)
tknmww (92)
qgvsuqv (206) -> jhgsgy, enbnfw, uflldd, jvhxfl
cfaniht (17)
jrgwnfh (57)
rdmrzdv (99)
krcoaft (59)
ifelr (77)
hpaggeh (88)
wlaslo (18)
wouprxh (28)
hhmavd (359) -> pxzdv, qjzol
bqmqbi (47)
aokpx (70)
lteyo (157) -> mfvkd, swpfak
hrphtyk (19) -> xjcprt, bbzmsv, berje
pkoww (39)
jsjaxjk (88) -> aidknm, uiwaf
uebnns (127) -> iteizf, lzxrfk, tlfsn
wcevmtt (77)
yxkldyi (90)
qpefm (823) -> kmogwi, ufyavk, evnlr, rmlddp, fticuc
cmfkxo (87)
qldijf (16)
nstdz (65)
aoehbj (182) -> dyimc, scmiv
eceocsy (35)
hwrxn (53) -> dihzy, ibysnvc, zunhob
nglea (50) -> pfutotv, yrlks
fmqjogk (24)
tgblta (141) -> elmog, woves
jkfob (49)
vcvypf (113) -> mttvnpg, sduuog
wlpfcsr (99)
lybovx (18)
mbwenqu (34)
obkhwze (38)
vleydj (173) -> mhydkla, msthql
uduan (75)
ykruy (39)
aiunuvn (1629) -> gjrqs, ukkaxy, lijeejc, zlgpmfs, zksmijj
lqavsk (73)
afeqhu (78)
veggtf (104) -> wnyxznj, mnwefbn, fhzhqie
jbbivz (16)
uzjejte (48)
jhysc (274) -> cldbz, cluxnul
vhmijln (97) -> tmyhhql, ykpdiav
nqvflzq (61)
cotpovw (91)
yookz (5)
skjtvyz (18)
szutvca (32)
navhthp (136) -> rrywqx, vjxmbzm
okseah (78)
ebmcu (30) -> glsrg, ckhip, rqhhyc, jjvxxtt, rdnpoms
wuybunc (54)
oktfn (37)
tvnco (63)
tfpuqgs (24)
ijmfnt (28)
zswyy (36)
sqiac (24)
xbucnh (11)
zmpwz (12)
cotvk (91)
ohtply (49)
ysigfem (177)
jlfgvr (72)
kplegas (86)
mbezs (71)
rprjk (205)
jlfho (83)
melsryt (58)
oenxsfm (275) -> cuhrgp, qlgme, bbhyth, dnzdqz, ezdhr
sdnlegj (207) -> hdlovco, mmerpi
dzouqwl (57)
rugltaa (94)
imezigo (50)
iizbmoz (31)
vipurf (15)
awufxne (2141) -> qllluo, wpoga
gkkrlg (36)
dfeyr (94)
zmlmsuf (23)
lzxrfk (21)
jhcsmc (64)
ksyewek (54)
jzcmf (24)
xumsm (16)
eiyxf (32)
intfltq (33)
yhyxv (160) -> ahmitv, ozfzz, zcdzv, rhlllw
jopyit (51)
rmlddp (64) -> rufvv, gxiwcqv, cjcapa, exwxepj, qvbfob, zimrsr, nrcbij
dnazkh (80)
raqjoxn (84)
cgfykiv (38)
tygnc (13) -> ucxedq, jgwvp, kgevjdx
ojhlp (137) -> zqnul, vmvxwar, cgfykiv
hpuyku (82)`

const test = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

type disk struct {
	name     string
	weight   int
	children []*disk
}

func (d disk) recweight() int {
	w := d.weight
	for _, c := range d.children {
		w += c.recweight()
	}
	return w
}

func (d disk) balanced() bool {
	if len(d.children) == 0 {
		return true
	}
	w := d.children[0].recweight()
	for _, c := range d.children[1:] {
		if c.recweight() != w {
			return false
		}
	}
	return true
}

func parse(s string) *disk {
	names := make(map[string]*disk)
	for _, line := range utils.Lines(s) {
		d := new(disk)
		utils.Sscanf(line, "%s (%d)", &d.name, &d.weight)
		names[d.name] = d
	}

	for _, line := range utils.Lines(s) {
		if !strings.Contains(line, "->") {
			continue
		}

		var name string
		utils.Sscanf(line, "%s (", &name)
		d := names[name]

		under := line[strings.Index(line, "->")+len("->"):]
		for _, u := range strings.Fields(under) {
			u = strings.Trim(u, ",")
			d.children = append(d.children, names[u])
		}
	}

	unheld := make(map[string]struct{})
	for name := range names {
		unheld[name] = struct{}{}
	}
	for _, d := range names {
		for _, c := range d.children {
			delete(unheld, c.name)
		}
	}

	if len(unheld) != 1 {
		panic("expected only one remaining disk")
	}
	for name := range unheld {
		return names[name]
	}
	panic("unreachable")
}

func wrongweight(d *disk) int {
	for _, c := range d.children {
		if !c.balanced() {
			return wrongweight(c)
		}
	}
	// all children are balanced, so one of the children's weights must be
	// wrong
	counts := make(map[int]int)
	for _, c := range d.children {
		counts[c.recweight()]++
	}
	var goodweight, badweight int
	for w, n := range counts {
		if n == 1 {
			badweight = w
		} else {
			goodweight = w
		}
	}
	for _, c := range d.children {
		if c.recweight() == badweight {
			return c.weight + (goodweight - badweight)
		}
	}
	panic("unreachable")
}

func main() {
	// part 1
	d := parse(input)
	utils.Println(d.name)

	// part 2
	fmt.Println(wrongweight(d))
}
