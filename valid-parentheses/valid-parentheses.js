/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {   
    // sanity check
     if ((s.length % 2) !== 0 ){
         return false
     }
     
     const pattern = /[^{}()[\]]/
 
     if (pattern.test(s)) {
         return false
     }
 
     let container = []
     
     for (let i = 0; i < s.length ; i++) {
         // store open bracket
           if (s[i] === "{" ||  s[i] === "[" || s[i] === "(") {
               container.push(s[i])
           }
           
           //once meet close bracket, pop the stored open bracket and match the bracket, if not return false
           if (s[i] === "}" && container.pop() !== "{") {
                return false
           } else if (s[i] === "]" && container.pop() !== "[") {
                return false
           } else if (s[i] === ")" && container.pop() !== "(") {
                return false  
           }
     }
     return container.length === 0;
 };