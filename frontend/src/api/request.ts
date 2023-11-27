import axios, { AxiosError, AxiosResponse } from "axios";

const baseURL = "/api"; //此处的 /api 与跨域有关,请跳转下方跨域部分

interface IResponeData {
  data: any;
  status: number;
  error: string;
  msg: string;
}

//fetch定义请求的类型约束
//当不传输值的时候默认请求方法为POST ,默认不携带token
export default function fetch(
  url: string,
  params: any,
  method: string = "POST",
  oheaders: object = {}
) {
  return new Promise<IResponeData>((resolve, _) => {
    //请求头部分
    let headers = Object.assign(
      {},
      { "x-auth-token": localStorage.getItem("token") },
      oheaders
    );

    axios({
      // 定义发起请求时的基本类型
      method: method,
      url: baseURL + url,
      data: ["POST", "PUT"].includes(method) ? params : null,
      params: ["GET"].includes(method) ? params : null,
      headers: headers,
    })
      .then((res: AxiosResponse<any, any>) => {
        resolve(res.data);
      })
      .catch((err: AxiosError) => {
        resolve(err.response?.data as IResponeData);
      });
  });
}

// 添加请求拦截器
axios.interceptors.request.use(
  function (config) {
    // 在发送请求之前做些什么
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
axios.interceptors.response.use(
  function (config) {
    // 在请求响应之前做些什么
    return config;
  },
  function (error) {
    // 对响应错误做些什么
    return Promise.reject(error);
  }
);
