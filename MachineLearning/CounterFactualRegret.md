# Counter Factual Regret

## Defining key ideas

Regret:
-at a time t, is a different between our algorithm total loss and the best single expert loss.
-expresses how much we regret for not following the single best advice.

$$ R = L_{H}^{T} - min_{i}(L_{i}^{t})$$

An online algorithm H learns without regret if in the limit as T goes into infinity, its **average** regret goes to zero in the worst case. Meaning no single expert is better than H in the limit its average regret goes to zero in the worst case - no single expert better than H in the limit.

Average regret is the cumulative regret divided by time steps.

Information sets: players perspective of game state (i.e. my cards + opp cards + community cards)


## No regret learning
Repeatedly making decision in an uncertain environment. Receive advice from N experts about a single phenomenon. Our online algorithm H's goal is to distribute trust among experts.


After outcome is revealed, there is a loss vector that evaluates the N experts.

## Counter factual regret minimisation
Optimises entity called immediate counterfactual regret, which is an upper bound on average overall regret.

Guaranteed to go to zero in the limit if regret matching is applied.

### Counter factual utility
$$ \sum_{h' part of Z} \pi^{\sigma} (h, h')u(h')$$

Z is the set of all terminal nodes

If this is considered for every possible game state (every possible opponent's hand), we will end up with a vector of utilities, one for each hand.


For counterfactual utility, another weighting scheme is used. The probability of reaching h, assuming that we wanted to get to h is used instead.

Instead of using the regular strategy from strategy profile,


## Memory Requirements

Distinguish: reward for playing action a at time t
Every action will be rewarded via unnormalized counter factual utility assuming action was played.
