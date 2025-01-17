package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DB_tbl_counter          string = "tbl_counter"
	DB_tbl_admin            string = "tbl_admin"
	DB_tbl_admingroup       string = "tbl_admingroup"
	DB_tbl_mst_version      string = "tbl_mst_version"
	DB_tbl_trx_sdsb4d_day   string = "tbl_trx_sdsb4d_day"
	DB_tbl_trx_sdsb4d_night string = "tbl_trx_sdsb4d_night"
	STATUS_RUNNING          string = "background:#FFEB3B;font-weight:bold;color:black;"
	STATUS_COMPLETE         string = "background:#8BC34A;font-weight:bold;color:black;"
	STATUS_CANCEL           string = "background:#E91E63;font-weight:bold;color:white;"
)

const Sourcechar string = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 !@#$%^&*()-_=+[]{};:,.<>?/`

var Keymap = [...]string{
	`JUNj@'LPg[!Zd<Do1k:+\-.5(?W%9F_/8T*heA$Q)Ha>=z0SlwMsKx,R }4IEqXntp2BCmr3O7cybi{]6YVfGv;^u`,
	`VHAX0 ZCT9zMQ[Sera+BEK1@_n86,DJ?m$wb/>h-;YP{Ly2NguGqFx]k=O.j:sUlvo7\cpt!')34Rd^(5WIfi%<}*`,
	`WZq<[ilDOx!tIAB'J}>Hk\=TF h?wsP+dN,VoKL()jz$e7b@MgQE/fa8r394]SX25{c-u^ny.C_pmYv0G:R*1;U6%`,
	`prmxVJX HFjsyf-L'@di{bn!/W0YB3:,vMuQZ[a1k(%I4A\]Cg8E_;qe<>2.GR6*^$}TlwNhc9)=zKUP?5tD7o+SO`,
	`9wU/,c=;Sl[H]yJurD!C*_d-F.ZV5Kzkf0?P1gt+npLs4^xObeYj\'I$><TaiQ@:)vMR6GX2qEo}{m7(W8N%A3Bh `,
	`i60WuABw$,p=n<V(Xrf%h-v^ )Y1@{;dT\?q2DM'tklgxQby_ZU7msS+aC:4R}3c9/ojKL]EJ*H!Nz[ePO>8F.GI5`,
	`R9V-vnT4mXw^ihK/k EecDL7ogU50;l+:?IOt=<GZP*Cr{QF]x8!3%paqAYb6$)z'WM\y1,s2_Bj}fN[JH(d>u@.S`,
	`KZG912y8hs7;pn}z5m.qHD*o/%t]63cd<IX+S W!\0wi-)P>JQMb$^LNa[gAY:Ckvf'=?{,ujxR@V4(_lFeOBErTU`,
	`2gxA?ptLv6}0Qehcn/BN(';XflU^C+1SY!Rb4wMq><r8-j{d_K@JkW9Z[s$]%.Ou7*\TH =iy:E,mVaPoFI3)Gz5D`,
	`7fy91_Mlz3[UL%(eq\B-jacYkITm^RQwdnS)W?=tP8sX2]opHJ{>E@A!Fu*rhZ$5 .';Gvx,<D0N}C6/ObK+4ig:V`,
	`RzP6'U:<}NIwXoyr$CEpt!L{e7Zi)?%cTWSfA\B4K OahJ5V[umM2nYx*klF+q-g,d.D9vjQ@3/0_>]GH1=(;bs^8`,
	`j ^bf'Iq4mC,./A@_Z{(L?;*X}lhF)SE]NdBD87ap3UyT2GYH>-VQzgu0v!MOwx$W=n6ir[<9eRckoJ:K%s\Pt15+`,
	`{sftH3b6$)0F5T> 1A-4<K}*pW[Rl@e!jqC_^y',Z(ScVov/QXm7uJh+MwiO?BEDL\9=g2z:GN;d8rU.Y]%PkxIna`,
	`boCxX17$vA0O+}y;)z=F%@lR*Nh8a5fn:(9s>\V!MWQkq_PTi3rEe]ItgL/',D{2^Zup?.B-UmdJ[ K6<HYwG4cSj`,
	`B'H{?Fi+ RE-h<tqN\bL>UlIJ(*gw@c5Z!KeMC/dX6mz71}QT03)Y=p$ku_A[ofjasO:P9V.nrxS%W,4yG^2v];D8`,
	`c5@(0G4T>-1A_3<K)*rW}Sm$g9kqC]%x',b^RdXpu\UYn6vIi MtjP/BEDL+8=h2z:FJ;f7sV.?Z[Oe!HyoNlwQ{a`,
	`JwuE@C) e=G.aW8L0mg1?R3hx\orMt5{zPcfbl/'K^;>UZjS$:]vOT7NY2qHn+-k6}d*X[D9FpQ!V(s%_4,Biy<IA`,
	`o4;!)X('=SNzCreYs]Juap<[hkEmy@tIWZUg_>F7/+PVdM6,*vKQ2HTxnDj -l3^c$b{B5GiR8f}w9\1.L0%?:AOq`,
	`M+j]Ff'cEN5mhV2x;l .\KRs=/HZS%Co)QG[t673@napAYX09$wbO-vz>r1_Dk^iP(LJ!e*gu:q{4<?T8IB,Ud}Wy`,
	`(-d F.ZU7Iylfz/P1gv\orJt4*xNbeam+'K$;>TYiQ@:[wMS6LW0qGn=_k5{c%X}D8EjR9V)s3^!<Cp2]?HBu,OhA`,
	`S)]mx>G+qhjfrOap(XW3isV{b<\yY@wc5_toMN*.I!D9Hv^ukU$6B[:CP01;'K4R F}/TeL-872=gZlAzQ%?ndE,J`,
	`wBbP=uz.r2 Fn%oS^NL!d8ip:q@0<+X*QD'Rt[U;K3sGJ)MYjV],/Cv5(e1y-?A}cxkO{_Z9\IH6ElhTam>g$f7W4`,
	`M1sFJ^LVjR[;.<Br3*ezy=>}b0nO{-W(+NG!AqpDUl:g?k@Y$']T cCXhav%IZfu\9mtEQ47xKHP,S)/dw_8i6o52`,
	`la61Cp>D;QWw,:NfV}Js-UK/x@$8(qerGdc4%)0FhS?z3u7+Io{nT[PO^g!jtAv]5'\Y_MEBXy=ZHb*2Ri<km9L. `,
	`9vU\>c=;SM[GyzIYsDmBt_f-F,b:lKdoh2+Qgj1 rwLV7*3OHJET]'Rx<?aCW/u.65Z}nXAk8ei)(4@^q$p%N0P!{`,
	`hH.I8?\TafQ6;}tNS1MXu@GjK l[)e:cBJ3OkY2x,n-$4_Pr=5(WL%*m{/A]pyzv7FZiUVw0!o'RD+qs>dg^9E<Cb`,
	`l9\wvYZ_'S(L^Ry{xrh%7I-BKEa34DGX$g>U=;eqb?[}*+umzN6i8)/1M5V,0F<t:O@!kT]Qj2csdnC pfAHW.JPo`,
	`H\rjkgsQbq{YX3lpW}c>+yZ$xd6_toNO(,J@D!Iw*vmV%7B]:E'P02;L5T G)/ShM-984[ianC1U^=?fAzK.<eFuR`,
	`8Izle0?P2fv\orJt5{xNbdZk+'K^;,UaiR$:_wOT7MX1qGm= j4[c*Y}F9EhS!V(p6)@<Hs3%/LC-]Wn.A>QDguBy`,
	`tfYu-Jxaq,_jmEo0$vIXZUh].F9</QWeN8;{wLP4KTzpDk= l3)c^b(G7HiS6V*r!}@?My5%>RC+\gs'A:nd1O2B[`,
	`0u;h :\KPo=/GYQ$Dk)OE]q451!jZlAXWy6@tbM-sw>pz[Ci%mR*LJ9c8en.r^x<+T(IB,Sg}U'H2fFN{Vd3a?_7v`,
	`2xj\%6+Rp?5_YP*{i[>F n13y9MXhCVUu4!rBaL/qv:mw-)f@NO^,I0EzJg'e8o.sA$<G;Qd(=]KkbDSt7TZH}lcW`,
	`cp)YX1joW{d<+vZ@ue4]qnOP*>J!D8Is%tka95C[;E'Qyz:L3T H(\ShN_$76-lbmFxV2}=rB0M,/iA?g.GwfUK^R`,
	`@TSudiU$X=]qV5oZy(lkKL%\F6A3Gp9nhY70'*<>Jst?,HzP}D!_OeM{481)gajBwW2[-rCxN.+m: f/EvcRQ^Ib;`,
	`Cq2*czx\,]Z1mM} U)+LG$ArpDSk:f.j%W^'_R=bBTeYt(JXdu?@inFO07vIEH;N{<Ps-[a>K!o9ywV/hQ8l4g365`,
	`<K)(uX}To$h!mvB]^2'>d*UeYsy+SZp6zIk-NxlQ\CDAJ 8=i54;EH,g9tV?.b_Pj%L{wMrqF7G@[c/O:aW1R0nf3`,
	`p^h$nvB[*3'<c{UeXtz+VYq8yKl_N0kR\CEAJ !-j65;DI,g9wT?.b=:Qm)H}uLsrG7F%]a/OZW2Pxif4>S1d@(oM`,
	`c7iq;^9x< X@OZSnu[QTl3tFg)JvhN_:'>C*2(b0y\,A-Y1mL}=U%/Kd!B$pEkjz.46R8?{MHoDfWVr]Isae+5PGw`,
	`fu(Z8W^GcHC>YT7iz91-MO03[U/n)hr+A_ldeaoJ'K%R.VjpQ$X\yvP!tb4]qsF,w?DkgBL{2xS}@:5<;=EN6 *Im`,
	`B%=b]}L>:)Z-mjF[UE+NfAC!@cq5V0T4DYH; WQzgt1v^KMwx$S_l8hp{.9iadXnI<J6R?eosP*k\2yO7ur(,3G/'`,
	`sKu3)yPegcm\L%;,WdjS@'-xQU6NZzrIn/+l4]f*b}H8JiV5X(p7{9?Mt2!=RF[_hv.B<kYaTqDOw GE$01C>o:^A`,
	`nx@sJYaWh_,F8?\RZeN6.{tLQ1KUvoGj -kz)d$b(E4HgS3V%m5*7/Mr2!+TI[]lu;A<pcfXwDPy=Oi^09C:q'}B>`,
	`Q*El[OF=q562$kZmAXWz7@vbM+sx,p0 Dj^nR)LJ!c8fr;o%y>\U(NB:Pt_S.I3iGK}TahY-'</Cw4{e1u?]9gHdV`,
	`Hl@ +Etj:A4*MZJ912y8gs7'pn[x5m.qKF(o?%r]30bc>IW\S-V6/zvf_)P<LRNa$!GQY}eBX;DiudC,={^kUhOwT`,
	`a!td2[poRS(<L9E7Kr%qkb@3C{.F'Twy;O4W-J)+ViQ]$85_menGzZ1^}vDxP>\sA=u,I*6gU Nf0YlchB?/j:XHM`,
	`oXx^jiMN!=G5A3Fm8ngW4z'$/<Lqs\>JwR)E@]ScO(671*hZlBuV09}tCyQ?-r;[p+H%2dP_Kb{Yvaf: kU.,IeDT`,
	`O[=W{/MH%BtsDTo'h.m*Z^:-U\eCVgav}JYdu<)lpEP08xKAFN(?Lq+_b;I]n zyS$iR>fwQX!9r37k5j6G1c2,4@`,
	`e(AZ>W'GzgaPwr/f}?_FLm[-CUM8:i@J^n01x5hTj,RQu26q<VH)ot=lv*;c7bI9DE4X!dp+k$y ]N%K.\BY3O{Ss`,
	`Whav1+Xcs!0Ko]PzpT\BDAJ-$ l97;EI>k%xV?.e='Un}L[2MutG@F(_f{H:bY5Oymi4<R3gw/jqN68*rQZdS^,)C`,
	`Upv[SWm4uGi(KwjP_:',C^3*d1z\.A-b5oN]=X{/Me@D%rEnk2>7!T$; QLtBlaYq)IsZh?9cfFx08gJRVO6<+H}y`,
	`6dw7z]LNy1{T\k^ep=A[iZbXjH:I@R,UclQ$V xrS8qW0)nmG<o/DgfCK%usP!5't?;+FM2->EaO(B3}YJ_4*9v.h`,
	`inY(c<=vZ!td2[poPQ*>K9E7Jr%skX83C{;D'Rxy.M4U_I)+TgO]@65}lbmFzW1^-uBwN,\qA/j?H$0eS LahV:Gf`,
	`UbgQ9;]uNS3MXwpIk+=l0}e%c{H5JhV2Y^m4*7/Lr6!\WK[-jv,B?odfasFTt_PR8nxO Z:>zG(1i$.qyCD<A)E'@`,
	`GzHbRwr<f]81[Ll-2)UK@^e*+{ivxt5FTd;POpy6m.VE_ko/hs}$a7IJ9\B3'uCc>Z4j=n:0 ,(DWq%?AXSMNg!QY`,
	`ZmAXVz6@vbM sx>q1-Dl%nP^LJ9c7gp.o$0?=U*NC;Qt[R:I3rFH)KYiT],<\Bu4(e2y+/'!akfO{}S58GEwWjhd_`,
	`mD7FOcK%\/ ;kv6Xtr}=>^UuhI@)Q![HB4.jl:Me+Y_d3R5?$N(W,LZSo8CTap{2fi'GszqA-<*E09]byxJ1wnPgV`,
	`F8Mr$qlb92D(,E'Svx.:P3X_L*=WiR[!@6}oepHzc1%{wG0T<+uA y?J)^jV]Qg5a4fmB\7n>UtNYOhk/KZd-sCI;`,
	`lG WE\OgBA()cr!X4U9DaF'+VQ2fw5y}KMx0*T-m$hq=;{jbdYnH,I8R<SipP%e_1uN7vk6]stJ.z?C3o:L[@^Z>/`,
	`n:f,k*V(' R+bBTeXt}IWZr?{joDMz9vHAC;K)/Gm=_SF]l\yuO-gP<YsLN$@p18h6d!JqQE>wi704%3[ac52.U^x`,
	`kF VE\NgAB(*cr!X3U9DaG:+WQ2fv5y{JLx0)S-m$hp_;}jZdYnH,I8R<TiqP@b=1uO7wl6]stK.z?C4o'M[%^e>/`,
	`QyoW=CEAL]!_l75.DJ<j9wX/,e+:Vn{K[zMusG@F*-f}I;bZ4Oxmg0?T2hv pqPt6^rRcdak\'N%>)SYiH8($1BU3`,
	`Af1gaQys>4], Gkm_=EtM3^KXH9z0w8eUh;pRr27n.qI}*l/%o['5xYJ!DWvSuVb?(6d\-N$FOLci+CPZ)jB{:T<"`,
	`mIH 0Kw{E@C(=eNG.cX8M1!3-?Tz4[/n^Oq\%tRfgbkL'P9;>YdjV7:]uSW2Qav*FoJ+p_)iAlDZ6Uxr$5,s<yh}B`,
	`B}^3'<f(VgZty=Waq8zKm]OxnS+CEAL-!_k65;DI>j9vU/,d :Tl)J{wMsrG7F%[c\N.bX2P0oh1?R4eu*iYH@p$Q`,
	`RZyrMl?/m2_g*d[L7NiW4Y^n5(9<Pt6!+XO}-jv;E>pefbqIVs{TS1huQ a:,zK$xU8=okFG]'C%A@D0\wcJ).B3H`,
	`EjKS@rmb72Cq>D;RYy,:OgX]Lu WM/z%$9{sftGed6^)1FiT?04w8\Ip[oU_PQ(h*lvAx-5'<Z}NHBa3+cJk=!Vn.`,
	`w>qz_Cl!nQ$MK7a5go.m9x? W%OB;Ps}S:J1rFI*LXhT[,'\Dt3^d0y=/)ZviN(]V@-HE6AkjGYp<e+f4b8{2RuUc`,
	`mzMjh@{_6LnY'49F3%:?w-ab=BU}O(TxGy]7D$P+EQJe2,INd*1Wc/ksg>[ ^<t5\RvuA!)qSrZX;pH.oCV80ifKl`,
	`AkTEdvQP>9S5=L{J-<nUNCmg(V8)@,:a6!+G;u]Wy.[0ZlojsRFY}DBirxf_K?4ep%ct7\X1bz/ qIwMh^O32*$H'`,
	`dK^?<='7v8Xsq[+,*UuhH$)P!]GZ4;5k:fc/r y2N3\%JFn.eSOi9BjQb}6WYaoxgCMRLV(_Az{@ETlI0>1pDmwt-`,
	`UfYt_IWar<[jnDOz@vHAB'K}>Gl/+TF=k?yuP\eN,XpLJ)(m1%g8c$MiQE.od9w5!4]SZ67{b x*q0;C-shV23R:^`,
	`Pfhdo+L$;,WekT@'_xQV5OZztJn\=p2[i^c{I6KjU4Y*q7(8?Mu39 SG}]lv>C/mabXsERw)HN1ryF<g:-0B%!A.D`,
	`SdWr[_UYo?xHkBMy!sG'.E(5)g=1\;C-e+qQK]aJ/RjFI^zLv@A6:8%cOD fZ3Pw40{}Vtu$,*l2Nm<7iThbnpX>9`,
	`mfCZ{)br@U4S!DxEXMvP6dw70_Iky2[QH%(ep\:-iYaVjFRn*NLzgoO]W?+uT$tc8=sqBG/,A}^J3 15h><'.;9lK`,
	`!S>VchQ9W-voT4nXx*jkJ/l EfdDK7qiU50:p+;\HNu=?FZO(Bs}PC]y8@3^ratAbY6%{1'eL<z2.m$_Gg)RI[Mw,`,
	`w9Wtr]\,(UvhH^}O$-GB7;kl:Le?Z+d6Q8<)M[X.NYRm!CSVn_5fiFs1oA>/ D3*'b@2J{=za0gcE4PIyTpKj%xqu`,
	`*I1JcSyV9hz!3 Nn05]WM^{gr?B=ldeamHUo)RQxipT}Y>\uP@tZ6-qkGK[.D%A$Ew_vsb(8/;',Fj2<:CfO7L+4X`,
	`PZoU BCJy9[k54>AG/i8vV+<d-.TO{IxzLbsEnDt]g}K,e:mNfqj3=Shp2'_u0MX!^6RHQFa\;Y7?)lWc(1*@$r%w`,
	`b5ho.p8x/-W@NB,Os{R:K1rFI%LViS};>\Cq2^dzw <)aykM(]U$_JG7AnmDXl'f+j9Y!?*Q[cEZgev6HTt0=4uP3`,
	`poST*?M9F7Ls$rlc83D).E'Uwy;P2Y_K(=XiR[@%5}neqGzb1^]vC0Q>+uA x<J{!hW-Og4a6fjB/tk:,NmIVZd\H`,
	`eL*673%mZnBwWz8$uA0O\]v;}y+F^@kR)Mg9a4dl:{5r,/U!KVQip NSh1qEc-HsfJ<'>C[x(Xot?.D_PbYI=2GjT`,
	`K,SlGF]_gv*b8Y%HdIC<aV7jz92-OQ03}W/p(is\A nehcqL'M$U;XmrR^Z+1wT!xf5[tuE.y>DokBN{64P=):@?J`,
	`Ow!rHAB'J}<Ge_[P>:{c okI]XF+RhCD$%ds7Z2W6EbK;=aU1jx4y*NSz0^Y-n9iq\.@lfgVpL,M8T?muvt)3/5(Q`,
	`9U2S8BXC:=TN0bs3v)GJux*Q_i@dm].(gYZVhE>F7O/RekM!W-wpP5qaz{nlH?r+Aj'fI%ytL^6,1\< DK4[;co$}`,
	`Ph)C}vEro5;!(X{'+TO0DseZt-Kwap<]ilGmz@uJWbVg ,H8?\QYfM7>*xLR3IU1qFk=_n4^d$c[B9NyS%j/2.:6A`,
	`km-=DtM!3^KYG8z1x7fUh;pQs06o,rH]*n/$q['42bJ9FXySwWd<%5i\_O@EPLcg}IRavjAZ CueN?.:+>(l)TV{B`,
	`hQq8cdG_f)AZ,WF0iaPyu?g[<-ELo} CUM!'l%K:*p12x6jTk>ROw47s/VH{rv+mz(;b9YI@DB5X$et\n^3.]JSN=`,
	`Ybs90Ln]QzoV+CDAK-@ l76;EI>j$wX?.e\:Um}J[yNvtG!F*_f{M,cZ4P1kh2<S3gu=pqOx8(rRadW5/'B^)%HTi`,
	`Nn03}VM$(hs\B_lceamJWk^SRyfoT*Y/ vU8uZ2[qpIK)<E9A7Ft%wjX@5].',Hxz>;G1O-D!=PdL{64r+gQiC?b:`,
	`/.)XxkJ*[R%-Ib8:9n'ig<v+06P7?(MGr;hWQl@CoTf_$Zdcq1mBOSLa^ Az]}HUYF2\5sENwKeutDy={p4>j!V,3`,
	`\Jn$M?GwmC7EQdL%+/-;kv6Ytr{ >^VugI@(R9}HB4.ji:Oe=Z]c2T3<!N*X,PaUo5AWbp[1hl'FxzqD_)8Ky0sSf`,
	`'e%I{jyzv5dSf>POr04n<TF}lq=it);b7cH9EC3U2Yh\k8s ]N!J.+Gm^L/Dwo:A@KRZQ(-?[,gu1Wpa*_$xMXVB6`,
	`5!uBzO={w.*y E@9lR%Ng6a3fm:(7s<+W8MXSjp]PTi0qFd}IreL\',C^x)Zvt/>A_VohJ[-Q$?HY2D1cGkb;n4UK`,
	`/KRq+<HaQ^Dm{PF-s562@kbnAYXz7$vcN=uy.r1_Cl%jS(MJ!e8gp;o*0>\V)LB:Tt]U'I3iGO}WdwZ ,?[Ex49fh`,
	`]K>'}a njF_UE+NfAB%$bq7W2T6CZG:=VQ0eu3w*JMxy^S-l!hp[.(iYdXmH,I8P?ckrR9g\1vO@zo5/tsD;)<L4{`,
	`U@ wRW5PaxqLl?/m1]g*e[K7MiX3Z^n4(8<Os6!+VJ}-ju.D>odfbpHTr{QS0htN_Y;,yG%zI9=kcCE\'B$A)F2:v`,
	`Y{+TO2Ctgau_Jxbq>]jnEp0$vIXcVi-.G9?/QZfN8,)wLR4KUzrFl= m5(e%d*D7HkS!W}s@[6;My3^<PB\:ho'A1`,
	`vDXMsP4bu6y{Hix0(RG!%dn :}gYZVhEQj$NLwckO*U+-qS8rW2]moFI)?A@'7Bt[pla^9;=>.<Jz1/,C3T_5fKe\`,
	`'FyeZPvq/d}?_ELl[-BUM9:h$J*kz0w5fTg,RQs26o<VG)nr=ju(.b7cI8DC4W3Ym\i@x {N!H;+Kp%O>Ata]S1X^`,
	`}Q%_HC9.nm:Mh?a+f8S7<)N[X;OZTo$DRWp-@ej'Gt1qB>/=A5{c^!K ,4b6liF3UI(PkEJ0zguyYxd2LrV\]sv*w`,
	`mdfbnKVl%SRygkT*X+_uU8tY2{poJL^/E9A7Fs$vjZ!4)>',Iwx?.G1O]C( PeM[653}iaqB0W@-=rDzN;<h:\QcH`,
	`y+,[a0nM{-U(\LG@BrqCSl:f>j$X%;]R cAVeYt*IWbs/^koEOz7vJDF'N)?Kp=}Z.H!m_xwT9hQ<g1Pd65238i4u`,
	`bJ$/? ;jv6Vsq}=>%TugG@(O9[FB4:kl'Ld\X-c3P5<^M)U.KYRm7CQWn]2ehErzoA+_{D0!,Z1yI*8wStfaHxNip`,
	`mb82D^,F:Svx.'O3W]K*=VjR}!@6)peqGza19(uE0Q< wA-y?J{%lU[Ph7d5gnB+$r>X4NYTio\LMftkHc/IsZC;_`,
	`$@pR^Nj7c5hq;(8x<=X!OYSnt]PUl2uEf}IvgL ':,B)1{azy/.A-Z0kJ_+T*?Hb9C6mDiow>34M%\sKGdFeVWr[Q`,
	`901x7fUh;pPsz6n.rH]^o/$t[53cJ8FYySwXe?%2k=_N!GQKdj{IMavlAZ)CmbL ,><'*u(Riq+\:@EWTB}4O-gDV`,
	`MfAB)(aq!V3T9DYE:+UP1du4x}IKwz*R k$go-.{iZbWlF>G8O?ShmN@X=ysQ7te2]prH,v/Cn'jJ^50c_[;6<\%L`,
	`r1u*FItw%P]f9Zk[.^cVWTeC<D6N\QYgL7R_snO3mUx$ijG+o Ad:bH8qlS52,p=>/EKz-?BaM)'y(X;{04!v@hJ}`,
	`vZq>_imEoz$tJXaVg].G8?\QWeN7;{uLP3KTxpFj =k1)c%b(D5HfS4U^n6*9/Ms!@+RI[-lw:A,hdrY0CO2'yB<}`,
	`^GcIC<ZU7hy81]MPz3}W\n*gr+A_ldeamJ'K$S.XioR@V=0tT9uY2{pqH>s/EjfDN(xvQ%!:w?; FO5-,Bbk)L6[4`,
	`Kk] CUL9'f%I)ly0w5eTg,RPsz4p<VF}nr+ju{;c7dH8ED3X2ai/m!t=[O@J.\Go*N?Bxh:A$MWbS(-_^>Zv1YqQ6`,
	`DAL]9_k64.CI?j8vW\,d=:Um{J[xMtrG!E^-e}K;aX3Oylgz/S0fu npNs2%wQbcZo+'H$><TYhR@)*qFP5Vi71B(`,
	`T!X [oU4nYw*jiLM@+F5A2Gm8lgW6z'%?>Krs/,IxQ)D9]PdN(370{hakBuVy$}qCvO<=t:_f\E^1bR-JcpS.Ze;H`,
	`@:$n'ig<x\29Q!>}MIt.kXRl*DoTh-^ader3mCPSNZ(=A1_[JUYG0+8qFKyHVupBj]{fz)b7W%E5Lvc6s/4 ,?;Ow`,
	`+>*VvjH%{Q@_IB7;mn:Mg/Z=e6R8<(N[X.OaSo!CTWl]9fiGt0qD?, A3)'c54K-\1Y2khF$UE^LpJPyxruzdwb}s`,
	`.RPsz4o>UG{lr=jv)'f6gJ8FD2W1ai\h5t-}Q9I;/Km%M?EwnB!CScN*+ ],ex7Xqp(_<@TkZH$^O0[ALy:bd3VYu`,
	`9y< Z@PbTou[RVl3tFi(JvjO_:',C^4*d0z\.A-a2nM}=U)/Lf!D$qEmk1;58S%>{NKrBhYXp]HsWg?7ceGwx6QI+`,
	`E p562$jYlAWUy7@uZM=rw,oz-Dk^mP(LJ!b8fn;q%x<+T)NB.Os[R:I1iFH{QXgV_>'/Ct3*d0v\?]9SecG}4Kha`,
	`WB: SN1as2v*GIux%Q]h9cm[.^fXYUgD<E7M\RbjL8T_wpO4oVz(lkF/q=Ai;eH!tnP@5,y+>-CK3}?JdZ)'0$r{6`,
	`kM8Ie4X2dm?!5t=[T7KVOjq)LQhzrAc^EsgJ],./'$y%Ywv <;{UxiF(_P*+GZ63l:ba\p-u1H0>9DCf}SRNn@BWo`,
	`aXh[,G7?\RYeN5.)sLQ0KVtoFj-_ky*d!b^H2IgT1W@m3%6+Mq48 SJ{]lu>A/ncfZrDUw=PO9vxi;p':$C}zE<B(`,
	`ux%PG6!bm];(fWYUgEQh8MLsaiN9S-}pR3oVy*klFH$+A5:2Bq@neT70.)/,?Itv\<DzK[4_JcO{1^w dXj'rC=>Z`,
	`jz1w7eSf.POs06o,TG[mr\ku}'d9gI!ED4V3Zi/h@t=]Q$H:?Jn(L<CxlB%FRaM)+ _;cv5Wqp{->8NbXA^2Ky*UY`,
	`/?l0]f%c}M5NhV2Y$m3*7\Qr48=XO{_ju.E<odgbpJUq)SRyinT[Z;>wP!vW1 skGH(:D@B9Fx-taL^6A,C'Ize+K`,
	`NO9 H4A1Gm7lhY3x'@\?Mpr/<JwT(E!}SdP*560^jakCuXy8)sBvQ+_q;[t=F%$fU{Lc2Zzbg:]oi,>KnIVRe-.DW`,
	`NH(ButDTo:h.m)Y{'=U\eCWgaw]KZcv?[lqFP2!yJAEM}<Lp+-X;I_n/0zR iS>dxOQ%$r49j7k@VsGb,3185^6*f`,
	`!.mn:Lg/Y+e9Q@?{M[W,NZRo%CPUl_$di'Gt2qB<> A7}b*^J=;6a8khF5TE)KjDH31fuzXxV0IpO\]rcysw4(-Sv`,
	`a/g!Q@>}M_X;NZSn^CRVm %fj'Ht3pB,.\A8]d)*J+:9b(okF$UE[OhDG54es1YyW2IiK<-cTzrw0x{LP6v=l?u7q`,
	`l3tGh*LuiQ]:'>D^2(czy+,A_a1oO[=V{/Ne@C$qFnk4.6!U%;-RMvEmZWp)JrXj?9dfHgx7sITYSw <B8\}PbK05`,
	`LOz2{V+m%fr A[kbdZlI:J!S,WejR@X-ytT6sY1(onH?q\DhgCM9vpU85;u/'<GN0=_EaP)Bx}QF>4$^7*ic3Kw].`,
	`tw$O[g8Zl}.^dUWSeC<E4L+PYfK5Q{snN1mTv@ijF=o_Ac;bG6qhR30>p ,\DJx-/BXM*'u(I:%y729!kVr?Ha)]z`,
	`kcdZlI:J@S>VejR$W-ysT6rX0(noK/p DhfCL9umU83;t+.\GOz=?FaP)Bw}NE]1!%7*qbxAgY^{[4'QM<25,iv_H`,
	`,{rMQ0KVslGg -hx(b!Z*H3IdS1W@j2$5+No47_UL)[it<A/mceapFTq]RP6kuX=Y;>yO^wf9?vnDC.'B}E\J%:8z`,
	`q.tI}@r\9u{43fL6Gb0UyZh?81n [R5JSNgl(KPdwmAa^DocH_,>/'$v!Vsp-+;*QkYE%]M7=CWz:xe<TX)i2OjBF`,
	`o,TH)mr+kv(:f4gK7FE2V0ai/j5t-}R8J;\In%M?DxlB9GQbO^=<_.hw6Xsq{]>@UucL!$P3*CAz'dZ SY[W1eyNp`,
	`O4X_K*\WkR[@$8{rfsG0c3!)yE2T,=xA 1.J}(oV]Ql9e7ipB+%v;Z^PbUmt?NYjzqHg/IuhM>CDFS<6:n5w'La-d`,
	`1w'7+:/Lmp=?IvQ%E8{PcM$34z@iYjBsVx5^qAuO }r.)t-F!9fR*Kd2X0bh;[6l>\TyJUNek_HSaonCZ<G(gD,]W`,
	`,[OJx:naTo(ErVl=*ehiu6pDSUQc)\A4 _LWZI3/$sGM0FRvmBj]}g1%d9Y^C5HkX7b{w8-!<Ny@+>PK?.tz';2fq`,
	`<}N_Y;OZSn^CRVl %ei'Ht4pB>,\A!]d)*K/.@b(ojF$UG[MgDE67fs2XzW3IcJ:-aTymw1x9PQ0v{h=r8qu+?5Lk`,
	`lVjQ]$%9}sftG2b5@{zF4S,=0A 3>J)(pW_Rk^d8iqC+*x.Z!PaUou?OTm1vIg\KreM<BDEL/7;h6y'HN:c[nX-wY`,
	`D*3(e0y/;A-b2oP]+W}?Ng%C^rFnl4:7$V@' SMvEpcXq[JtYk>{fiGmx8sKUZTh=,H6<_ORdL5.!uIQ1ajz9B)\w`,
	`49tYM-qw>oy_Ck@lP%LJ7a5fm.n!x? V$NB;Qr{R,I0pEH(KWgT[:<\As2^czv=/'*ZueO)}U6]GF3DihSbj+d8X1`,
	`XUeD?E5N=RZfM6S{tmP1lVv@ihG k[Ad.bH4ogQ3y>p_<-FLu]+BYO%'r*K:(w780!nWq,Ta29^x;JC}sz/j$\)Ic`,
	`k=(cghs5nDSTQa*\A2-]LUYI1?%pGMxFPrlCf[}ew$Z8X^E3HiR0W!o4)7/Kt69+VJ_<mv.' jdbqyBOz,uN>{@;:`,
	`gv)a9Y*H1IdRyU!i0@3=Mn25_TL([hr?B+lceZoJVp{QPzjqS]X>/uO%wW7 smFG-;C^A$Ex\4kb<},:'.Dft86KN`,
	`ZWgERf8OMsahN9T])oQ3nVx^kjGH@ B5;1Ap7meU4z.%\,/Iru+?DyL{!}KcJ(260*iYl'vSw[_q:tF=-b>$X<dPC`,
	`jL7GE2X0cl/k5t }T8K:\Jo^N<FwnC9DReO*?+-;iv6Zsr{_,$VufI%(S![HA3.hm'Qd=Y)bzU1]4M@W>PagqyBpx`,
	`n@oQ%MJ7b5hq.p9y? X$OB;Pt}S:K1sEH(LWiT],'/Cu2*ezx=<{awkN)_V!-IF6AmlDZj>f\g8Y4+^U[cGdrv03R`,
	`R(F9]QfL^562%manBwXz7$uAyO/_v;}x\E*@kS)Ng8b3el'[4r,?V!KWPip=MTh0qDd HscJ>:.C{t-Yoj<+G1IZU`,
	`X)'=S\dBUgZt[IWbr<}knDNz!vHAC:K{>Jm+ TG-l?yuP_fO,aqLM^%o1@i7e9QpRE.jh6x4$5]VY38(s/w*c0;F2`}

func GetApiPath() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	appsApi := os.Getenv("PATH_API")
	if appsApi == "" {
		appsApi = "http://174.138.22.4:7071/"
	}

	return appsApi
}
