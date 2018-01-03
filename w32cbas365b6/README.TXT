*** *** Chipmunk Basic README *** *** 	(version 3.6.5b6, 2011-Oct) 
Chipmunk Basic is a simple old-fashioned Basic Interpreter.  Versions
of Chipmunk Basic exist for the MacOS, linux and MSWindows.

Version 3.6.5 is a both bug fix and new feature release.  It is still
free for educational and non-commercial distribution and use (but only
for those who don't believe that using the BASIC language causes brain
damage. :-)

chipmunkbasic.exe has the beta test graphics commands
basic.exe can be used from the cmd console command-line, no graphics

Chipmunk Basic only comes with a simple man page.  But there is more
documentation and some tutorials available on the web site:
  http://www.nicholson.com/rhn/basic
The Chipmunk Basic language is mostly compatible with books on
programming in BASIC that were written between circa 1978 and 1990.
There is also a Chipmunk Basic support and discussion group on
Yahoo.com:  http://groups.yahoo.com/group/chipmunkbasicforum/
Please email bug reports to rhn AT nicholson.com 

Shortened version history:
Version 3.6.5(b6): 
        fixed mat multiply
	initialized mat origin
Version 3.6.5(b3): 
	fixed bug in format$() with leading zeros
	fixed bug in exit while statement
	fixed bugs in empty for and empty edit statements
	fixed object string bug
Version 3.6.4(b7): 
        fixed some mat command syntax omissions
	fixed a function subroutine return bug
	allow subroutines to return 1d and 2d mat arrays
Version 3.6.4b6:
	added the local keyword for declaring local sub variables
Version 3.6.4b1:
	added det() function to calculate 2d matrix determinants
	added pset(x,y) to Windows graphics version
Version 3.6.3b8:
        bug fix in win32 rnd()
Version 3.6.3b2: 
	bug fix for clicking on .bas files
	icon added
Version 3.6.3b1: 
	added some beta test windows graphics commands:
		graphics 0	:' turn on graphics window
		moveto x,y
		lineto x,y	:' draw a line
		graphics color r,g,b	:' set color (range 0..100)
		graphics text s$		:' draw some text
		graphics cls		:' clear graphics window to white

	added fn version$()
	improved fn format$() precision capability

Version 3.6.1: changed exponentiation precedence, default mat origin
               added mat print command
Version 3.6.0: added exit status, fn kill(), fn bigendian(),
Version 3.5.9b3: fixed long if, instr(), mid$(), val("..") bugs
Version 3.5.9b2: fixed intl. string compares, format$()
Version 3.5.8: fixed format$(). changed int() to round to -infinity.
Version 3.5.7: fixed some integer array bugs, added w32 sys().
Version 3.4.7: changed not operator precedence 
Version 3.4.6: added matrix ops

copyright notice:

 * Chipmunk Basic is (c)Copyright 1996-2007 by Ronald H. Nicholson, Jr.
 * ALL RIGHTS RESERVED
 *
 * Permission to use, copy and distribute this software without fee
 * for non-commercial or educational purposes is hereby granted,
 * provided that the above copyright notice appear in all copies,
 * and that both the copyright notice and this permission notice with
 * the following disclaimer appear in the supporting documentation.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY OF ANY KIND; not even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR ANY PARTICULAR PURPOSE.

 * Portions of this release of Chipmunk Basic are
 *   Copyright (c) 1983 Regents of the University of California.
 *   Copyright (c) 1993 Regents of the University of California.
 * Portions of Chipmunk Basic are Copyright (C) 1989 Dave Gillespie.
 
