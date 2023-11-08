import fetch from "./request";

export const LoginApi = (params: { password: string; username: string;}) => fetch("/login", params)


