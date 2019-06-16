import axios from 'axios';

export default {
  make: (method, path) => axios({
    method,
    url: `${process.env.VUE_APP_API_BASE_URL || 'http://localhost:3000'}/v1/${path}`,
  }),
};
