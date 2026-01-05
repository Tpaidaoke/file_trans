// vite.config.js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// 直接使用 import.meta.url 和 process.cwd() 的组合
const root = process.cwd()

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(root, 'src')
    }
  },
  server: {
    host: '0.0.0.0',
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:9003',
        changeOrigin: true,
        rewrite: (path) => path
      }
    }
  }
})