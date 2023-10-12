/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let result = []
    for (let i = 0 ; i < nums.length ; i++) {
        for (let b = i+1 ; b < nums.length ; b++) {
            if ((nums[i] + nums[b]) == target) {
                result.push(i)
                result.push(b)
                return result
            }
        }
    }
};