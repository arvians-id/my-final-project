export const getStatusSubmision = (submision) => {
  if (submision.hasOwnProperty('file')) {
    return true;
  }
  return false;
};
