export const setLocal = (key, value) => {
  localStorage.setItem(key, JSON.stringify(value));
};

export const getLocal = (key, fallback) => {
  const result = localStorage.getItem(key);
  if (result && result !== "undefined") {
    return JSON.parse(result);
  } else {
    return fallback ?? false;
  }
};

export const clearLocal = (key) => {
  localStorage.removeItem(key);
};
