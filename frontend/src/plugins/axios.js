import globalAxios from 'axios';
export default ({store}) => {
  const manager = globalAxios.interceptors.request;
  if (manager.handlers.length == 0){
    globalAxios.interceptors.request.use(req => {
      const token = store.$auth.getToken('local');
      if (token)
      {
        req.headers.Authorization = token;
      }
      return req
    })
  }
}
