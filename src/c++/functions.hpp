#include <numeric>
#include <vector>

namespace math {
    float arithmeticMean(std::vector<int>);
    float arithmeticMean(std::vector<float>);

    float geometricMean(std::vector<int>);
    float geometricMean(std::vector<float>);

    float harmonicMean(std::vector<int>);
    float harmonicMean(std::vector<float>);

    float weightedHarmonicMean(std::vector<int>, std::vector<float>);
    float weightedHarmonicMean(std::vector<float>, std::vector<float>);





    // SECTION Arithmetic Mean


    float
    arithmeticMean(std::vector<int> nums)
    {
        // Number of elements
        size_t len = nums.size();

        // Sum of elements in vector
        int sum = std::accumulate(nums.begin(), nums.end(), 0);

        return sum / len;
    };

    float
    arithmeticMean(std::vector<float> nums)
    {
        // Number of elements
        size_t len = nums.size();

        // Sum of elements in vector
        float sum = std::accumulate(nums.begin(), nums.end(), 0);

        return sum / len;
    };

    // !SECTION
    // SECTION Geometric Mean

    float
    geometricMean(std::vector<int> nums) {
        // Number of elements
    }

    geometricMean(std::vector<float> nums) {

    }

    // !SECTION
    // SECTION Harmonic Mean

    float
    harmonicMean(std::vector<int> nums){
        // Number of elements
        size_t len = nums.size();

        std::vector<float> inv_nums;

        // Sum of reciprocal of elements in vector
        for (int i = 0; i < len; i++)
            nums.push_back(1 / nums[i]);

        //
        float inv_sum = std::accumulate(inv_nums.begin(), inv_nums.end(), 0);

        return len / inv_sum;
    };

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

    };

    // !SECTION
    // SECTION Weighted Harmonic Mean

    float
    weightedHarmonicMean(std::vector<int>)
    {
        // Number of elements
        size_t
    };

    float
    weightedHarmonicMean(std::vector<float>)
    {

    };

    // !SECTION
}