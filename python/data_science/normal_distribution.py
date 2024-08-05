import numpy as np
import matplotlib.pyplot as plt


np.random.seed(0)
sample_size = 10000
# sample = np.array([np.random.binomial(1, 0.5) for _ in range(sample_size)])
# head_count = sample.sum()
# head_count_frequency = head_count / sample_size

frequencies = np.random.binomial(sample_size, 0.5, 100000) / sample_size

sample_means = frequencies
likelihoods, bin_edges, _ = plt.hist(sample_means, bins='auto', edgecolor='black', density=True)

plt.xlabel('Binned Sample Mean')
plt.ylabel('Relative Likelihood')
plt.show()

# 计算均值和标准差
mean_normal = np.average(bin_edges[:-1], weights=likelihoods)
var_normal = 
