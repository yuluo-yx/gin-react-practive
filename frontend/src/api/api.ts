import fetch from "./request";

// 登陆 api
export const LoginApi = (params: { user_name: string; password: string }) =>
  fetch("/api/v1/user/login", params, "POST");

// 获取留言信息列表（获取全部）
export const MsgList = () => fetch("/api/v1/msg/list", null, "GET");

// 获取留言信息列表（分页获取）
export const MsgPages = (params: { page_num: number; page_size: number }) =>
  fetch("/api/v1/msg/list", params, "GET");

// 保存留言信息
export const publishMsg = (params: { text: string; user_id: number }) =>
  fetch("/api/v1/msg/publish", params);
