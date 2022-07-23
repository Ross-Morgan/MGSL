#include <cmath>
#include <iostream>
#include <numeric>
#include <vector>


namespace math {
    // SECTION Generic Stubs

    // !SECTION
    // SECTION Math Stubs

    float arithmeticMean(std::vector<double>);
    float geometricMean(std::vector<double>);
    float harmonicMean(std::vector<double>);
    float weightedHarmonicMean(std::vector<double>, std::vector<double>);

    float arithmeticMean(std::vector<float>);
    float geometricMean(std::vector<float>);
    float harmonicMean(std::vector<float>);
    float weightedHarmonicMean(std::vector<float>, std::vector<float>);

    float arithmeticMean(std::vector<int>);
    float geometricMean(std::vector<int>);
    float harmonicMean(std::vector<int>);
    float weightedHarmonicMean(std::vector<int>, std::vector<float>);


    // !SECTION
    // SECTION Arithmetic Mean

    float
    arithmeticMean(std::vector<float> nums)
    {
        // Number of elements
        size_t len = nums.size();

        // Sum of elements in vector
        float sum = std::accumulate(nums.begin(), nums.end(), 0);

        return sum / len;
    }

    // !SECTION
    // SECTION Geometric Mean

    float
    geometricMean(std::vector<float> nums) {
        // Number of elements
        size_t len = nums.size();

        // Product of elements
        auto product = std::accumulate(std::begin(nums), std::end(nums), 1, std::multiplies<>());

        return pow(product, 1 / len);
    }

    // !SECTION
    // SECTION Harmonic Mean

    float
    harmonicMean(std::vector<float> nums){
        // Number of elements
        size_t len = nums.size();

        // Vector for reciprocals of elements
        std::vector<float> inv_nums;

        // Fill vector with reciprocals
        for (int i = 0; i < len; i++)
            inv_nums.push_back(1 / nums[i]);

        // Sum of reciprocal of elements in vector
        float inv_sum = std::accumulate(inv_nums.begin(), inv_nums.end(), 0);

        return len / inv_sum;

    }

    // !SECTION
    // SECTION Weighted Harmonic Mean

    float
    weightedHarmonicMean(std::vector<float> nums, std::vector<float> weights)
    {
        // Number of elements
        size_t len = nums.size();
        // Number of weights
        size_t nWeights = weights.size();

        if (len != nWeights) {
            if (len > nWeights)
            { std::cerr << "Not enough weights provided"; }
            if (len < nWeights)
            { std::cerr << "Not enough numbers provided"; }
        }
    }

    // !SECTION
}