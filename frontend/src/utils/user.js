import { usernameRegexPattern } from "./reqex";

export const checkIsValidUsername = (username) => {
  if (
    usernameRegexPattern.test(username) &&
    username.length > 2
  ) {
    return true;
  }
  return false;
};