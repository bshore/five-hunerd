# Five-Hunerd

Playing card game I learned when I was younger, recreated while tinkering with the [Pixel](https://github.com/faiface/pixel#readme) and [Beep](https://github.com/faiface/beep#beep---) packages, experimenting with game dev.

## Download

Visit the [release page](https://github.com/bshore/five-hunerd/releases) and download the asset for your operating system.

## Example Round
https://user-images.githubusercontent.com/14914028/157559458-b690857a-5ebb-4f41-8881-5e510a40af9b.mp4

## Rules

Real Five-Hundred is played at a table with two teams of two, you can read about the table rules [here](https://www.pagat.com/euchre/500.html#w7st).

In Five-Hunerd though, (this game), you play against a computer team, with a computer partner. Your goal is to bid on how many tricks you think your hand can take in order to get points for your team. The game ends with a team winning once reaching 500 points, or losing by reaching -500 points.

### Card Rank / Bidding

There are special rules regarding the Jacks, often referred to as Left or Right "Bauers" if the bid names a suit. The "Left Bauer" is the Jack of the same color of the named suit. The "Right Bauer" is the Jack of the named suit.

So for example, if I bid **8 Clubs** - the Jack of Spades is the **Left Bauer** and the Jack of Clubs is the **Right Bauer**. Similarly a bid of **8 Diamonds** would make the Jack of Hearts the **LB** and the Jack of Diamonds the **RB**

Rank if a suit is bid: 4, 5, 6, 7, 8, 9, 10, Q, K, A, **LB**, **RB**, Joker

Rank if no trump bid: 4, 5, 6, 7, 8, 9, 10, J Q, K, A, Joker

### Phases

The game consists of 3 phases:

- Bidding Phase - At the beginning of a fresh round, everyone has 10 cards, and you can sort your hand by the different suits and determine if you'd like to bid or pass (passing just shuffles & deals out new hands for now)

- Blind Exchange Phase - If you place a bid, the 5 cards in the middle are revealed to you. You can swap cards from the blind with cards in your hand in case you'd like to get rid of some pesky low cards.

- Trick Phase - Starting with you, the players play cards from their hand one trick at a time attempting to take as many as possible.
  - Highest card takes the trick, players must follow the suit of the first card played if possible
  - If a player cannot follow suit, they can "sluff" a low card of any suit, or choose to trump the trick, meaning you play a card of the suit that was bid. A trump suit played on a trick wins unless one of the next players also cannot follow suit and can trump higher.
  - Your goal is to take _at least_ as many tricks as you bid. So for example if you bid **8 Diamonds**, your team must take 8 of the 10 tricks, or you are penalized.

#### Credits

- Really nice playing card asset pack downloaded for free from [Kenny Vleugels](https://www.kenney.nl/), thanks!
- Awesome wiki's going over the basics of working with [Pixel](https://github.com/faiface/pixel/wiki) and [Beep](https://github.com/faiface/beep/wiki)
