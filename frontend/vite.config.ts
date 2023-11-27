import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],

  // 解决跨域
  server: {
    // 前端服务启动地址 默认为 ‘0.0.0.0’
    port: 9000, // 设置服务启动端口号
    open: true, // 设置服务启动时是否自动打开浏览器

    proxy: {
      '/api': {
        target: 'http://127.0.0.1:3000',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },

})
