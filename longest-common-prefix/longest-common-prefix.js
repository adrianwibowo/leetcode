/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    if (strs.length == 1) return strs[0]
    
    let result = ""

    strs.sort() // sort alphabetically, so we can just compare the first and last element only because if it's sorted the middle words will be automatically same as first and last

    for (let i = 0; i < strs[0].length; i++) {
        if (!strs[0].length) return ""

        // check every char for first and last index
        if (strs[0][i] == strs[strs.length-1][i]) { 
            result += strs[0][i]
        } 
        // stop compare whenever the char is not match and return result   immediately
        else { 
        break
        }
    }

    return result
};