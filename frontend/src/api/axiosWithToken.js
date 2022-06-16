import Axios from 'axios';
import { localLoadToken } from '../utils/token';
import { localClearToken } from '../utils/token';

// ini request axios tapi kita upgrade dengan masang token di header buat otentifikasi user di server
// lalu otomatis logout jika token sudah ga valid

export const axiosWithToken = (options = {}) => {
    const token = localLoadToken();
    const config = token
        ? {
            headers: {
                'Authorization': token,
            },
        }
        : {};
    const instance = Axios.create(config);

    instance.interceptors.response.use(
        function (response) {
            return response;
        },
        function (error) {
            switch (error.response?.status) {
                case 400:
                    if (
                        !error.response.data.success &&
                        error.response.data.message ===
                        'expired token, please relogin'
                    ) {
                        localClearToken();
                        window.location.href = '/login';
                    }
                    break;
                default:
            }
            return Promise.reject(error);
        }
    );
    return instance;
};
