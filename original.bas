10  'JADWAL SALAT SELURUH DUNIA ANTARA LINTANG 65 S - 65 U
20  'OLEH T.DJAMALUDDIN (LAPAN BANDUNG)
30  'MENGIKUTI KRITERIA DEPAG RI
40  '*********************************************
50  INPUT "NAMA FILE, e.g. D:SALAT.JAD"; FILE$
60  OPEN FILE$ FOR OUTPUT AS #1
70  RAD = 3.14159 / 180
80  INPUT "NAMA KOTA                             ="; CITY$
90  INPUT "BUJUR ;  BB:-, BT:+ (DERAJAT)         ="; LAMD
100 INPUT "LINTANG; LS:-, LU:+ (DERAJAT)         ="; PHI
110 PRINT "BEDA WAKTU= WAKTU STANDAR - GMT, e.g. WIB:7 "
120 INPUT "BEDA WAKTU; BB:-, BT:+ (JAM)          ="; TD
130 PRINT #1, "       JADWAL SALAT "; CITY$: PRINT #1, " ": PRINT #1, " "
140 LAMD = LAMD / 360 * 24
150 PHI = PHI * RAD
160 DATA "JANUARI",31,"FEBRUARI",28,"MARET",31,"APRIL",30,"MEI",31
170 DATA "JUNI",30,"JULI",31,"AGUSTUS",31,"SEPTEMBER",30
180 DATA "OKTOBER",31,"NOVEMBER",30,"DESEMBER",31
190 N0 = 0
200 FOR MN = 1 TO 12
210   READ MONTH$, D
220   PRINT #1, "            ", MONTH$
230   PRINT #1, " "
240   PRINT #1, "TGL SHUBUH  TERBIT  DHUHUR  ASHAR  MAGHRIB  ISYA"
250   FOR K = 1 TO D
260     N = N0 + K
270     A = 6
280     Z = 110 * RAD: GOSUB 540
290     IF ABS(X) > 1 THEN GOTO 310
300     T(1) = ST
310     Z = (90 + 5 / 6) * RAD: GOSUB 540
320     T(2) = ST
330     A = 18
340     Z = (90 + 5 / 6) * RAD: GOSUB 540
350     SUNSET = ST: T(5) = ST + 2 / 60
360     Z = 108 * RAD: GOSUB 540
370     IF ABS(X) > 1 THEN GOTO 390
380     T(6) = ST
390     A = 12: GOSUB 540
400     MIDDAY = ST: T(3) = MIDDAY + 2 / 60
410     ZD = ABS(DEK - PHI): A = 15: GOSUB 540: T(4) = ST
420     PRINT USING "###"; N: PRINT #1, USING "##"; K; : PRINT #1, "  ";
430     FOR I = 1 TO 6
440       TH = INT(T(I)): TM = INT((T(I) - TH) * 60)
450       PRINT #1, USING "##"; TH; : PRINT #1, ":";
460       PRINT #1, USING "##"; TM; : PRINT #1, "   ";
470     NEXT I: PRINT #1, " "
480   NEXT K: N0 = N
490   PRINT #1, " ": PRINT #1, " "
500 NEXT MN
510 CLOSE
520 SYSTEM
530 END
540   T = N + (A - LAMD) / 24
550   M = (.9856 * T - 3.289) * RAD
560   L = M + 1.916 * RAD * SIN(M) + .02 * RAD * SIN(2 * M) + 282.634 * RAD
570   LH = L / 3.14159 * 12: QL = INT(LH / 6) + 1
580   IF INT(QL / 2) * 2 - QL <> 0 THEN QL = QL - 1
590   RA = ATN(.91746 * TAN(L)) / 3.14159 * 12
600   RA = RA + QL * 6
610   SIND = .39782 * SIN(L)
620   COSD = SQR(1 - SIND * SIND)
630   DEK = ATN(SIND / COSD)
640   IF A = 15 THEN Z = ATN(TAN(ZD) + 1)
650   X = (COS(Z) - SIND * SIN(PHI)) / (COSD * COS(PHI))
660   IF ABS(X) > 1 THEN GOTO 720
670   ATNX = ATN(SQR(1 - X * X) / X) / RAD
680   IF ATNX < 0 THEN ATNX = ATNX + 180
690   H = (360 - ATNX) * 24 / 360
700   IF A = 18 THEN H = 24 - H
710   IF A = 12 THEN H = 0
720   IF A = 15 THEN H = 24 - H
730   TLOC = H + RA - .06571 * T - 6.622
740   TLOC = TLOC + 24
750   TLOC = TLOC - INT(TLOC / 24) * 24
760   ST = TLOC - LAMD + TD
770   RETURN