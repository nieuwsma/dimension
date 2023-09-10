# a few thoughts on approaches
 
 * i could start by submitting an empty dimension, and seeing what i fail
 * i could start by trying to figure out what rules are in conflict
 * i could start by always working with QUANTITY rules first: exact, SUM, gt

## some good heuristics,
 
 * if a color is repeated a lot across many colors, it might be easier to just omit it
 * if there is a TOUCH-G-K; then a good pattern is to do GKGKGK around the equator.

# Thoughts on tasks

there are 7 types of Tasks

## what are the tasks that interact with eachother?

1. TOUCH, NOTOUCH
   * TOUCH & NOTOUCH interact if they are the same, or if any color is the same
   * TOUCH & TOUCH interacts if any color is the same
   * NOTOUCH & NOTOUCH interacts if any color is the same

1. BOTTOM,TOP
   *  BOTTOM & TOP interact if the same color is listed
   *  BOTTOM & BOTTOM interact if there arent any more places
   * TOP & TOP interact if there arent any more places

1. QUANTITY, SUM, GreaterThan (GT)
   *  QUANTITY & SUM interact if any color in SUM is QUANTITY color
   * QUANTITY & QUANTITY only interact if they are the same color
   * SUM & SUM interact if any color is in both pairs
   * QUANTITY & GT interact if any color is in QUANTITY
   * GT & SUM interact if any color is in either
   * GT & GT interact if any color is in both pairs

AH, I realize that these are a priori interactions, I know that TOUCHes can impact each other,
the other type of interactions are based on color itself; as that is the common thread, I really care as much,
if not more about all the places where G has a rule for it


## what are the tasks that conflict with eachother?

there are probably a few other cases, but lets start with these.

1. TOUCH, NOTOUCH 
   2. TOUCH & NOTOUCH conflict if the color pair is the same this is really only  a conflict IF its played. Which could be required by Q, S, or GT
   3. TOUCH & TOUCH could conflict when colors are played if there are too many rules (unlikely)
   4. NOTOUCH & NOTOUCH could conflict when colors are played if there are too many rules (unlikely)

1. BOTTOM,TOP 
   2. BOTTOM & TOP conflict if the same color is listed  this is really only  a conflict IF its played. Which could be required by Q, S, or GT

1. QUANTITY, SUM, GT 
   2. QUANTITY & SUM conflict  if QUANTITY-A-(1/2/3) & QUANTITY-B-(1/2/3) && SUM-4-A-B  as long as QUANTITY-A + QUANTITY-B != 4); 
   3. QUANTITY & GT conflict  IF QUANTITY-A <= QUANTITY-B && GT-A-B 
   4. GT & GT conflict if the chain is > 5  : A > B > C > D > E


# Rule types
 
there are a few types of rules:

* `SHALL DO`, like QUANTITY, SUM, and greater than, you MUST do these, you cannot avoid them by inaction 
* `PREDICATE`, meaning if you place the color criteria, it must conform to the rules, e.g. TOP, BOTTOM, TOUCH, no TOUCH.
    * TOP and BOTTOM apply even if there is just one sphere of a color, -> they apply to the whole color
    * TOUCH and NOTOUCH only apply if there are more than 2 speheres placed: ACOLOR & BCOLOR


