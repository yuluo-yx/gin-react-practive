import {useRoutes} from "react-router-dom"
import router from  "./router/router"

function App() {

  const route = useRoutes(router);

  return (
    <>
      {/* 引用路由组件 */}
      {route}
    </>
  )
}

export default App
