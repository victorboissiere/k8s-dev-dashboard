import axios from 'axios';

export default {
  make: (method, path) => axios({
    method,
    url: `http://localhost:3000/v1/${path}`,
  }),
};
