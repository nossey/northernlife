import globalAxios from 'axios';
export default ({store}) => {
  globalAxios.interceptors.request.use(req => {
    const token = store.$auth.getToken('local');
    if (token && !globalAxios.defaults.headers.Authorization)
    {
      req.headers.Authorization = token;
    }
    return req
  })
}
