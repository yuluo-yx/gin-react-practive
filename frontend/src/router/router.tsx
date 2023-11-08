// import React, { lazy } from "react";
import { Navigate } from "react-router-dom";
import { RouteObject } from "react-router-dom";
import Login from "../views/login";
import Test from "../views/test"
import Message from "../views/message"

//懒加载模式的组件写法，外面需要套一层loading 的提示加载组件
// const Login= lazy(()=>import("../views/Login"));
// 使用 element: withLoadingComponent(<Login />)
// const withLoadingComponent = (comp: JSX.Element) => {
//   return (
//     <React.Suspense fallback={<div>loading...</div>}>{comp}</React.Suspense>
//   );
// };

const routes: RouteObject[] = [
  {
    path: "/",
    element: <Navigate to="/login" />,
  },
  {
    path: "/test",
    element: <Test />,
  },
  {
    path: "/msg",
    element: <Message />,
  },
  {
    path: "/login",
    element: <Login />,
  },
];

export default routes;
