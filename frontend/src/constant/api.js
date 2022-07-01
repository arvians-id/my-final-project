export const BASE_URL = '' + process.env.REACT_APP_TEENAGER_BACKEND_URL;

// USER AUTH
export const ENDPOINT_API_POST_LOGIN_USER = `${BASE_URL}/api/users/login`;
export const ENDPOINT_API_GET_CHECK_LOGIN_USER = `${BASE_URL}/api/userstatus`;
export const ENDPOINT_API_POST_REGISTER_USER = `${BASE_URL}/api/users`;

// COURSE
export const ENDPOINT_API_GET_COURSE_BY_USER_LOGIN = `${BASE_URL}/api/usercourse/courses`;
export const ENDPOINT_API_GET_COURSE_BY_CODE = `${BASE_URL}/api/courses`;
export const ENDPOINT_API_CREATE_COURSE = `${BASE_URL}/api/courses`;
export const ENDPOINT_API_GET_ALL_COURSE = `${BASE_URL}/api/courses`;
export const ENDPOINT_API_UPDATE_COURSE = (courseCode) => `${BASE_URL}/api/courses/${courseCode}`;
export const ENDPOINT_API_DELETE_COURSE = (courseCode) => `${BASE_URL}/api/courses/${courseCode}`;
export const ENDPOINT_API_LIST_USER_IN_COURSE = (courseCode) => `${BASE_URL}/api/courses/${courseCode}/users`;
export const ENDPOINT_API_REMOVE_USER_IN_COURSE = (userId, courseCode) => `${BASE_URL}/api/usercourse/${userId}/${courseCode}`;
export const ENDPOINT_API_ADD_USER_IN_COURSE = `${BASE_URL}/api/usercourse`;

// SUBMISSION
export const ENDPOINT_API_GET_SUBMISSIONS_BY_USER = `${BASE_URL}/api/users/submissions`;
export const ENDPOINT_API_USER_SUBMIT_SUBMISSIONS = (courseCode, submissionId) => `${BASE_URL}/api/courses/${courseCode}/submissions/${submissionId}/user-submit`;

// DISCUSSION   
export const ENDPOINT_API_GET_QUESTIONS_BY_USER = `${BASE_URL}/api/questions/by-user`;
export const ENDPOINT_API_GET_ALL_QUESTIONS = `${BASE_URL}/api/questions/all`;

// USER
export const ENDPOINT_API_GET_USER_BY_ID = `${BASE_URL}/api/users`;
export const ENDPOINT_API_UPDATE_USER_DETAIL_PROFIE = `${BASE_URL}/api/users`;
export const ENDPOINT_API_GET_LIST_USER = `${BASE_URL}/api/users`;

// MODULE ARTICLES
export const ENDPOINT_API_CREATE_MODULE_ARTICLES = (courseCode) => `${BASE_URL}/api/courses/${courseCode}/articles`;
export const ENDPOINT_API_GET_LIST_MODULE_ARTICLES_BY_COURSE_CODE = (courseCode) => `${BASE_URL}/api/courses/${courseCode}/articles`;
export const ENDPOINT_API_DELETE_MODULE_ARTICLES = (courseCode, articleId) => `${BASE_URL}/api/courses/${courseCode}/articles/${articleId}`;
export const ENDPOINT_API_UPDATE_MODULE_ARTICLES = (courseCode, articleId) => `${BASE_URL}/api/courses/${courseCode}/articles/${articleId}`;
export const ENDPOINT_API_NEXT_MODULE_ARTICLES = (courseCode, articleId) => `${BASE_URL}/api/courses/${courseCode}/articles/${articleId}/next`;
export const ENDPOINT_API_PREV_MODULE_ARTICLES = (courseCode, articleId) => `${BASE_URL}/api/courses/${courseCode}/articles/${articleId}/previous`;
export const ENDPOINT_API_GET_DETAIL_ARTICLE_COURSE_BY_ARTICLE_ID_AND_COURSE_CODE = (code, articleId) => `${BASE_URL}/api/courses/${code}/articles/${articleId}`
