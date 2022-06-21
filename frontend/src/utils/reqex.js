/**
 *  RegEx	            Description
 * ^	                The password string will start this way
 * (?=.*[a-z])	      The string must contain at least 1 lowercase alphabetical character
 * (?=.*[A-Z])	      The string must contain at least 1 uppercase alphabetical character
 * (?=.*[0-9])	      The string must contain at least 1 numeric character
 * (?=.*[!@#$%^&*])	  The string must contain at least one special character, but we are 
 *                    escaping reserved RegEx characters to avoid conflict
 * (?=.{8,})	        The string must be eight characters or longer
 */

export const strongRegexPatternPassword = new RegExp("^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{10,})");

export const mediumRegexPatternPassword = new RegExp("^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{7,})");

/* 
  Username can only have: 
  - Uppercase Letters (A-A) 
  - Lowercase Letters (a-z) 
  - Numbers (0-9)
  - Dots (.)
  - Underscores (_)
*/
export const usernameRegexPattern = new RegExp(/^[A-Za-z0-9_\.]+$/)