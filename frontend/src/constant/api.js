export const BASE_URL = '' + process.env.REACT_APP_TEENAGER_BACKEND_URL;

// USER AUTH
export const ENDPOINT_API_POST_LOGIN_USER = `${BASE_URL}/api/users/login`;
export const ENDPOINT_API_GET_CHECK_LOGIN_USER = `${BASE_URL}/api/userstatus`;
export const ENDPOINT_API_POST_REGISTER_USER = `${BASE_URL}/api/users`;

// COURSE
export const ENDPOINT_API_GET_COURSE_BY_USER_LOGIN = `${BASE_URL}/api/usercourse/courses`;
export const ENDPOINT_API_GET_COURSE_BY_CODE = `${BASE_URL}/api/courses`;

// SUBMISSION
export const ENDPOINT_API_GET_SUBMISSIONS_BY_USER = `${BASE_URL}/api/users/submissions`;
export const ENDPOINT_API_USER_SUBMIT_SUBMISSIONS = (courseCode, submissionId) => `${BASE_URL}/api/courses/${courseCode}/submissions/${submissionId}/user-submit`;

// DISCUSSION
export const ENDPOINT_API_GET_QUESTIONS_BY_USER = `${BASE_URL}/api/questions/by-user`;

// USER
export const ENDPOINT_API_GET_USER_BY_ID = `${BASE_URL}/api/users`;
export const ENDPOINT_API_UPDATE_USER_DETAIL_PROFIE = `${BASE_URL}/api/users`;

