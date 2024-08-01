from scipy import stats


num_heads = 16
num_flips = 20
prob_head = 0.5
prob = stats.binomtest(num_heads, num_flips, prob_head)

prob_16_heads = stats.binom.pmf([4, 16], num_flips, prob_head)

print(prob_16_heads)
