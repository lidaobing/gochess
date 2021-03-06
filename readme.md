![Alt text](https://travis-ci.org/jonpchin/gochess.svg?branch=master "Travis CI Go Play Chess Image")
[![Build status](https://ci.appveyor.com/api/projects/status/96kvdw3mr190y854?svg=true)](https://ci.appveyor.com/project/jonpchin/gochess)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonpchin/gochess)](https://goreportcard.com/report/github.com/jonpchin/gochess)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)
-<br><br>
![Alt text](/img/gif/goplaychessdemo_shrink.gif?raw=true "Go Play Chess Demo")
<br>
![Alt text](/img/screenshots/lobbyResize.png?raw=true "Chess Lobby")
<br>
![Alt text](/img/screenshots/gameResize.png?raw=true "Chess Room")
<br><br>
Go Play Chess - Free Online Real time chess web server using websockets - Creator: Jonathan Chin<br><br>
12/10/2015 - Start date of project.<br><br>
5/31/2016 This project is officially open source. - All commit history has been cleared as some of the git history contained sensitive data. 
<br><br>
4/30/2017 I've decided to add a MUD. Chess and MUD's don't usually go together but by having them on the same web server and domain allows me to save server costs. More details will be revealed soon.
<br><br>If you are interesting in contributing please open a new issue here:
https://github.com/jonpchin/gochess/issues
In the issue state what you want to work on and I'll add you as a contributor. 
If you are open to doing anything I can assign you one of the open tasks in the back log.
Once your changes are ready, make a pull request and I"ll review it.
<br>
<br><b>Features:</b><br>
<p>1. User accounts</p>
<p>2. Time controls</p>
<p>3. Player ratings and matchmaking system </p>
<p>4. Save and review games</p>
<p>5. Server side validation of chess moves</p>
<p>6. Security(HTTPS, WSS, CAPTCHA, input validation, password encryption)(partially shown on GitHub)</p>
<p>7. Cron jobs</p>
<p>8. Logging(partially shown on GitHub)</p>
<p>9. Graceful shutdown and recovery of active games</p>
<p>10. Automated deployment of web server to production environment(not shown on GitHub)</p>
<p>11. Highscore board containing top bullet, blitz standard, and correspondence ratings.</p>
<p>12. Configurable user settings, chess piece theme, sound toggle, chat toggle, etc.</p>
<p>13. Database of over 1.8 million chess games, most of which are master/grandmaster games.</p>
<p>14. Player profiles to view rating and past games.</p>
<p>15. Player's name has the country flag where their IP originates from when they register</p>
<p>16. Players can view on going chess games in real time as spectators</p>
<p>17. Lobby and game room chat.</p>
<p>18. Rating history timeline on player's profile using Google Charts.</p>
<p>19. Players can play against cinnamon chess engine and stockfish chess engine (web workers).</p>
<p>20. Web server runs in Windows and Linux.</p>
<p>21. Forum</p>
<p>22. Daily news feed, contains mostly tech articles but it can be configured to get other types of news.</p>
<br>
<p><b>To Do List:</b></p>
<p>1. Chess TV where live games are randomly broadcasted.</p>
<p>2. Admin control</p>
<p>3. Friends list</p>
<p>4. Inbox system</p>
<p>5. Update help, documentation and screenshots</p>
<p>6. Post game analysis with Stockfish engine</p>
<p>7. Statistics page</p>
<p>8. Store zipped source files and databas backup in Google Drive or Dropbox.</p>
<p>9. Add release notes.</p>
<br>

<b>Special thanks to those who provided amazing third party libaries</b>:
<p>1. chess.js</p>
<p>2. chessboard.js</p>
<p>3. goglicko by Josh Hoak aka Kashomon</p>
<p>4. Tock.js</p>
<p>5. Cinnamon JS Engine</p>
<p>6. download.js v4.1, by dandavis</p>
<p>7. dchest/captcha</p>
<p>8. robfig/cron</p>
<p>9. gopkg.in/gomail.v2 </p>
<p>10. golang.org/x/crypto/scrypt </p>
<p>And many more...</p>

<br>
<b>Other Dependencies</b>:
<p>1. jsonlint</p>
<br>

<b>MUD</b>

Classes:

1. Warrior
2. Barbarian 
3. Monk
4. Mage
5. Thief
6. Ranger
7. Swordmaster
8. Illusionist
9. Priest
10. Necromancer
11. Witch
12. Paladin
13. Alchemist
14. Jester

Breakdown of classes and description:

1. <b>Warriors</b>:
Warriors are well rounded fighter that can use a variety of powerful armor and weaponry. 
Warriors can learn some basic magic.

2. <b>Barbarian</b>: 
Barbarians are ancient warrior who uses brute strength and raw fury
to excel in combat. After many generations of fierce training they have become resistant to magic.

3. <b>Monk</b>:
Monks are masters at martial arts and they can use their spirtual powers to heal and purge disease.

4. <b>Mage</b>:
Mages are spell casters who wield powerful spells. They can do great damage but are physically weak
as a trade off.

5. <b>Thief</b>:
Makes a living by robbing others.

6. <b>Ranger</b>:
Cunning forestmen who lay traps.

7.  <b>Swordmaster</b>:
Swordmasters train with the blade day and night until they become experts.

8.  <b>Illusionist</b>:
Manipulates space and time

9.  <b>Priest</b>:
Holy magic casters

10. <b>Necromancer</b>:
Can raise people from the dead

11. <b>Witch</b>:
Uses dark voodo magic and can create healing potions

12. <b>Paladin</b>:
Holy warriors who can use holy magic
13. <b>Alchemist</b>:
Able to manipulate bodily fluid's and create rare elements
14. <b>Jester</b>:
The utlimate prankster

Races
1. Human
2. Dwarf
3. Elf
4. Troll
5. Ogre
6. Hobbit
7. Vampire