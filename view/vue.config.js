const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api/v1': {
          target: 'http://121.4.98.41:80', //API服务器的地址
          changeOrigin: true,
          pathRewrite: {
              
          }
      }
  }}
})