-- cut here --
</PRE>
<H2>Quick Reference</H2>
<PRE>
/*
 * Quick Reference and FAQ
 *
 * Chipmunk Basic is (c) Copyright 1994,2004 Ronald H. Nicholson, Jr.
 * ALL RIGHTS RESERVED
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY OF ANY KIND; not even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

Updates to this document may be found on the WWWeb at:
  http://www.nicholson.com/rhn/basic/

---
Quick Reference for Chipmunk Basic
ported to the MacOS by Ronald H. Nicholson, Jr.	  rhn@nicholson.com 

Supported operators, functions and statements (reserved words):

	+ - * / ^ mod  and or xor not  > < >= <= <> =
	sqr() log() exp() sin() cos() tan() atn()  pi
	abs() sgn() int() rnd() peek() val() asc() len()
	mid$() right$() left$() str$() chr$()
	goto  if then else endif  gosub return
	for to step next  while wend  select case
	rem  let  dim erase  data read restore   field$()
	input print open for output append as close# load save
	random lof loc get put   
	inkey$  input$ eof()  files  fgetbyte# fseek# fputbyte
	run stop end exit quit cont  renum  new clear
	date$ time$ timer  sound morse say  doevents
	home cls gotoxy htab vtab pos() 
	graphics sprite pset moveto lineto window scrn mouse()
	varptr peek() poke fre push() pop  isarray
	sub call usr  def fn
	type class string integer single double
	asin acos sinh cosh tanh log10 floor true false ubound

	
- The "basic.man" file is included with the distribution.  It describes
  the above functions and statements in more detail.
- Line numbers are required for console input.
- The interpreter includes an dumb old fashioned line number based
  editor.  (enter a line, delete a line, edit one line at a time).
- The variable types are long floats, short integers and strings with
  a maximum length of 254 characters.  Variable names can be up to
  31 characters long and are case insensitive.
- Math is done using the built-in double precision IEEE math.
  All trigonometric functions express angles in radians.

---
Some Simple Examples:

  for i=1 to 10 : print i : next i

  c$ = a$ + b$		    : rem string concatenation
  a$ = inkey$		    : rem inkey$ is non-blocking 
  
  load "sieve.bas"	    : rem load TEXT file as a Basic program
  			    : rem file names should have .bas extension
  
  open "filename" for output as #1 : print #1, "hello world" : close #1 
  
---
MSWIndows specific commands:

  	SYS( "command.com /c cls" ) ' Win98/Win2K Clear Screen


---
Frequently Asked Questions:

Q: Is more documentation on learning how to program in the BASIC language?
A: The language is mostly compatible with books on programming in the
   BASIC language that were published between 1978 and 1988.  Since these
   are mostly out of print, your best bet is to try your local public
   library.
   
   Also, there is a short tutorial on the Chipmunk Basic web page:
        http://www.nicholson.com/rhn/basic/

Q: Why won't this Microsoft QuickBasic program run? ...
A: Chipmunk Basic was never intended to be exactly compatible with any
   other commercial BAISC language product.  It does include most of the
   features of ANSI BASIC language definition.  Many BASIC programs that
   are saved as ASCII text files and that don't use special features of
   the computer system can be easily ported to Chipmunk Basic.
   
Q: Is there a compiler for Chipmunk Basic?
A: No.

Q: What is the history of Chipmunk Basic?
A: In March 1990, p2c, a pascal to c translator, was posted to the usenet
   newsgroup "comp.source.misc" by David Gillespie.  One of the test
   input files was "basic.p", Chipmunk Basic in Pascal.  Other versions
   have been written in C for the Mac and PalmPilot.
   
---
Other notes or Bugs:

- Programs without line numbers can be loaded from a file; however
    sequential line numbers will be added.  The target of a GOTO can be
    a label followed by a colon. e.g.
    	foo: x=x-1 : if x>0 then goto foo:

- All input works like the LINE INPUT statement of other Basic versions.
- Line numbers above 999999999 will not list.
- Other reserved words (don't use these!) :
	msgbox  do loop until break  function
	method private public local   menu dialog memstat()
	draw play  bload bsave min max mat
	each redim resume  static
	eqv imp key is each option set width swap

A  sample program:

--- cut here ---
1 rem "sieve.bas" , a prime number sieve benchmark
2 t = timer
3 dim f(8194)
4 for i = 0 to 8191 : f(i) = 1 : next i
5 s = 8191
6 for i = 0 to s
7   if f(i) = 0 then goto 11
8   p = i+i+3
9   for k = i+p to s step p : f(k) = 0 : next k
10   c = c+1
11 next i
12 print c;" primes found in ";
13 t = timer-t
14 print t;" seconds"
15 end
--- cut here ---
*/
</PRE>
</body>
</html>

