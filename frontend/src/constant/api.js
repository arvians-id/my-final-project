export const BASE_URL = '' + process.env.REACT_APP_TEENAGER_BACKEND_URL;

// USER AUTH
export const ENDPOINT_API_POST_LOGIN_USER = `${BASE_URL}/api/users/login`;
export const ENDPOINT_API_GET_CHECK_LOGIN_USER = `${BASE_URL}/api/userstatus`;
export const ENDPOINT_API_GET_REGISTER_USER = `${BASE_URL}/api/users`;